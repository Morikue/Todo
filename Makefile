up:
	docker-compose -f docker-compose.yaml down -v
	docker-compose -f docker-compose.yaml up -d postgres
	docker-compose -f docker-compose.yaml up --build migrate-users
	docker-compose -f docker-compose.yaml up --build migrate-todo
	docker-compose -f docker-compose.yaml up -d users-service todo-service
	docker-compose -f docker-compose.yaml up -d gateway-service
	docker-compose -f docker-compose.yaml ps

down:
	docker-compose -f docker-compose.yaml down -v

generate-users:
	protoc -I api/protos/users \
		--go_out=api/protos/users \
		--go_opt=paths=source_relative \
		--go-grpc_out=api/protos/users \
		--go-grpc_opt=paths=source_relative \
		api/protos/users/users.proto

	mkdir -p users/pkg/grpc_stubs/users
	cp -r api/protos/users/* users/pkg/grpc_stubs/users

	mkdir -p gateway/pkg/grpc_stubs/users
	cp -r api/protos/users/* gateway/pkg/grpc_stubs/users

generate-todos:
	protoc -I api/protos/todos \
		--go_out=api/protos/todos \
		--go_opt=paths=source_relative \
		--go-grpc_out=api/protos/todos \
		--go-grpc_opt=paths=source_relative \
		api/protos/todos/todos.proto

	mkdir -p todo/pkg/grpc_stubs/todos
	cp -r api/protos/todos/* todo/pkg/grpc_stubs/todos

	mkdir -p gateway/pkg/grpc_stubs/todos
	cp -r api/protos/todos/* gateway/pkg/grpc_stubs/todos
