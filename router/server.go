package router

import (
	"ConfCenter/config"
	"ConfCenter/log"
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run() {
	wg := &sync.WaitGroup{}
	exit := make(chan os.Signal)
	mux := http.NewServeMux()
	srv := &http.Server{
		ReadTimeout:  time.Duration(config.Conf.HttpConf.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Conf.HttpConf.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Conf.HttpConf.IdleTimeout) * time.Second,
		Addr:         config.Conf.HttpConf.Addr,
		Handler:      mux,
	}
	Router(mux)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go shutdown(exit, wg, srv)
	log.Debug(config.Conf.HttpConf.Addr)
	config.Log.Debug("[%v] http run %v",time.Now(), config.Conf.HttpConf.Addr)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	log.Warn("waiting for the remaining connections to finish...")
	config.Log.Warn("[%v] waiting for the remaining connections to finish...",time.Now())
	wg.Wait()
	close(exit)
	config.Log.Warn("[%v] gracefully shutdown the http server...",time.Now())
	log.Warn("gracefully shutdown the http server...")
}

func shutdown(exit chan os.Signal, wg *sync.WaitGroup, srv *http.Server) {
	<-exit
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancel()
		wg.Done()
	}()
	err := srv.Shutdown(ctx)
	if err != nil {
		config.Log.Error("[%v] http shutdown err",time.Now(), err)
		return
	}
}
