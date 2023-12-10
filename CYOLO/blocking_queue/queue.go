package blocking_queue

import (
	"strings"
	"sync"
)

type BlockingQueue struct {
	items []string
	mu    sync.Mutex
}

func NewQueue() *BlockingQueue {
	return &BlockingQueue{
		items: make([]string, 0),
	}
}

func (q *BlockingQueue) Enqueue(words string) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if IsSingleWord(words) {
		q.items = append(q.items, words)
		return
	}

	for _, word := range strings.Split(words, ",") {
		q.items = append(q.items, word)
	}
}

func (q *BlockingQueue) Dequeue() string {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return ""
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func IsSingleWord(s string) bool {
	return !strings.Contains(s, ",")
}
