proto:
	# protoc --go-grpc_out=proto proto/whoami.proto
	# protoc --go_out=plugins=grpc:whoami proto/whoami.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/whoami.proto
