package main

import "math"

// Найти максимальный элемент в массиве.
func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// Найти минимальный элемент в массиве.
func findMin(arr []int) int {
	min := arr[0]
	for _, value := range arr[1:] {
		if value < min {
			min = value
		}
	}
	return min
}

// Подсчитать количество положительных чисел в массиве.
func countPositive(arr []int) int {
	count := 0
	for _, value := range arr {
		if value > 0 {
			count++
		}
	}
	return count
}

// Найти сумму всех элементов массива.
func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

// Найти среднее значение всех элементов массива.
func averageArray(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

// Удалить все вхождения заданного числа из массива.
func removeElement(arr []int, target int) []int {
	result := []int{}
	for _, value := range arr {
		if value != target {
			result = append(result, value)
		}
	}
	return result
}

// Умножить все элементы массива на заданное число.
func multiplyArray(arr []int, factor int) []int {
	for i := range arr {
		arr[i] *= factor
	}
	return arr
}

// Найти все индексы заданного числа в массиве.
func findAllIndexes(arr []int, target int) []int {
	indexes := []int{}
	for i, value := range arr {
		if value == target {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// Создать копию массива.
func copyArray(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	return copyArr
}

// Объединить два массива.
func mergeArrays(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
}

// Поменять местами максимальный и минимальный элементы массива.
func swapMaxMin(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	minIdx, maxIdx := 0, 0
	for i, value := range arr {
		if value < arr[minIdx] {
			minIdx = i
		}
		if value > arr[maxIdx] {
			maxIdx = i
		}
	}
	arr[minIdx], arr[maxIdx] = arr[maxIdx], arr[minIdx]
	return arr
}

// Проверить, является ли массив палиндромом.
func isPalindrome(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

// Найти второе наибольшее число в массиве.
func findSecondMax(arr []int) int {
	max, secondMax := arr[0], arr[0]
	for _, value := range arr {
		if value > max {
			secondMax = max
			max = value
		} else if value > secondMax && value < max {
			secondMax = value
		}
	}
	return secondMax
}

// Перевернуть массив.
func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// Удалить дубликаты из массива.
func removeDuplicates(arr []int) []int {
	result := []int{}
	for i, value := range arr {
		found := false
		for j := 0; j < i; j++ {
			if arr[j] == value {
				found = true
				break
			}
		}
		if !found {
			result = append(result, value)
		}
	}
	return result
}

// Переместить все нули в конце массива, сохраняя порядок ненулевых элементов.
func moveZerosToEnd(arr []int) []int {
	nonZeroCount := 0
	for _, value := range arr {
		if value != 0 {
			arr[nonZeroCount] = value
			nonZeroCount++
		}
	}
	for i := nonZeroCount; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

// Найти пересечение двух массивов.
func findIntersection(arr1, arr2 []int) []int {
	result := []int{}
	for _, value1 := range arr1 {
		for _, value2 := range arr2 {
			if value1 == value2 {
				found := false
				for _, v := range result {
					if v == value1 {
						found = true
						break
					}
				}
				if !found {
					result = append(result, value1)
				}
			}
		}
	}
	return result
}

// Проверить, является ли массив подмножеством другого массива.
func isSubset(arr1, arr2 []int) bool {
	for _, value1 := range arr1 {
		found := false
		for _, value2 := range arr2 {
			if value1 == value2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Объединить два отсортированных массива в один отсортированный.
func mergeSortedArrays(arr1, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result
}

// Найти длину самого длинного подмассива, в котором все элементы различны.
func lengthOfLongestUniqueSubarray(arr []int) int {
	maxLength := 0
	for i := 0; i < len(arr); i++ {
		visited := make(map[int]bool)
		for j := i; j < len(arr); j++ {
			if visited[arr[j]] {
				break
			}
			visited[arr[j]] = true
			if j-i+1 > maxLength {
				maxLength = j - i + 1
			}
		}
	}
	return maxLength
}

// Найти все подмассивы, сумма которых равна заданному числу.
func findSubarraysWithSum(arr []int, target int) [][]int {
	result := [][]int{}
	for start := 0; start < len(arr); start++ {
		sum := 0
		for end := start; end < len(arr); end++ {
			sum += arr[end]
			if sum == target {
				result = append(result, arr[start:end+1])
			}
		}
	}
	return result
}

// Найти пару элементов в массиве, сумма которых равна заданному числу.
func findPairWithSum(arr []int, target int) (int, int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return arr[i], arr[j]
			}
		}
	}
	return 0, 0
}

// Найти наименьший положительный элемент, отсутствующий в массиве.
func findMissingPositive(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for arr[i] > 0 && arr[i] <= n && arr[arr[i]-1] != arr[i] {
			arr[arr[i]-1], arr[i] = arr[i], arr[arr[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// Найти максимальную сумму подмассива с условием, что подмассив не должен содержать более двух различных элементов.
func maxSumSubarrayWithTwoDistinct(arr []int) int {
	maxSum, start, currentSum := 0, 0, 0
	distinct := make(map[int]int)
	for _, value := range arr {
		distinct[value]++
		currentSum += value
		for len(distinct) > 2 {
			distinct[arr[start]]--
			if distinct[arr[start]] == 0 {
				delete(distinct, arr[start])
			}
			currentSum -= arr[start]
			start++
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// Найти максимальную длину подмассива, сумма элементов которого равна заданному числу.
func maxLengthSubarrayWithSum(arr []int, target int) int {
	maxLength := 0
	for start := 0; start < len(arr); start++ {
		sum := 0
		for end := start; end < len(arr); end++ {
			sum += arr[end]
			if sum == target {
				if end-start+1 > maxLength {
					maxLength = end - start + 1
				}
			}
		}
	}
	return maxLength
}

// Найти максимальное произведение трех чисел в массиве.
func maxProductOfThree(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32
	for _, value := range arr {
		if value > max1 {
			max3 = max2
			max2 = max1
			max1 = value
		} else if value > max2 {
			max3 = max2
			max2 = value
		} else if value > max3 {
			max3 = value
		}
		if value < min1 {
			min2 = min1
			min1 = value
		} else if value < min2 {
			min2 = value
		}
	}
	return max(max1*max2*max3, min1*min2*max1)
}

// Найти подмассив с максимальной суммой.
func maxSubArraySum(arr []int) int {
	maxSum, currentSum := math.MinInt32, 0
	for _, value := range arr {
		currentSum += value
		if currentSum > maxSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}

// Переместить все отрицательные числа в начало массива, сохраняя порядок остальных чисел.
func moveNegativesToFront(arr []int) []int {
	negCount := 0
	for _, value := range arr {
		if value < 0 {
			arr = append(arr[:negCount], append([]int{value}, arr[negCount:]...)...)
			negCount++
		}
	}
	return arr
}

// Найти подмассив с наибольшей длиной, сумма элементов которого равна нулю.
func maxLengthSubarrayWithZeroSum(arr []int) int {
	maxLength := 0
	for start := 0; start < len(arr); start++ {
		sum := 0
		for end := start; end < len(arr); end++ {
			sum += arr[end]
			if sum == 0 {
				if end-start+1 > maxLength {
					maxLength = end - start + 1
				}
			}
		}
	}
	return maxLength
}

// Найти наибольший общий делитель всех элементов массива.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfArray(arr []int) int {
	result := arr[0]
	for _, value := range arr[1:] {
		result = gcd(result, value)
	}
	return result
}
