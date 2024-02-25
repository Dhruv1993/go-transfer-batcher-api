package microBatchService

import (
	"fmt"
	gomicrobatcher "github.com/Dhruv1993/go-microbatcher"
	"github.com/go-redis/redis/v8"
	"time"
)

type Domain struct {
	// we can put any shared service here as the size of services grows
	db *redis.Client
}

func NewDomain(redisClient *redis.Client) *Domain {
	return &Domain{
		db: redisClient,
	}
}

func (d Domain) processTransfers(job Transfer) (Transfer, error) {
	// processing done here
	// Initialise the batcher
	batcher := gomicrobatcher.NewBatcher(&TransactionProcessor{RedisClient: d.db})

	// set the batch-size and frequency
	batcher.SetBatchSize(2)
	batcher.SetFrequency(20 * time.Millisecond)

	// submit the job to the batching system
	res, err := batcher.AddJob(job)

	if err != nil {
		// return empty
		return Transfer{}, err
	}
	// Create a channel for receiving the result
	resultCh := make(chan interface{}, 1)
	go func() {
		// Retrieve the result from the JobResult
		result := res.GetValue()

		// Send the result through a channel to the main goroutine
		resultCh <- result
	}()

	// Wait for the result from the goroutine
	processedResult := <-resultCh

	// Type assertion to convert result to TransferResponse
	transferResult, ok := processedResult.(Transfer)
	if !ok {
		fmt.Println("Type assertion failed")
		return Transfer{}, nil
	}

	return Transfer{
		FromAccount: transferResult.FromAccount,
		ToAccount:   transferResult.ToAccount,
		Amount:      transferResult.Amount,
	}, nil
}
