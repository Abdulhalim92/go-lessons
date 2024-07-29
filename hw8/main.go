package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findMax(arr)) // 5

	arr = []int{1, 2, 0, 4, 5}
	fmt.Println(findMin(arr)) // 0

	arr = []int{-1, 2, -3, 4, 5}
	fmt.Println(countPositive(arr)) // 3

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(sumArray(arr)) // 15

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(averageArray(arr)) // 3

	arr = []int{1, 2, 3, 2, 4, 2, 5}
	fmt.Println(removeElement(arr, 2)) // [1 3 4 5]

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(multiplyArray(arr, 2)) // [2 4 6 8 10]

	arr = []int{1, 2, 3, 2, 4, 2, 5}
	fmt.Println(findAllIndexes(arr, 2)) // [1 3 5]

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(copyArray(arr)) // [1 2 3 4 5]

	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	fmt.Println(mergeArrays(arr1, arr2)) // [1 2 3 4 5 6]

	arr = []int{1, 3, 4, 5, 2}
	fmt.Println(swapMaxMin(arr)) // [5 3 4 1 2]

	arr = []int{1, 2, 3, 2, 1}
	fmt.Println(isPalindrome(arr)) // true

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(findSecondMax(arr)) // 4

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(reverseArray(arr)) // [5 4 3 2 1]

	arr = []int{1, 2, 3, 2, 4, 2, 5}
	fmt.Println(removeDuplicates(arr)) // [1 2 3 4 5]

	arr = []int{0, 1, 0, 3, 12}
	fmt.Println(moveZerosToEnd(arr)) // [1 3 12 0 0]

	arr1 = []int{1, 2, 3, 4, 5}
	arr2 = []int{3, 4, 5, 6, 7}
	fmt.Println(findIntersection(arr1, arr2)) // [3 4 5]

	arr1 = []int{1, 2, 3}
	arr2 = []int{1, 2, 3, 4, 5}
	fmt.Println(isSubset(arr1, arr2)) // true

	arr1 = []int{1, 3, 5}
	arr2 = []int{2, 4, 6}
	fmt.Println(mergeSortedArrays(arr1, arr2)) // [1 2 3 4 5 6]

	arr = []int{1, 2, 3, 1, 2, 3, 4, 5}
	fmt.Println(lengthOfLongestUniqueSubarray(arr)) // 5

	arr = []int{1, 2, 3, 4, 5}
	target := 5
	fmt.Println(findSubarraysWithSum(arr, target)) // [[2, 3], [5]]

	arr = []int{2, 7, 11, 15}
	target = 9
	x, y := findPairWithSum(arr, target)
	fmt.Printf("Pair: (%d, %d)\n", x, y) // Pair: (2, 7)

	arr = []int{3, 4, -1, 1}
	fmt.Println(findMissingPositive(arr)) // 2

	arr = []int{1, 2, 1, 3, 4, 2, 1, 1, 2}
	fmt.Println(maxSumSubarrayWithTwoDistinct(arr)) // 7

	arr = []int{1, 2, 3, 4, 5}
	target = 5
	fmt.Println(maxLengthSubarrayWithSum(arr, target)) // 2

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProductOfThree(arr)) // 60

	arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArraySum(arr)) // 6

	arr = []int{1, -2, 3, -4, 5}
	fmt.Println(moveNegativesToFront(arr)) // [-2 -4 1 3 5]

	arr = []int{1, 2, -3, 3, -2, 1}
	fmt.Println(maxLengthSubarrayWithZeroSum(arr)) // 5

	arr = []int{8, 12, 16}
	fmt.Println(gcdOfArray(arr)) // 4
}
