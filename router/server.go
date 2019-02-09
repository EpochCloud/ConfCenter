package router

import (
	"net/http"
	"time"
	"sync"
	"syscall"
	"os/signal"
	"os"
	"context"
	"ConfCenter/config"
	"ConfCenter/log"
)

func Run(){
	wg := &sync.WaitGroup{}
	exit := make(chan os.Signal)

	mux := http.NewServeMux()
	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:          "127.0.0.1:8080",
		Handler:      mux,
	}
	Router(mux)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go shutdown(exit,wg,srv)

	log.Debug("127.0.0.1:8080")
	config.Log.Debug("127.0.0.1:8080")

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	log.Warn("waiting for the remaining connections to finish...")
	config.Log.Warn("waiting for the remaining connections to finish...")
	wg.Wait()
	config.Log.Warn("gracefully shutdown the http server...")
	log.Warn("gracefully shutdown the http server...")
}


func shutdown(exit chan os.Signal,wg *sync.WaitGroup,srv *http.Server){
		<-exit
		wg.Add(1)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer func(){
			cancel()
			wg.Done()
		}()
		err := srv.Shutdown(ctx)
		if err != nil {
			config.Log.Error("http shutdown err",err)
			return
		}
}