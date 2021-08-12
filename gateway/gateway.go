package gateway

import (
	"context"
	"io"
	"net/http"
)

func HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"success": "yay"}`)
	}
}

func New(ctx context.Context) (http.Handler, error) {
	mux := http.NewServeMux()
	mux.Handle("/", HomePage())
	return mux, nil
}
