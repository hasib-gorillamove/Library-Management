{
	"info": {
		"_postman_id": "d5da1876-6e3f-4c70-af7e-9a8453b567b5",
		"name": "golden-infotech-service",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "25328998"
	},
	"item": [
		{
			"name": "Create",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"title\" :\"Harry Potter\",\r\n    \"author\":\"J. K. Rowling\",\r\n    \"publication_year\":1997\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "{{URL}}:2001/api/v1/books",
					"host": [
						"{{URL}}"
					],
					"port": "2001",
					"path": [
						"api",
						"v1",
						"books"
					]
				}
			},
			"response": []
		},
		{
			"name": "ListAllBooks",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{URL}}:2001/api/v1/books?limit=2&page=4",
					"host": [
						"{{URL}}"
					],
					"port": "2001",
					"path": [
						"api",
						"v1",
						"books"
					],
					"query": [
						{
							"key": "keyword",
							"value": "javascript",
							"disabled": true
						},
						{
							"key": "author",
							"value": "media",
							"disabled": true
						},
						{
							"key": "publication_year",
							"value": "2017",
							"disabled": true
						},
						{
							"key": "limit",
							"value": "2"
						},
						{
							"key": "page",
							"value": "4"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "Get A Book",
			"request": {
				"method": "GET",
				"header": []
			},
			"response": []
		},
		{
			"name": "Update",
			"request": {
				"method": "PUT",
				"header": [],
				"url": {
					"raw": "{{URL}}:2001/api/v1/books/10",
					"host": [
						"{{URL}}"
					],
					"port": "2001",
					"path": [
						"api",
						"v1",
						"books",
						"10"
					]
				}
			},
			"response": []
		},
		{
			"name": "Delete",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "{{URL}}:2001/api/v1/books/5",
					"host": [
						"{{URL}}"
					],
					"port": "2001",
					"path": [
						"api",
						"v1",
						"books",
						"5"
					]
				}
			},
			"response": []
		}
	]
}