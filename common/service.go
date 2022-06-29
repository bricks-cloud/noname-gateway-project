package common

type ServiceConfig interface {
	GetRouteConfigs() []RouteConfig
}
