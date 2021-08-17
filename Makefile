GCC:=x86_64-w64-mingw32-gcc
GXX:=x86_64-w64-mingw32-g++

build-wsl-for-win:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=$(GCC) CXX=$(GXX) go build -x ./

pb: 
	protoc \
		--go-grpc_out=./proto/.  --go-grpc_opt=paths=source_relative \
		--go_out=./proto/.  --go_opt=paths=source_relative \
 		-I. event.proto

clean:
	rm -rf ./pb 