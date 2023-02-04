GOSOURCEFILE="./main.go"
SWAGDOCS="./docs/swagger"

gen:
	@echo "Running Progo Generator ..."
	@protoc --go_out=plugins=grpc:. model/*.proto
	@echo "Success"

gen-run:
	@echo "Running Progo Generator ..."
	@protoc --go_out=. --go-grpc_out=. model/*.proto
	@echo "Success"

r:
	go run main.go

run:
	echo "Update Swagger Docs"
	swag init -g ./$(GOSOURCEFILE) -o $(SWAGDOCS)
	nodemon --exec go run $(GOSOURCEFILE)  --signal SIGTERM

swag:
	echo "Create Swagger files"
	swag init -g ./$(GOSOURCEFILE) -o $(SWAGDOCS)