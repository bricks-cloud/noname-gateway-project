package snowflake

// To install this provider, copy and paste this code into your Terraform configuration. 
// Then, run terraform init. Terraform 0.13+

deployment: {
	terraform: {
		required_providers: {
			snowflake: {
				source:  "chanzuckerberg/snowflake"
				version: "0.33.1"
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

// snowflake user and roles
#snowflake_user: [_]: {
	name: string
	default_role?: string
	default_warehouse?: string
	password?: string
	rsa_public_key?: string
}

#snowflake_role: [_]: {
	name: string
}

#snowflake_role_grants: [_]: {
	role_name: string
	roles?: [...string]
	users?: [...string]
}

// snowflake database
#snowflake_database: [_]: {
	name: string
	data_retention_time_in_days?: number
}

#snowflake_database_grant: [_]: {
	database_name: string
	privilege?: string
	roles?: [...string]
	shares?: [...string]
	with_grant_option?: bool
}

// schema, table
#snowflake_schema: [_]: {
	database: string
	name: string
	data_retention_days: number
}

#column: [_]: {
	name: string
	type: string
	nullable?: bool
}

#primary_key: [_]: {
	name?: string
	keys: [...string]
}

#snowflake_table: [_]: {
	name: string
	database: string
	schema: string
}

#tag: [_]: {
	name: string
	value: string
}