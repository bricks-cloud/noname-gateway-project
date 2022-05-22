package snowflake

deployment: {
	terraform: {
		required_providers: {
			snowflake: {
				source : "chanzuckerberg/snowflake"
				version : "0.33.1"
    		}
  		}
	}

	provider: {
		snowflake: {
			username: string
			account:  string
			region:   string

			password?:           string
			oauth_access_token?: string
			private_key_path?:   string
		}
	}

	resource: {
		snowflake_user: #snowflake_user
	}
}

#snowflake_user: [_]: {
	name: string
}
