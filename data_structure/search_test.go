package data_structure

import (
	"testing"

	"github.com/go-playground/assert"
)

var intArray = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

// check if datalist contain key sequentially
func LinearSearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}

	return false
}

func TestLinearSearch(t *testing.T) {
	assert.Equal(t, true, LinearSearch(intArray, 70))
	assert.Equal(t, false, LinearSearch(intArray, 75))
}

// check if datalist contain item by skip half of datalist
func BinarySearch(datalist []int, key int) bool {
	lowestIndex := 0
	highestIndex := len(datalist) - 1

	for lowestIndex <= highestIndex {
		meadianIndex := (lowestIndex + highestIndex) / 2

		if datalist[meadianIndex] < key {
			lowestIndex = meadianIndex + 1
		} else if datalist[meadianIndex] > key {
			highestIndex = meadianIndex - 1
		} else {
			// condition if median == key
			return true
		}
	}

	return false
}

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, true, BinarySearch(intArray, 70))
	assert.Equal(t, false, BinarySearch(intArray, 75))
}

// check if datalist caontains key by using guessed number
func InterpolationSearch(datalist []int, key int) bool {
	min, max := datalist[0], datalist[len(datalist)-1]
	lowestIndex, highestIndex := 0, len(datalist)-1

	for lowestIndex <= highestIndex {
		//make a guess
		var guess int
		if highestIndex == lowestIndex {
			guess = highestIndex
		} else {
			diff := highestIndex - lowestIndex
			offset := int(float64(diff-1) * (float64(key - min)) / float64(max-min))
			guess = lowestIndex + offset
		}

		if datalist[guess] > key {
			highestIndex = guess - 1
			max = datalist[highestIndex]
		} else if datalist[guess] < key {
			lowestIndex = guess + 1
			min = datalist[lowestIndex]
		} else {
			// condition if index guess == key
			return true
		}
	}

	return false
}

func TestInterpolationSearch(t *testing.T) {
	assert.Equal(t, true, InterpolationSearch(intArray, 100))
	assert.Equal(t, false, InterpolationSearch(intArray, 75))
}
