protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. proto/helloworld.proto

go build -buildmode=plugin -o plugin/grpc-post.so ./plugin

#export PATH="$PATH:$(go env GOBIN)"
