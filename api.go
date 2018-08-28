package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var gameserverRedis *redis.Client

func main() {
	corsObj := handlers.AllowedOrigins([]string{"*"})
	router := mux.NewRouter()
	router.HandleFunc("/games", Games).Methods("GET")
	gameserverRedis = connectToRedis("redis-gameservers:6379")
	log.Fatal(http.ListenAndServe(":6001", handlers.CORS(corsObj)(router)))
}

func connectToRedis(addr string) *redis.Client {
	var client *redis.Client
	for {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
		})
		_, err := client.Ping().Result()
		if err != nil {
			fmt.Println("Could not connect to redis")
			fmt.Println(err)
		} else {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Connected to redis")
	return client
}

func Games(w http.ResponseWriter, r *http.Request) {
	keys, _ := gameserverRedis.Keys("*").Result()
	json.NewEncoder(w).Encode(keys)
}
