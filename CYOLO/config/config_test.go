package config_test

import (
	"CYOLO/config"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"os"
	"testing"
)

var _ = Describe("Config", func() {
	Context("Valid configuration", func() {
		It("should not return an error", func() {
			err, _ := config.ReadAndValidateConfig("config", "json", "../.")
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("Invalid address in configuration", func() {
		BeforeEach(func() {
			configData := map[string]interface{}{
				"server": map[string]interface{}{
					"address": "invalid address",
					"port":    "8080",
				},
			}

			jsonData, err := json.MarshalIndent(configData, "", "  ")
			Expect(err).To(BeNil())
			filePath := "invalid_address_config.json"
			file, err := os.Create(filePath)
			Expect(err).To(BeNil())
			defer file.Close()

			_, err = io.WriteString(file, string(jsonData))
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			os.Remove("invalid_address_config.json")
		})

		It("should return an error", func() {
			err, _ := config.ReadAndValidateConfig("invalid_address_config", "json", ".")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("server addrress is not valid"))
		})
	})

	Context("Invalid port in configuration", func() {
		BeforeEach(func() {
			configData := map[string]interface{}{
				"server": map[string]interface{}{
					"address": "127.0.0.1",
					"port":    "804480",
				},
			}

			jsonData, err := json.MarshalIndent(configData, "", "  ")
			Expect(err).To(BeNil())
			filePath := "invalid_port_config.json"
			file, err := os.Create(filePath)
			Expect(err).To(BeNil())
			defer file.Close()

			_, err = io.WriteString(file, string(jsonData))
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			os.Remove("invalid_port_config.json")
		})

		It("should return an error", func() {
			err, _ := config.ReadAndValidateConfig("invalid_port_config", "json", ".")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("server port is not valid"))
		})
	})

	Context("Missing port in configuration", func() {
		BeforeEach(func() {
			configData := map[string]interface{}{
				"server": map[string]interface{}{
					"address": "127.0.0.1",
				},
			}

			jsonData, err := json.MarshalIndent(configData, "", "  ")
			Expect(err).To(BeNil())
			filePath := "invalid_port_config.json"
			file, err := os.Create(filePath)
			Expect(err).To(BeNil())
			defer file.Close()

			_, err = io.WriteString(file, string(jsonData))
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			os.Remove("invalid_port_config.json")
		})

		It("should return an error", func() {
			err, _ := config.ReadAndValidateConfig("invalid_port_config", "json", ".")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("missing required fields in server configuration"))
		})
	})

	Context("Missing address in configuration", func() {
		BeforeEach(func() {
			configData := map[string]interface{}{
				"server": map[string]interface{}{
					"address": "127.0.0.1",
				},
			}

			jsonData, err := json.MarshalIndent(configData, "", "  ")
			Expect(err).To(BeNil())
			filePath := "invalid_address_config.json"
			file, err := os.Create(filePath)
			Expect(err).To(BeNil())
			defer file.Close()

			_, err = io.WriteString(file, string(jsonData))
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			os.Remove("invalid_address_config.json")
		})

		It("should return an error", func() {
			err, _ := config.ReadAndValidateConfig("invalid_address_config", "json", ".")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("missing required fields in server configuration"))
		})
	})
})

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}
