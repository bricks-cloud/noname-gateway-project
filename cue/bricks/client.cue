package bricks

// #Client: {
// 	$bricks: type: _name: "Client"
// 	name: string
// }

// #Resource: {
// 	$bricks: type: _name: "Resource"
// 	name: string
// }

// spotifyClient: #Client & {
// 	name: "spotifyClientNameBLahblah"
// }

#service: {
	name: string
	description?: string
	url: string
	// routes: [...#route]
}

#route: {
	name: string
	description?: string
	paths: [...string]
}

config: #service & {
    name: "my-service"
    description: "mock description for service"
    url: "sample.com/my-service"
    // routes: [
    //     { 
    //     name: "my-first-route"
    //     paths: ["/first", "/route"]
    //     },
    //     {
    //     name: "my-second-route"
    //     paths: ["/home"]
    //     }
    // ]
}

// Vanguard is a data monitor alerting system for SQL databses
// #VanguardClient: #Client & {
// 	source:       "BigQuery"
// 	name:         "VanguardClient"
// 	alertChannel: "Webhook"

// 	// credentials
// 	credentials: {
// 		if *source == "BigQuery" {
// 			serviceAccountKey: string
// 		}
// 	}

// 	if *alertChannel == "Webhook" {
// 		webhookURL: =~"https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"
// 	}
// }

// DataQualityAlert is a periodically running SQL scheduled job
// that detects data quality anomalies when it happens
// #DataQualityAlert: #Resource & {
// 	name: "DataQualityAlert"

// 	// schedule shoud be in unix cron tab format
// 	schedule: =~"/^(\\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\\*\\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\\*|([0-9]|1[0-9]|2[0-3])|\\*\\/([0-9]|1[0-9]|2[0-3])) (\\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\\*\\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\\*|([1-9]|1[0-2])|\\*\\/([1-9]|1[0-2])) (\\*|([0-6])|\\*\\/([0-6]))$/"

// 	sqlQuery:  string
// 	condition: "=" | "<" | ">" | ">=" | "<="
// 	threshold: int | string | bool
// }

// #Deployment: {
// 	$bricks: type: _name: "Deployment"
// 	environmentSettings: [
// 		{
// 			region： "",
// 		},
// 		{
// 			region： "",
// 		}
// 	]
// }
