package server

import (
	"net/http"
	"encoding/json"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type MyCacheServer struct {
	http.Handler
}

func NewServer() *MyCacheServer {
	server := new(MyCacheServer)

	router := http.NewServeMux()
	router.Handle("/get", http.HandlerFunc(server.getValueHandler))
	router.Handle("/set", http.HandlerFunc(server.setValueHandler))

	server.Handler = router
	return server
}

func (s *MyCacheServer) getValueHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
  		Addr:     "localhost:6379",
  		Password: "",
  		DB:       0,
	})

	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			w.WriteHeader(http.StatusBadRequest)

			return
	}

	key := keys[0]
	value, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		value = "Err: Key Not Found"
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(value)
}

type ExpectedBody struct {
	Key   string
	Value string
}


func (s *MyCacheServer) setValueHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		log.Println("Wrong method sent")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var b ExpectedBody

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = rdb.Set(ctx, b.Key, b.Value, 10*time.Second).Err()

	if err != nil {
		log.Fatal(err)
	}
}
