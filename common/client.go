package common

import "time"

type ClientConfig interface {
	GetMaxIdleConns() int
	GetMaxIdleConnsPerHost() int
	GetMaxConnsPerHost() int
	GetDisableCompression() bool
	GetIdleConnTimeout() time.Duration
	GetTLSHandshakeTimeout() time.Duration
	GetExpectContinueTimeout() time.Duration
	GetDialTimeout() time.Duration
	GetDialKeepAlive() time.Duration
}
