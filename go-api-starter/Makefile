ENV=development
VERSION=latest
LOCAL_PORT=9000

docker.build:
	docker build -t oliversavio/api:$(VERSION) .

docker.run:
	docker run --name api_test -d \
		-e FOO_ENV=$(ENV) \
		-p $(LOCAL_PORT):80 \
	 	--log-driver local \
		--log-opt max-size=10m \
		--log-opt max-file=5 \
		oliversavio/api:$(VERSION) 

docker.clean:
	docker container rm api_test

docker.stop:
	docker stop api_test

swag:
	swag init

.PHONY build: swag docker.build

.PHONY run: docker.clean docker.run

