package main

import (
	"log"
	"sync"

	"github.com/bhagas/go-svc-echo/config"
	"github.com/bhagas/go-svc-echo/src"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Printf(".env is not loaded properly")
	}

	Cfg := config.NewConfig()
	Srv := src.InitServer(Cfg)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		Srv.Run()
	}()

	wg.Wait()
}
