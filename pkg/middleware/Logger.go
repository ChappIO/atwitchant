package middleware

import (
	"log"
	"net/http"
)

type Logger struct {
}

type statusInterceptingWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusInterceptingWriter)  WriteHeader(statusCode int)  {
	if w.status == 0 {
		w.status = statusCode
	}
	w.ResponseWriter.WriteHeader(statusCode)
}

func (l *Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		newWriter := &statusInterceptingWriter{ResponseWriter: writer, status: 200}

		defer func(writer *statusInterceptingWriter) {
			log.Printf("HTTP %s %s - %d", request.Method, request.URL, newWriter.status)
		}(newWriter)

		next.ServeHTTP(newWriter, request)
	})
}
