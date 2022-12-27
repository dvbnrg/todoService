# List all recipes (you just called this)
default: dev
# run a local dev build
dev: build run
# Build Macro
build: protoc 

# Run all unit tests
test:
    echo "test called"
# Run Local versions of binaries
run:
    sleep 2
    goreman start
    echo "run called"
# Build protoc artifacts
protoc:
    protoc --go-grpc_out=pb --go_out=pb pb/todoService.proto

# Clean all build artifacts
clean:
    rm pb/*.pb.go
    rm bin/*