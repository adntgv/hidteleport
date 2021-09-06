GCC:=x86_64-w64-mingw32-gcc
GXX:=x86_64-w64-mingw32-g++
SRC:=./
OUT:=hidteleport

all: build-wsl-for-win build-wsl-for-lin

build-wsl-for-win:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=$(GCC) CXX=$(GXX) go build -o $(OUT).exe -x $(SRC)

build-wsl-for-lin:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(OUT) -x $(SRC)

pb: 
	protoc \
		--go-grpc_out=./proto/.  --go-grpc_opt=paths=source_relative \
		--go_out=./proto/.  --go_opt=paths=source_relative \
 		-I. event.proto

clean:
	rm  $(OUT) $(OUT).exe

lint:
	golangci-lint run

scp:
	scp $(OUT) aid@192.168.0.192:/home/aid/Downloads/hidteleport