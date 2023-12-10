package main

import (
	blockingQueue "CYOLO/blocking_queue"
	"CYOLO/config"
	"CYOLO/consumer"
	"CYOLO/data_holder"
	"CYOLO/server"
	"log"
	"sync"
)

func main() {
	var serverConfig *config.Config
	var err error
	log.Printf("Starting app...")

	if err, serverConfig = config.ReadAndValidateConfig("config", "json", "."); err != nil {
		log.Printf("Error reading configuration: %v\n", err)
		return
	}
	log.Printf("Valid configurations")

	wg := sync.WaitGroup{}
	wg.Add(1)

	errChan := make(chan error)
	q := blockingQueue.NewQueue()
	dh := data_holder.NewDataHolder()

	c := consumer.NewConsumer(q, dh, errChan, &wg)
	go func() {
		err := <-errChan
		if err != nil {
			log.Printf("Consumer got un expected error, starting a new one")
			c = consumer.NewConsumer(q, dh, errChan, &wg)
			go c.StartConsuming()
		}
	}()
	go c.StartConsuming()

	log.Printf("Bringing up server...")

	server.NewServer(q, dh, serverConfig).Run()

	wg.Wait()
}
