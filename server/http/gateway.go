package http

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"

	"github.com/bricks-cloud/noname-gateway-project/common"
)

type Middleware func(http.RoundTripper) http.RoundTripper

type router interface {
	HandlerFunc(method, path string, handler http.HandlerFunc)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Gateway struct {
	middlewareFactory map[MiddlewareName]Middleware
	client            *client
	router            router
}

func NewRouter(mf map[MiddlewareName]Middleware, cc common.ClientConfig, rc common.ServiceConfig) (*Gateway, error) {
	g := &Gateway{
		middlewareFactory: mf,
		client:            NewClient(cc),
		router:            httprouter.New(),
	}

	err := g.configureRouter(rc.GetRouteConfigs())

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Gateway) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			buf := make([]byte, 64<<10)
			n := runtime.Stack(buf, false)
			log.Error().Msgf("panic recovered: %s", buf[:n])
		}
	}()

	g.router.ServeHTTP(w, req)
}

func (g *Gateway) configureRouter(rcs []common.RouteConfig) error {
	for _, rc := range rcs {
		handlerFunc, err := g.buildHandlerFunc(rc)
		if err != nil {
			return err
		}

		g.router.HandlerFunc(rc.GetMethod(), rc.GetPath(), handlerFunc)
	}

	return nil
}

// buildHandlerFunc builds a http request handler function for the router
func (g *Gateway) buildHandlerFunc(rc common.RouteConfig) (http.HandlerFunc, error) {
	tripper, err := g.useMiddleware([]MiddlewareName{}, g.client)
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		setXFFHeader(req)
		ctx, cancel := context.WithTimeout(req.Context(), rc.GetRetryTimeOut())
		defer cancel()

		body, err := io.ReadAll(req.Body)
		if err != nil {
			return
		}
		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body)
			return ioutil.NopCloser(reader), nil
		}

		var resp *http.Response
		if err = ctx.Err(); err != nil {
			return
		}

		tryCtx, cancel := context.WithTimeout(ctx, rc.GetPerRetryTimeout())
		defer cancel()
		reader := bytes.NewReader(body)
		req.Body = ioutil.NopCloser(reader)
		resp, err = tripper.RoundTrip(req.WithContext(tryCtx))
		if err != nil {
			return
		}

		if err != nil {
			return
		}

		headers := w.Header()
		for k, v := range resp.Header {
			headers[k] = v
		}

		w.WriteHeader(resp.StatusCode)
		if body := resp.Body; body != nil {
			_, err := io.Copy(w, body)
			if err != nil {
				return
			}
		}

		// see https://pkg.go.dev/net/http#example-ResponseWriter-Trailers
		for k, v := range resp.Trailer {
			headers[http.TrailerPrefix+k] = v
		}

		if resp.Body != nil {
			resp.Body.Close()
		}
	}), nil
}

func (g *Gateway) useMiddleware(middlewares []MiddlewareName, next http.RoundTripper) (http.RoundTripper, error) {
	for _, mw := range middlewares {
		m, ok := g.middlewareFactory[mw]
		if !ok {
			return nil, NewNotFoundError("middleware not found")
		}
		next = m(next)
	}
	return next, nil
}
