package gateway

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	helloworldpb "test-proto-grpc/proto"
	"time"
)

const portServerHello = ":50051"

func HelloPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := grpc.Dial(portServerHello, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect to service: %v", err)
		}
		defer conn.Close()
		helloServiceClient := helloworldpb.NewGreeterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, err := helloServiceClient.SayHello(ctx, &helloworldpb.HelloRequest{
			Name: "Hi World",
		})
		if err != nil {
			log.Fatalf("Could not create request: %v", err)
		}
		fmt.Println(response.Message)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, response.Message)
	}
}

func HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"success": "yay"}`)
	}
}

func New(ctx context.Context) (http.Handler, error) {
	mux := http.NewServeMux()
	mux.Handle("/test", HelloPage())
	mux.Handle("/", HomePage())
	return mux, nil
}
