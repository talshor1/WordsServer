package blocking_queue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"

	"CYOLO/blocking_queue"
)

func TestBlockingQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BlockingQueue Suite")
}

var _ = Describe("BlockingQueue", func() {
	var (
		queue *blocking_queue.BlockingQueue
	)

	BeforeEach(func() {
		queue = blocking_queue.NewQueue()
	})

	Describe("Enqueue", func() {
		Context("with a single word", func() {
			It("should enqueue the word", func() {
				word := "apple"
				queue.Enqueue(word)
				Expect(queue.Dequeue()).To(Equal(word))
			})
		})

		Context("with multiple words", func() {
			It("should enqueue each word separately", func() {
				words := "apple,banana,orange"
				expectedWords := []string{"apple", "banana", "orange"}

				queue.Enqueue(words)

				for _, expectedWord := range expectedWords {
					Expect(queue.Dequeue()).To(Equal(expectedWord))
				}
			})
		})
	})

	Describe("Dequeue", func() {
		Context("when the queue is not empty", func() {
			It("should return the front item", func() {
				queue.Enqueue("apple")
				Expect(queue.Dequeue()).To(Equal("apple"))
			})
		})

		Context("when the queue is empty", func() {
			It("should return an empty string", func() {
				Expect(queue.Dequeue()).To(Equal(""))
			})
		})
	})

	Describe("IsSingleWord", func() {
		It("should return true for a single word", func() {
			Expect(blocking_queue.IsSingleWord("apple")).To(BeTrue())
		})

		It("should return false for multiple words", func() {
			Expect(blocking_queue.IsSingleWord("apple,banana,orange")).To(BeFalse())
		})
	})
})
