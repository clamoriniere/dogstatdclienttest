package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	log.SetOutput(os.Stdout)
	rand.Seed(42)

	client, err := statsd.New("unix:///var/run/datadog/dsd.socket",
		statsd.Buffered(),                        // enable buffering
		statsd.WithMaxMessagesPerPayload(16),     // sets the maximum number of messages in a single datagram
		statsd.WithNamespace("testmetrics."),     // prefix every metric with the app name
		statsd.WithTags([]string{"source:test"}), // send the EC2 availability zone as a tag with every metric
		// add more options here...
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(200 * time.Millisecond)
		log.Printf("send distribution: %v", time.Now())
		if err = client.Distribution("dist", float64(rand.Intn(100)), []string{}, 1); err != nil {
			log.Printf("Error: %v", err)
		}
	}

}
