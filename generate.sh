protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. proto/helloworld.proto

go build -buildmode=plugin -o grpc-post.so ./plugin
