package data_holder

import (
	"CYOLO/helper"
	"fmt"
	"sort"
	"sync"
)

type Tuple struct {
	Word   string
	Amount int
}

type TopFive struct {
	Tuples []*Tuple
}

type DataHolder struct {
	FrequencyMap map[string]int
	TopFive      TopFive
	mu           sync.Mutex
}

func NewDataHolder() *DataHolder {
	return &DataHolder{
		FrequencyMap: make(map[string]int),
		TopFive:      TopFive{Tuples: make([]*Tuple, 0)},
	}
}

func (d *DataHolder) AddWord(word string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.FrequencyMap[word]++

	d.updateTopFive(word, d.FrequencyMap[word])
}

func (d *DataHolder) updateTopFive(wordToUpdate string, wordAmount int) {
	for _, tuple := range d.TopFive.Tuples {
		if tuple != nil && tuple.Word == wordToUpdate {
			tuple.Amount = wordAmount
			d.TopFive.sort()
			return
		}
	}

	if len(d.TopFive.Tuples) < 5 {
		d.TopFive.Tuples = append(d.TopFive.Tuples, &Tuple{Word: wordToUpdate, Amount: wordAmount})
		d.TopFive.sort()
		return
	}

	if d.TopFive.Tuples[4] != nil && d.TopFive.Tuples[4].Amount < wordAmount {
		d.TopFive.Tuples[4] = &Tuple{Word: wordToUpdate, Amount: wordAmount}
		d.TopFive.sort()
	}
}

func (tf *TopFive) sort() {
	sort.Slice(tf.Tuples, func(i, j int) bool {
		return tf.Tuples[i] != nil && tf.Tuples[j] != nil && tf.Tuples[i].Amount > tf.Tuples[j].Amount
	})
}

func (d *DataHolder) GetTopFive() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	str := ""
	for _, t := range d.TopFive.Tuples {
		if t == nil {
			return str
		}
		str += fmt.Sprintf("%s %d ", t.Word, t.Amount)
	}
	return str
}

func (d *DataHolder) GetLeast() string {
	if len(d.FrequencyMap) == 0 {
		return ""
	}

	var minWord string
	var minValue int

	firstIteration := true

	for word, value := range d.FrequencyMap {
		if firstIteration || value < minValue {
			minValue = value
			minWord = word
			firstIteration = false
		}
	}

	return fmt.Sprintf("%s %d", minWord, minValue)
}

func (d *DataHolder) GetMedian() string {
	var values []int
	for _, value := range d.FrequencyMap {
		values = append(values, value)
	}

	helper.QuickSort(values)
	medianValue := values[len(values)/2]

	for k, v := range d.FrequencyMap {
		if v == medianValue {
			return fmt.Sprintf("%s %d", k, v)
		}
	}

	return ""
}
