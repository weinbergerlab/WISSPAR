package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gitlab.com/weinbergerlab/immunogenicity-data-project/app"

	"github.com/go-chi/chi"

	"github.com/go-chi/valve"
)

// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func healthCheck(r chi.Router) {
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		codeParams, ok := r.URL.Query()["code"]
		if ok && len(codeParams) > 0 {
			statusCode, _ := strconv.Atoi(codeParams[0])
			if statusCode >= 200 && statusCode < 600 {
				w.WriteHeader(statusCode)
			}
		}
		requestID := uuid.Must(uuid.NewUUID())
		fmt.Fprintf(w, requestID.String())
	})
}

func main() {
	// Our graceful valve shut-off package to manage code preemption and
	// shutdown signaling.
	vv := valve.New()
	baseCtx := vv.Context()

	cfg, err := app.LoadConfig("app")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("creating router")
	r := app.Router(baseCtx, cfg)
	workDir, _ := os.Getwd()
	public := http.Dir(filepath.Join(workDir, "./", "public", "assets"))
	fileServer(r, "/static", public)
	healthCheck(r)

	srv := &http.Server{
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSecs) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSecs) * time.Second,
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		Handler:      r,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Println("shutting down..")

			// first valv
			vv.Shutdown(20 * time.Second)

			// create context with timeout
			ctx, cancel := context.WithTimeout(baseCtx, 20*time.Second)
			defer cancel()

			// start http shutdown
			srv.Shutdown(ctx)

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()

	fmt.Println("http server is listening...")
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
