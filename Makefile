swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

run:
	go run main.go

get:
	curl -v localhost:9091/ | jq

put:
	curl -v localhost:9091/1 -XPUT -d '{"id": 12345, "name": "tea", "description": "hello world 5", "price": 1.0, "sku": "fdc-ert-hfs"}' | jq

post:
	curl -v -d '{"name": "tea", "price": 20.12, "sku": "abc-dfv-dsg"}' localhost:9091/ | jq