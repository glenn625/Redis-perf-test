package main

import (
	"context"
	"fmt"
	"time"
	"math/rand"
	"unsafe"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
)

var (
	PUBSUB_NAME  = "redis-pubsub"
	PUBSUB_TOPIC = "constraint-violations"
	NUM_TEST_VIOLATION_MESSAGES = 500000
)

// Gatekeeper constraint status violation details.
type DetailedStatusViolation struct {
	Message              string
	Details              interface{}
	ConstraintGroup      string
	ConstraintAPIVersion string
	ConstraintKind       string
	ConstraintName       string
	ConstraintNamespace  string
	ConstraintAction     string
	ResourceGroup        string
	ResourceAPIVersion   string
	ResourceKind         string
	ResourceNamespace    string
	ResourceName         string
}

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	ctx := context.Background()

	fmt.Printf("Starting publishing at: %s\n", time.Now().String())

	// Construct a violation message.
	rand.Seed(time.Now().UnixNano())
	message := "Mock Gatekeeper constraint violation" + strconv.Itoa(rand.Intn(100))
	violation := DetailedStatusViolation{
		Message: message,
	}

	for i := 1; i <= NUM_TEST_VIOLATION_MESSAGES; i++ {
		// Publish an event using Dapr pub/sub
		if err := client.PublishEvent(ctx, PUBSUB_NAME, PUBSUB_TOPIC, violation); err != nil {
			panic(err)
		}

		fmt.Printf("Published data: %v size: %d at time: %s\n", violation, unsafe.Sizeof(violation), time.Now().String())
	}

	time.Sleep(360) // wait 6 minutes before starting again.

	fmt.Printf("Done publishing data!\n\n\n\n\n")
}
