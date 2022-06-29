package common

import "time"

type ServerConfiguration interface {
	GetReadTimeout() time.Duration
	GetReadHeaderTimeout() time.Duration
	GetWriteTimeout() time.Duration
	GetIdleTimeout() time.Duration
}
