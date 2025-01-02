package resthttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/resthttp/dto"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideHTTPServer(
	lc fx.Lifecycle,
	config *raas.Config,
	router http.Handler,
	logger *zap.Logger,
) (*http.Server, error) {
	timeoutErrMeta := &dto.Meta{
		Code:    0,
		Message: fmt.Sprintf("handler timeout: %d ns", config.Server.HandlerTimeout.Nanoseconds()),
	}
	timeoutErr := dto.ResponseError{Meta: timeoutErrMeta}
	timeoutMsg, err := json.Marshal(timeoutErr)
	if err != nil {
		return nil, fmt.Errorf("problem occurred while marshalling handler timeout message: %w", err)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      http.TimeoutHandler(router, config.Server.HandlerTimeout, string(timeoutMsg)),
	}
	lc.Append(fx.StartHook(func(_ context.Context) error {
		ln, err := net.Listen("tcp", srv.Addr)
		if err != nil {
			return err
		}
		logger.Info("api server started", zap.Int("port", config.Server.Port))

		//nolint:errcheck
		go srv.Serve(ln)

		return nil
	}))
	lc.Append(fx.StopHook(func(ctx context.Context) error {
		logger.Info("api server stopped")
		return srv.Shutdown(ctx)
	}))

	lc.Append(fx.StartHook(func(_ context.Context) error {
		debugSrv := &http.Server{Addr: fmt.Sprintf(":%d", config.Server.DebugPort), Handler: http.DefaultServeMux}

		ln, err := net.Listen("tcp", debugSrv.Addr)
		if err != nil {
			return err
		}
		logger.Info("debug server started", zap.Int("port", config.Server.DebugPort))

		// http.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		// 	if !appProbesService.Ready(r.Context()) {
		// 		w.WriteHeader(http.StatusServiceUnavailable)
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// })
		// http.HandleFunc("/health/live", func(w http.ResponseWriter, _ *http.Request) {
		// 	if !appProbesService.Live() {
		// 		w.WriteHeader(http.StatusServiceUnavailable)
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// })
		// http.Handle("/logLevel", atomicLevel)

		//nolint:errcheck
		go debugSrv.Serve(ln)

		return nil
	}))

	return srv, nil
}

var Providers = fx.Provide(
	MakeRoutes,
	ProvideHTTPServer,
)

var Launcher = fx.Invoke(func(*http.Server) {})
