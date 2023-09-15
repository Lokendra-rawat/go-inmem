# run sh protoc-gen.sh in wall directory
# sudo apt install golang-goprotobuf-dev
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc --proto_path=proto \
    --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
    feature_store.proto

protoc --proto_path=proto \
    --go_out=proto --go_opt=paths=source_relative \
    feature_store.proto
