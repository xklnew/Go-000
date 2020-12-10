package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)
//同时关闭
func serverApp(stop <- chan struct{}) error  {
	return serve("127.0.0.1:8080",http.DefaultServeMux,stop)
}

func serverDebug(stop <- chan struct{}) error {
	return serve("127.0.0.1:8001",http.DefaultServeMux,stop)
}

func serve(addr string,handler http.Handler,stop <- chan struct{}) error {
	s := http.Server{
		Addr: addr,
		Handler: handler,
	}
	go func() {
		<-stop
		fmt.Println("I am closed addr=%v",addr)
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func main()  {
	done := make(chan error,2)
	stop := make(chan struct{})//有一个有问题关闭所有的都关闭
	go func() {
		 serverDebug(stop)
	}()
	go func() {
		serverApp(stop)
	}()
	close(stop)
	time.Sleep(time.Second * 10)
	var stopped bool
	for i:=0;i<cap(done);i++ {
		if err :=<-done;err != nil{
			fmt.Println("error: %v",err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
