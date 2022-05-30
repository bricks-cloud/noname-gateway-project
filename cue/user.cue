package cue

import "github.com/bricks-cloud/bricks/cue/snowflake"

snowflake.deployment & {
	provider: snowflake: {
		username: "testuser"
		account:  "testaccount"
		region:   "testregion"
		password: "fake"
	}
	resource: snowflake_user: {
		"test": {
			name: "testname"
		}
	}
}
