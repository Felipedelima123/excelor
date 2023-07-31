# Excelor
Excelor aims to be a simple API to generate Excel, from anywhere. Powered by [Gin](https://gin-gonic.com) and [Excelize](https://xuri.me/excelize/) Streaming API, is pretty fast, lightweight and robust.


## Features
- Create excel file from a JSON;
- Uploads the created file to S3;
- Run with docker;
- Standalone app;
- Can be a binary to distribute to your servers and run locally.

## How to use
In the moment the api is in an embryonic stage, so there is only one endpoint. 

HTTP POST -> /generate-excel
```JSON
{
	"columns": [
		"Name",
		"Age",
		"Role"
	],
	"rows": [
		[
			"José",
			"31",
			"Admin"
		],
		[
			"Lucia",
			"32",
			"Manager"
		],
		[
			"Pedro",
			"32",
			"Developer"
		]
	],
	"sheetName": "A Cool Name"
}
```

CURL Code:
```SHELL
curl --request POST \
  --url http://localhost:8080/generate-excel \
  --header 'Content-Type: application/json' \
  --data '{
	"columns": [
		"Name",
		"Age",
		"Role"
	],
	"rows": [
		[
			"José",
			"31",
			"Admin"
		],
		[
			"Lucia",
			"32",
			"Manager"
		],
		[
			"Pedro",
			"32",
			"Developer"
		]
	],
	"sheetName": "A Cool Name"
}'
```

## TODO
- Everything :D

