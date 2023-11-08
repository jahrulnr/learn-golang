app=goservice:1.0
container=go-a

build-app: ###
ifeq ($(shell docker images -q ${app} 2> /dev/null),)
	@echo $(app)
	docker build -f Dockerfile-local -t ${app} .
endif
.PHONY: build-app

up: build-app ### Run docker-compose if app image not found will build
	docker-compose up --build -d service-a && docker-compose logs -f
.PHONY: up

down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: down

clean: down ###remove image fpm custom
	docker rmi -f $(shell docker images -q ${app})
.PHONY: clean

rebuild: clean
	make up
### docker exec $(container) sh -c "go build -o binary"

shell:
	docker exec -it $(container) sh

go-build:
	go build -o ./build/app1
go-run:
	INSTANCE_ID=test PORT=7080 ./build/${app} 