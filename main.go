package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(1000000000)
	}

	return result
}

func maximum(data []int) int {

	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}

	return maxVal
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	chunkMaxes := make([]int, CHUNKS)

	var wg sync.WaitGroup

	baseChunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		startIdx := i * baseChunkSize
		endIdx := startIdx + baseChunkSize
		if i == CHUNKS-1 {
			endIdx = len(data)
		}

		go func(idx int, start, end int) {
			defer wg.Done()

			localMax := data[start]
			for k := start + 1; k < end; k++ {
				if data[k] > localMax {
					localMax = data[k]
				}
			}
			chunkMaxes[idx] = localMax
		}(i, startIdx, endIdx)
	}
	wg.Wait()

	finalMax := chunkMaxes[0]
	for i := 1; i < len(chunkMaxes); i++ {
		if chunkMaxes[i] > finalMax {
			finalMax = chunkMaxes[i]
		}
	}

	return finalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)
	if len(data) == 0 {
		fmt.Println("Слайс пуст, завершаем работу.")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d микросекунд\n\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d микросекунд\n", maxMulti, elapsedMulti.Microseconds())
}
