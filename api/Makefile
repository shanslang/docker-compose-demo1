prod_name = api
version = v1.2.0
container_name = api:${version}-beta
container_name_online = api:${version}

build: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o release/bin/$(prod_name) main.go
	docker buildx build -t $(container_name) -f Dockerfile --platform=linux/amd64 .
	docker tag $(container_name) beta/$(container_name)
	#docker push 镜像仓库地址/zta/beta/$(container_name) 将本地镜像推到镜像仓库

build-prod: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o release/bin/$(prod_name) cmd/main.go
	docker buildx build -t $(container_name_online) -f Dockerfile --platform=linux/amd64 .
	docker tag $(container_name_online) zta-harbor.moresec.cc:4443/zta/$(container_name_online)
	docker push 192.168.180.195/zta/$(container_name_online)

clean:
	rm -rf release/bin/$(prod_name)