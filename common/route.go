package common

type RouteConfig interface {
	GetTargetURL() string
	GetURL() string
}
