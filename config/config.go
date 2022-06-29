package config

import "time"

type ClientConfig struct {
	maxIdleConns          int
	maxIdleConnsPerHost   int
	maxConnsPerHost       int
	disableCompression    bool
	idleConnTimeout       time.Duration
	tLSHandshakeTimeout   time.Duration
	expectContinueTimeout time.Duration
	dialTimeout           time.Duration
	dialKeepAlive         time.Duration
}

func (cc *ClientConfig) GetMaxIdleConns() int {
	return cc.maxIdleConns
}

func (cc *ClientConfig) GetMaxIdleConnsPerHost() int {
	return cc.maxIdleConnsPerHost
}

func (cc *ClientConfig) GetMaxConnsPerHost() int {
	return cc.maxConnsPerHost
}

func (cc *ClientConfig) GetDisableCompression() bool {
	return cc.disableCompression
}

func (cc *ClientConfig) GetIdleConnTimeout() time.Duration {
	return cc.idleConnTimeout
}

func (cc *ClientConfig) GetTLSHandshakeTimeout() time.Duration {
	return cc.tLSHandshakeTimeout
}

func (cc *ClientConfig) GetExpectContinueTimeout() time.Duration {
	return cc.expectContinueTimeout
}

func (cc *ClientConfig) GetDialTimeout() time.Duration {
	return cc.dialTimeout
}

func (cc *ClientConfig) GetDialKeepAlive() time.Duration {
	return cc.dialKeepAlive
}

type RouteConfig struct {
	targetPath      string
	path            string
	method          string
	retryTimeout    time.Duration
	perRetryTimeout time.Duration
}

func (rc *RouteConfig) GetTargetPath() string {
	return rc.targetPath
}

func (rc *RouteConfig) GetPath() string {
	return rc.path
}

func (rc *RouteConfig) GetMethod() string {
	return rc.method
}

func (rc *RouteConfig) GetRetryTimeOut() time.Duration {
	return rc.retryTimeout
}

func (rc *RouteConfig) GetPerRetryTimeout() time.Duration {
	return rc.perRetryTimeout
}

type ServerConfig struct {
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
}

func (sc *ServerConfig) GetReadTimeout() time.Duration {
	return sc.readTimeout
}

func (sc *ServerConfig) GetReadHeaderTimeout() time.Duration {
	return sc.readHeaderTimeout
}

func (sc *ServerConfig) GetWriteTimeout() time.Duration {
	return sc.writeTimeout
}

func (sc *ServerConfig) GetIdleTimeout() time.Duration {
	return sc.idleTimeout
}

type ServiceConfig struct {
	routeConfigs []RouteConfig
}

func (sc *ServiceConfig) GetRouteConfigs() []RouteConfig {
	return sc.routeConfigs
}
