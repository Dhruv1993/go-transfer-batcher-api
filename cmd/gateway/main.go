package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	app "github.com/Dhruv1993/app"
	"github.com/Dhruv1993/app/services"
	"github.com/Dhruv1993/app/services/gateway"
	"github.com/Dhruv1993/app/services/microBatchService"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

// NewRedisClient creates a new instance of the Redis client.
func NewRedisClient() *Client {
	return &Client{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "secret", 
			DB:       0,        // Default DB
		}),
	}
}

func main() {
	//initilize db
	dbConn := NewRedisClient()
	// initialise all the services here
	svcs := initServices(dbConn.client)

	g := gateway.New(svcs, "/api/v1") // prefix should probably should come form a config file

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	app.RunInBackground(func() {
		defer cleanup()
		sig := <-sc
		log.Panicf("Caught sig: %+v", sig)
	})



	log.Printf("Listening at: http://%s", ":5001")
	if err := http.ListenAndServe(":5001", g.Handler); err != nil {
		log.Fatal(err)
	}

}

func cleanup() {
	// can do db connection closing here
	log.Printf("Service exited")
	os.Exit(0)
}

func initServices(dbConn *redis.Client) []services.Service {
	// here we can initilise all the services, dbconnections by providing them with JWT, configs etc
	log.Println("-----------microbatching test service started-------")

	microbatchTestDomain := microBatchService.NewDomain(dbConn)
	microbatchTestService := microBatchService.Service{}

	microbatchTestService.API = microbatchTestDomain

	return []services.Service{
		microbatchTestService,
		// add other services here (patrol, location, asset, incident, etc)
	}
}
