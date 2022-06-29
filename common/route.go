package common

import "time"

type RouteConfig interface {
	GetTargetPath() string
	GetPath() string
	GetMethod() string
	GetRetryTimeOut() time.Duration
	GetPerRetryTimeout() time.Duration
}
