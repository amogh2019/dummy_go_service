

all: greet

greet: buildGreetProto

buildGreetProto :
	protoc -Igreet/proto --go_opt=module=github.com/amogh2019/dummy_go_service --go_out=. --go-grpc_opt=module=github.com/amogh2019/dummy_go_service --go-grpc_out=.   greet/proto/*.proto

