package consumer

import (
	blockingQueue "CYOLO/blocking_queue"
	"CYOLO/data_holder"
	"fmt"
	"log"
	"sync"
	"time"
)

type Consumer struct {
	Queue      *blockingQueue.BlockingQueue
	DataHolder *data_holder.DataHolder
	mu         sync.Mutex
	errChan    chan error
	wg         *sync.WaitGroup
}

func NewConsumer(queue *blockingQueue.BlockingQueue, dataHolder *data_holder.DataHolder,
	c chan error, wg *sync.WaitGroup) *Consumer {
	return &Consumer{
		Queue:      queue,
		DataHolder: dataHolder,
		errChan:    c,
		wg:         wg,
	}
}
func (c *Consumer) StartConsuming() {
	defer func() {
		log.Fatal("Consumer got un expected error... exiting")
		c.wg.Done()
		c.errChan <- fmt.Errorf("unexpected exit")
	}()

	for {
		c.mu.Lock()
		word := c.Queue.Dequeue()
		log.Printf(fmt.Sprintf("Consumer consumed the word %s", word))

		if word != "" {
			c.DataHolder.AddWord(word)
		} else {
			log.Printf("Consumer is sleeping for 2 seconds")
			time.Sleep(2 * time.Second)
		}

		c.mu.Unlock()
	}
}
