package main

import (
	"go-rate-limit/api"
	"go-rate-limit/errs"
	"go-rate-limit/redis"
	"log"
	"net/http"
)

func initRedis() error {
	client := redis.CreateClient()
	_, err := client.Ping().Result()

	if err != nil {
		return errs.RedisConnectionErr
    }
    
    return nil
}

func main() {
    err := initRedis()
    if err != nil {
        log.Print(err.Error())
        return
    }

	http.HandleFunc("/docker-name", api.GetDockerName)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
