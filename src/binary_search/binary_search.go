package main

import (
	"io"
	"os"
	"fmt"
	"log"
	"math"
	"sort"
)

func binarySearch(key int, arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// N^2
func twoSum(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] + arr[j] == 0 {
				count++
			}
		}
	}
	return count;
}

// NlogN
func twoSumFast(arr []int) int {
	sort.Ints(arr)
	count := 0
	for i := 0; i < len(arr); i++ {
		if binarySearch(-arr[i], arr) > i {
			count++
		}
	}
	return count
}

// N
func twoSumFaster(arr []int) int {
	sort.Ints(arr)
	l := 0
	r := len(arr) - 1
	count := 0
	for l < r {
		if arr[l] + arr[r] == 0 {
			count++
			l++
			r--
		} else if arr[l] + arr[r] > 0 {
			r--
		} else {
			l++
		}
	}
	return count
}

// N^3
func threeSum(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i] + arr[j] + arr[k] == 0 {
					count++
				}
			}
		}
	}
	return count
}

// N^2logN
func threeSumFast(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if binarySearch( -(arr[i] + arr[j]), arr ) > j {
				count++
			}
		}
	}
	return count;
}

// N^2
func threeSumFaster(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		l := i + 1
		r := len(arr) - 1
		for l < r {
			if key + arr[l] + arr[r] == 0 {
				count++
				l++
				r--
			} else if key + arr[l] + arr[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return count
}

// NlogN
func closestPair(arr []float64) (float64, float64) {
	sort.Float64s(arr)
	min := math.MaxFloat64
	var p, q float64
	for i := 0; i < len(arr) - 1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < min {
			min = diff
			p = arr[i]
			q = arr[i+1]
		}
	}
	return p, q
}

func localMin(arr []int) int {
	low := 0
	high := len(arr) - 1
	return localMinRecur(arr, low, high)
}

// 2lgN
func localMinRecur(arr []int, low, high int) int {
	mid := (low + high) / 2
	if mid == len(arr) - 1 || mid == 0 {
		return -1
	}
	if arr[mid] < arr[mid - 1] && arr[mid] < arr[mid + 1] {
		return mid
	}
	// search in the half with the smaller neighbor
	if arr[mid - 1] < arr[mid + 1] {
		return localMinRecur(arr, low, mid - 1)
	} else {
		return localMinRecur(arr, mid + 1, high)
	}
}

// 3lgN
func bitonic(key int, arr []int) int {
	maxi := indexOfMax(arr)
	if arr[maxi] == key {
		return maxi
	}
	if left := binarySearch(key, arr[0:maxi]); left != -1 {
		return left
	}
	if right := binarySearch(key, arr[maxi+1:]); right != -1 {
		return right
	}
	return -1
}

func indexOfMax(arr []int) int {
	low := 0
	high := len(arr) - 1
	var mid int
	for low < high {
		mid = (low + high) / 2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid-1] > arr[mid+1] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return mid
}

func majority1(arr []int) int {
	sort.Ints(arr)
	mid := (0 + len(arr) - 1) / 2
	val := arr[mid]
	count := 0
	for _, x := range arr {
		if x == val {
			count++
		}
	}
	if count > len(arr) / 2 {
		return mid
	}
	return -1
}

func majority2(arr []int) int {
	ind := 0
	count := 1
	for i, x := range arr[1:] {
		if count == 0 {
			ind = i
			count = 1
		}
		if x == arr[ind] {
			count++
		} else {
			count--
		}
	}
	return ind
}

func contiSum(val int, arr []int) {
	partSum := make([]int, len(arr))
	sum := 0
	for i, x := range arr {
		sum += x
		partSum[i] = sum
	}
	
	var i, j int
	if j = binarySearch(val, partSum); j != -1 {
		fmt.Printf("from %d to %d\n", 0, j)
		return
	} else {
		for i = 0; i < len(partSum); i++ {
			j = binarySearch(val + partSum[i], partSum)
			if j > i {
				break
			}
		}
		if j != -1 {
			fmt.Printf("from %d to %d\n", i+1, j)
			return
		}
	}
	fmt.Println("can't find")
}

func tinyTest() {
	array := []int{30, -40, -20, -10, 40, 0, 10, 5}
	fmt.Printf("local minimum: %d\n", localMin(array))
	
	fmt.Printf("bitonic %d\n", bitonic(-10, []int{-40, -20, -10, 40, 30, 10, 5, 0, -1, -2, -3}))
	fmt.Printf("majority1: %d\n", majority1([]int{3, 4, 3, 5, 7, 3, 3, 3}))
	fmt.Printf("majority2: %d\n", majority2([]int{3, 4, 3, 5, 7, 3, 3, 3}))
	contiSum(4, []int{1, 2, 4, 7, 11, 15})
	/*
	fmt.Printf("twoSum: count of sum 0 pairs: %d\n", twoSum(array))
	fmt.Printf("twoSumFast: count of sum 0 pairs: %d\n", twoSumFast(array))
	fmt.Printf("twoSumFaster: count of sum 0 pairs: %d\n", twoSumFaster(array))
	fmt.Printf("threeSum: count of sum 0 triple: %d\n", threeSum(array))
	fmt.Printf("threeSumFast: count of sum 0 triple: %d\n", threeSumFast(array))
	fmt.Printf("threeSumFaster: count of sum 0 triple: %d\n", threeSumFaster(array))
	
	floats := []float64{0.3, 1.4, 5.5, 2.6, 3.8, 4.9, 7.2, 3.3}
	p, q := closestPair(floats)
	fmt.Printf("closestPair: %.2f, %.2f\n", p, q)
	*/
}

var testfile string = "1Kints.txt"
func bigTest() {
	fh, err := os.Open(testfile)
	if err != nil {
		log.Fatalln(err)	
	}
	array := []int{}
	var num int
	eof := false
	for !eof {
		_, err = fmt.Fscanf(fh, "%d\n", &num)
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			} else {
				eof = true
			}
		}
		array = append(array, num)
	}
	fmt.Printf("twoSum: count of sum 0 pairs: %d\n", twoSum(array))
	fmt.Printf("twoSumFast: count of sum 0 pairs: %d\n", twoSumFast(array))
	fmt.Printf("twoSumFaster: count of sum 0 pairs: %d\n", twoSumFaster(array))
	fmt.Printf("threeSum: count of sum 0 triple: %d\n", threeSum(array))
	fmt.Printf("threeSumFast: count of sum 0 triple: %d\n", threeSumFast(array))
	fmt.Printf("threeSumFaster: count of sum 0 triple: %d\n", threeSumFaster(array))
}

func main() {
	tinyTest()
	//bigTest()
}
