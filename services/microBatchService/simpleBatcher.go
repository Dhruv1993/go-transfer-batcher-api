package microBatchService

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Dhruv1993/app/utils"
	"github.com/go-redis/redis/v8"
)

type TransactionProcessor struct {
	// You can add any dependencies or configurations for transaction processing
	RedisClient *redis.Client
}

func (s *TransactionProcessor) Process(jobs []interface{}) []interface{} {
	// Implement the logic to process jobs in batches
	// For simplicity, this example just logs the jobs and returns them as results
	transferID := utils.NewUUID()
	//Marshal the slice to JSON
	jsonData, error := json.Marshal(jobs)
	if error != nil {
		log.Fatal(error)
	}
	err := s.RedisClient.Set(context.Background(), transferID, jsonData, 0).Err()
	if err != nil {
		fmt.Println("errorrrrr setting the data")
		log.Fatal(err)
	}

	//
	// Retrieve the payload from Redis using the UUID as the key
	payloadFromRedis, err := s.RedisClient.Get(context.Background(), transferID).Result()
	//payloadFromRedis, err := s.RedisClient.Get(context.Background(), transferID.String())
	if err != nil {
		log.Fatal(err)
	}

	//// Unmarshal JSON data
	var storedData []Transfer
	err = json.Unmarshal([]byte(payloadFromRedis), &storedData)
	var interfacesSlice []interface{}
	for _, transaction := range storedData {
		interfacesSlice = append(interfacesSlice, transaction)
	}

	return interfacesSlice
}
