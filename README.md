# gophkeeper

protoc --go_out=server/ --go_opt=paths=import --go-grpc_out=server/ --go-grpc_opt=paths=import proto/gophkeeper.proto

protoc --go_out=client/ --go_opt=paths=import --go-grpc_out=client/ --go-grpc_opt=paths=import proto/gophkeeper.proto