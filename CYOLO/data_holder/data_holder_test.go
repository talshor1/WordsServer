package data_holder_test

import (
	"CYOLO/data_holder"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestDataHolder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DataHolder Suite")
}

var _ = Describe("DataHolder", func() {
	var holder *data_holder.DataHolder

	BeforeEach(func() {
		holder = data_holder.NewDataHolder()
	})

	When("adding many words", func() {
		It("should update FrequencyMap and TopFive correctly", func() {
			for i := 0; i < 50; i++ {
				holder.AddWord("Apple")
			}
			for i := 0; i < 40; i++ {
				holder.AddWord("Banana")
			}
			for i := 0; i < 30; i++ {
				holder.AddWord("Cold")
			}
			for i := 0; i < 20; i++ {
				holder.AddWord("Damn")
			}
			for i := 0; i < 10; i++ {
				holder.AddWord("Egg")
			}
			for i := 0; i < 100; i++ {
				holder.AddWord("Fat")
			}

			expectedFrequency := map[string]int{
				"Fat":    100,
				"Apple":  50,
				"Banana": 40,
				"Cold":   30,
				"Damn":   20,
			}

			for word, expectedCount := range expectedFrequency {
				actualCount, exists := holder.FrequencyMap[word]
				Expect(exists).To(BeTrue(), "FrequencyMap should contain %s", word)
				Expect(actualCount).To(Equal(expectedCount), "FrequencyMap[%s] is incorrect", word)
			}

			expectedTopFive := "Fat 100 Apple 50 Banana 40 Cold 30 Damn 20 "
			actualTopFive := holder.GetTopFive()
			Expect(actualTopFive).To(Equal(expectedTopFive), "TopFive is incorrect")
		})
	})

	When("having less than 5 words", func() {
		It("should show only those words", func() {
			holder.AddWord("Apple")
			actualTopFive := holder.GetTopFive()
			expectedTopFive := "Apple 1 "
			Expect(actualTopFive).To(Equal(expectedTopFive), "TopFive is incorrect")

			holder.AddWord("Banana")
			actualTopFive = holder.GetTopFive()
			expectedTopFive = "Apple 1 Banana 1 "
			Expect(actualTopFive).To(Equal(expectedTopFive), "TopFive is incorrect")
		})
	})

	Context("Least", func() {
		When("adding words", func() {
			It("should return the least one properly", func() {
				for i := 0; i < 1000; i++ {
					for j := 0; j < 2; j++ {
						holder.AddWord(fmt.Sprintf("%d", i))
					}
				}
				holder.AddWord("LONE")
				Expect(holder.GetLeast()).To(Equal("LONE 1"))
			})
		})

	})
})
