package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

var sub = &common.Subscription{
	PubsubName: "redis-pubsub",
	Topic:      "constraint-violations",
	Route:      "/violations",
}

func main() {
	appPort := "6001"
	if value, ok := os.LookupEnv("APP_PORT"); ok {
		appPort = value
	}

	s := daprd.NewService(":" + appPort)
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	fmt.Printf("Subscriber received: %v at time: %s\n", e.Data, time.Now().String())
	return false, nil
}
