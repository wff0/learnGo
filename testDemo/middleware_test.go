package testDemo

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 执行handler之前的逻辑
		next.ServeHTTP(w, r)
		// 执行完毕handler后的逻辑
	})
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func TestMiddleware(t *testing.T) {
	http.Handle("/", loggingHandler(http.HandlerFunc(index)))

	http.ListenAndServe(":8000", nil)
}

func TestStatic(t *testing.T) {
	http.Handle("/", http.FileServer(http.Dir("D:/html/static/")))
	http.ListenAndServe(":8080", nil)
}
