package main

import (
	"context"
	"fmt"
	"net/http"
	"test-proto-grpc/gateway"
)

func init() {
	fmt.Println("krakend-grpc-post plugin loaded!!!")
}

var ClientRegisterer = registerer("grpc-post")

type registerer string

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), func(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
		return gateway.New(ctx)
	})
}

func main() {}
