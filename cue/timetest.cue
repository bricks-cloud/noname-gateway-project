package cue

import "github.com/bricks-cloud/bricks/cue/time"

time.deployment & {
	resource: time_offset: {
		"test": {
			offset_days: 7
		}
	}
}
