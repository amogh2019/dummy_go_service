

all: greet

greet: buildGreetProto

buildGreetProto : buildCalculatorProto
	protoc -Igreet/proto --go_opt=module=github.com/amogh2019/dummy_go_service --go_out=. --go-grpc_opt=module=github.com/amogh2019/dummy_go_service --go-grpc_out=.   greet/proto/*.proto


buildCalculatorProto :
	# protoc -Icalculator/proto --go_opt=module=github.com/amogh2019/dummy_go_service --go_out=. --go-grpc_opt=module=github.com/amogh2019/dummy_go_service --go-grpc_out=.   calculator/proto/*.proto
	protoc --go_opt=module=github.com/amogh2019/dummy_go_service --go_out=. --go-grpc_opt=module=github.com/amogh2019/dummy_go_service --go-grpc_out=.   calculator/proto/*.proto


