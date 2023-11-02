run:
	go run main.go

get:
	curl -v localhost:9090/ | jq

put:
	curl -v localhost:9090/1 -XPUT -d '{"id": 12345, "name": "tea", "description": "hello world 5"}' | jq

post:
	curl -v -d '{"name": "tea", "price": 20.12, "sku": "abc-dfv-dsg"}' localhost:9090/ | jq