package time

deployment: {
	terraform: {
		required_providers: {
			time: {
				source:  "hashicorp/time"
				version: "0.7.2"
			}
		}
	}

	provider: {
		time: {}
	}

	resource: {
		time_offset: #time_offset
	}
}

#time_offset: [_]: {
	offset_days: int
}
