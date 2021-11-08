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
	RedisClient *redis.Client
}

type ExpectedBody struct {
	Key   string
	Value string
}

type ErrorResponse struct{Error string}

func NewServer(redisClient *redis.Client) *MyCacheServer {
	server := new(MyCacheServer)

	router := http.NewServeMux()
	router.Handle("/get", http.HandlerFunc(server.getValueHandler))
	router.Handle("/set", http.HandlerFunc(server.setValueHandler))

	server.Handler = router
	server.RedisClient = redisClient
	return server
}

func (s *MyCacheServer) getValueHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var ctx = context.Background()
	rdb     := s.RedisClient
	encoder := json.NewEncoder(w)

	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			errResponse := ErrorResponse{"Param 'key' is missing"}
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(errResponse)

			return
	}

	key := keys[0]
	value, err := rdb.Get(ctx, key).Result()

	if err != nil {
		var errorMessage string

		if err == redis.Nil {
			errorMessage = "Key Not Found"
		  w.WriteHeader(http.StatusNotFound)
			} else {
			errorMessage = "Error fetching value"
			w.WriteHeader(http.StatusBadRequest)
		}

		errorResponse := ErrorResponse{errorMessage}
		encoder.Encode(errorResponse)
		return
	}

	response := ExpectedBody{key, value}
	encoder.Encode(response)
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
	rdb     := s.RedisClient
	encoder := json.NewEncoder(w)

	var b ExpectedBody

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Err: ", err)
		return
	}

	err = rdb.Set(ctx, b.Key, b.Value, 5*60*time.Second).Err()

	if err != nil {
		log.Fatal(err)
		errorResponse := ErrorResponse{"Error fetching value"}
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(errorResponse)

		return
	}

	encoder.Encode(b)
}
