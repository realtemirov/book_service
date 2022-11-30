include .env
CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

createdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} createdb --password --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DATABASE}

dropdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} dropdb ${POSTGRES_DATABASE}

psqlcontainer:
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -d -p ${POSTGRES_PORT}:5432 --env-file .env postgres:15-alpine3.16

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

copy-proto-module:
	rm -rf ${CURRENT_DIR}/protos
	rsync -rv --exclude={'/.git','LICENSE','README.md'} ${CURRENT_DIR}/../finiso_protos/* ${CURRENT_DIR}/protos

gen-proto-module:
	./scripts/gen_proto.sh ${CURRENT_DIR}

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

swag-init:
	swag init -g api/api.go -o api/docs

.PHONE: pull-proto-module, update-proto-module, copy-proto-module, migration-up, migration-down, build, build-image, push-image, swag-init, psqlcontainer
