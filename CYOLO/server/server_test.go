package server_test

import (
	blockingQueue "CYOLO/blocking_queue"
	"CYOLO/config"
	"CYOLO/data_holder"
	"CYOLO/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = Describe("Server", func() {
	var (
		queue      *blockingQueue.BlockingQueue
		dataHolder *data_holder.DataHolder
		testServer *server.Server
	)

	BeforeEach(func() {
		queue = blockingQueue.NewQueue()
		dataHolder = data_holder.NewDataHolder()
		testServer = server.NewServer(queue, dataHolder, &config.Config{})
	})

	Describe("POST /words", func() {
		Context("with valid input", func() {
			It("should return 200", func() {
				input := "apple,banana,orange"
				request, _ := http.NewRequest(http.MethodPost, "/words?words="+input, nil)
				response := httptest.NewRecorder()

				testServer.Router.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(http.StatusOK))
			})
		})

		Context("with invalid input", func() {
			It("should return a 400 Bad Request", func() {
				input := ""
				request, _ := http.NewRequest(http.MethodPost, "/words?words="+input, nil)
				response := httptest.NewRecorder()

				testServer.Router.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(http.StatusBadRequest))
				Expect(response.Body.String()).To(ContainSubstring("Invalid comma-separated string format"))
			})
		})
	})
})
