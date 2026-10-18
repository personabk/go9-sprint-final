package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	// Тест 1: Проверка на размер 0
	resultZero := generateRandomElements(0)
	if len(resultZero) != 0 {
		t.Errorf("Ожидался пустой слайс при размере 0, но длина равна %d", len(resultZero))
	}

	// Тест 2: Проверка на размер 10
	size := 10
	result := generateRandomElements(size)
	if len(result) != size {
		t.Errorf("Ожидалась длина слайса %d, но получена длина %d", size, len(result))
	}

	// Тест 3: Проверка на отрицательный размер (крайний случай)
	resultNeg := generateRandomElements(-5)
	if len(resultNeg) != 0 {
		t.Errorf("Ожидался пустой слайс при отрицательном размере, но длина равна %d", len(resultNeg))
	}
}
func TestMaximum(t *testing.T) {

	emptySlice := []int{}
	resEmpty := maximum(emptySlice)
	if resEmpty != 0 {
		t.Errorf("Для пустого слайса ожидалось 0, получено %d", resEmpty)
	}

	// Тест 2: Слайс из одного элемента
	singleSlice := []int{42}
	resSingle := maximum(singleSlice)
	if resSingle != 42 {
		t.Errorf("Для слайса [42] ожидалось 42, получено %d", resSingle)
	}

	// Тест 3: Обычный слайс (максимум в середине)
	normalSlice := []int{3, 7, 2, 9, 1}
	resNormal := maximum(normalSlice)
	if resNormal != 9 {
		t.Errorf("Для слайса [3, 7, 2, 9, 1] ожидалось 9, получено %d", resNormal)
	}

	// Тест 4: Максимум в начале
	startMaxSlice := []int{100, 1, 2, 3}
	resStart := maximum(startMaxSlice)
	if resStart != 100 {
		t.Errorf("Для слайса [100, 1, 2, 3] ожидалось 100, получено %d", resStart)
	}

	// Тест 5: Максимум в конце
	endMaxSlice := []int{1, 2, 3, 999}
	resEnd := maximum(endMaxSlice)
	if resEnd != 999 {
		t.Errorf("Для слайса [1, 2, 3, 999] ожидалось 999, получено %d", resEnd)
	}
}
func TestMaxChunks(t *testing.T) {
	// Генерируем небольшой слайс для теста (не нужно 100 млн, хватит 1000)
	data := generateRandomElements(1000)

	// Считаем максимум "эталонным" однопоточным способом
	expected := maximum(data)

	// Считаем максимум нашим многопоточным способом
	actual := maxChunks(data)

	if actual != expected {
		t.Errorf("maxChunks вернул %d, а maximum вернул %d. Результаты должны совпадать!", actual, expected)
	}
}
