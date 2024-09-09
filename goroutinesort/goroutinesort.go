package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func mergeArrays(arr1, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}
	return merged
}

func main() {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter integers separated by space (press Enter when done):")
	if scanner.Scan() {
		input := scanner.Text()
		fields := strings.Fields(input)

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Invalid input, please enter integers.")
				return
			}
			numbers = append(numbers, num)
		}
	}

	n := len(numbers)
	if n == 0 {
		fmt.Println("No numbers entered.")
		return
	}

	partitionSize := (n + 3) / 4

	var wg sync.WaitGroup
	wg.Add(4)

	partitions := make([][]int, 4)

	for i := 0; i < 4; i++ {
		start := i * partitionSize
		end := (i + 1) * partitionSize
		if end > n {
			end = n
		}

		partitions[i] = numbers[start:end]

		go func(idx int, subarr []int) {
			fmt.Printf("Sorting subarray %d: %v\n", i+1, subarr)
			sort.Ints(subarr) // Sort the subarray
			fmt.Printf("Sorted subarray %d: %v\n", i+1, subarr)
			wg.Done()
		}(i, partitions[i])
	}

	wg.Wait()

	sortedArray := mergeArrays(partitions[0], partitions[1])
	sortedArray = mergeArrays(sortedArray, partitions[2])
	sortedArray = mergeArrays(sortedArray, partitions[3])

	fmt.Println("Final sorted array:", sortedArray)
}
