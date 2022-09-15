package data_structure

import (
	"math/rand"
	"testing"

	"github.com/go-playground/assert"
)

var randomSlice = []int{100, 56, -34, 75, 23, 55, 80, 101, -45, -23}

// return slice of integer contains positive and negative number
func GenerateSlice(length int) []int {
	slice := make([]int, length, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func TestGenerateSlice(t *testing.T) {
	println(GenerateSlice(20))
}

// sort the given datalist by moving each element
func BubbleSort(datalist []int) []int {
	var highestIndex = len(datalist) - 1
	var sorted = false

	for !sorted {
		swapped := false
		for i := 0; i < highestIndex; i++ {
			if datalist[i] > datalist[i+1] {
				datalist[i+1], datalist[i] = datalist[i], datalist[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		highestIndex -= 1
	}

	return datalist
}

func TestBubbleSort(t *testing.T) {
	assert.Equal(t, []int{-45, -34, -23, 23, 55, 56, 75, 80, 100, 101}, BubbleSort(randomSlice))
}

// sort the given datalist by seperate element into 2 side
func QuickSort(datalist []int) []int {
	if len(datalist) < 2 {
		return datalist
	}

	var left, right = 0, len(datalist) - 1

	for k := range datalist {
		if datalist[k] < datalist[right] {
			datalist[left], datalist[k] = datalist[k], datalist[left]
			left++
		}
	}

	datalist[left], datalist[right] = datalist[right], datalist[left]

	QuickSort(datalist[:left])
	QuickSort(datalist[left+1:])

	return datalist
}

func TestQuickSort(t *testing.T) {
	assert.Equal(t, []int{-45, -34, -23, 23, 55, 56, 75, 80, 100, 101}, QuickSort(randomSlice))
}

// sort the given datalist by select the lowest value index
func SelectionSort(datalist []int) []int {
	var length = len(datalist)

	for i := 0; i < length; i++ {
		var minIdx = i
		for j := i + 1; j < length; j++ {
			if datalist[j] < datalist[minIdx] {
				minIdx = j
			}
		}
		datalist[i], datalist[minIdx] = datalist[minIdx], datalist[i]
	}

	return datalist
}

func TestSelectionSort(t *testing.T) {
	assert.Equal(t, []int{-45, -34, -23, 23, 55, 56, 75, 80, 100, 101}, SelectionSort(randomSlice))
}
