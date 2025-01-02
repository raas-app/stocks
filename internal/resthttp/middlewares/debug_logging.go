package middlewares

import (
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type DebugLogger struct {
	logger *zap.Logger
}
type DebugResponseWriter struct {
	w            http.ResponseWriter
	statusCode   int
	responseBody []byte
}

func NewDebugLogger(logger *zap.Logger) (*DebugLogger, error) {
	if logger == nil {
		return nil, errors.New("debug logger: provided logger is nil")
	}
	return &DebugLogger{
		logger: logger,
	}, nil
}

func NewDebugResponseWriter(w http.ResponseWriter) *DebugResponseWriter {
	return &DebugResponseWriter{
		w:            w,
		statusCode:   0,
		responseBody: []byte{},
	}
}

func (d *DebugLogger) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rw := NewDebugResponseWriter(w)
		next.ServeHTTP(rw, r)

		// ctx := r.Context()
		// userJWT := utils.GetUserFromContext(ctx)
		fields := []zap.Field{
			zap.Int("status_code", rw.statusCode),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("http_method", r.Method),
			zap.String("request_uri", r.RequestURI),
			zap.Duration("request_time", time.Since(startTime)),
		}
		// if userJWT != nil {
		// 	fields = append(fields, zap.Int64("user_id", userJWT.ID))
		// }
		if rw.statusCode >= 300 {
			fields = append(fields, zap.ByteString("response_body", rw.responseBody))
		}

		d.logger.Debug("http incoming request", fields...)
	})
}

func (d *DebugResponseWriter) Header() http.Header {
	return d.w.Header()
}
func (d *DebugResponseWriter) Write(b []byte) (int, error) {
	d.responseBody = b
	return d.w.Write(b)
}
func (d *DebugResponseWriter) WriteHeader(statusCode int) {
	d.w.WriteHeader(statusCode)
	d.statusCode = statusCode
}
