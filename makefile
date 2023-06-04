build-grpc:
	docker build -f grpcserver_dockerfile . -t grpc_server:v1.1
build-search:
	docker build -f search_dockerfile . -t search_app:v1.1
rm-tmp:
	docker rmi  $(docker images|grep none|awk '{print $3}')
build-all: build-grpc build-search 
	@echo "生成grpcs_server容器、search_app容器成功"
run-grpc:
	@echo $(shell docker run  --network=host --name grpc_server -d grpc_server:v1)
run-search:
	@echo $(shell docker run  --network=host --name search_app -d search_app:v1)
rm-container:
	docker stop grpc_server && docker rm grpc_server
	docker stop search_app  && docker rm search_app 
save:
	docker save -o ./deploy/grpc_server.tar  grpc_server:v1.1
	docker save -o ./deploy/search_app.tar  search_app:v1.1
push:
	scp ./deploy/grpc_server.tar ./deploy/search_app.tar  root@192.168.56.102:/root 
	scp ./deploy/grpc_server.tar ./deploy/search_app.tar  root@192.168.56.103:/root 
load:
	ssh root@192.168.56.102 && docker load -i grpc_server.tar && docker load -i search_app.tar
	ssh root@192.168.56.103 && docker load -i grpc_server.tar && docker load -i search_app.tar