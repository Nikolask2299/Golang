/*
Реализуйте сортировку слиянием, используя алгоритм из предыдущей задачи.
На каждом шаге делите массив на две части,
 сортируйте их независимо и сливайте с помощью уже 
 реализованной функции.
 В первой строке входного файла содержится число N — количество элементов массива (0 ≤ N ≤ 106).
Во второй строке содержатся N целых чисел ai, разделенных пробелами (-109 ≤ ai ≤ 109).
Выведите результат сортировки, то есть N целых чисел, разделенных пробелами, в порядке неубывания.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func merge(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		resMasiv := append(a, b...)
		return resMasiv
	}
	resMasiv := make([]int, len(a)+len(b))
	aind, bind := 0, 0
	for i := range resMasiv {
		if aind == len(a) {
			resMasiv[i] = b[bind]
			bind++
			continue
		}
		if bind == len(b) {
			resMasiv[i] = a[aind]
			aind++
			continue
		}
		if a[aind] < b[bind] {
			resMasiv[i] = a[aind]
			aind++
		} else {
			resMasiv[i] = b[bind]
			bind++
		}
	}
	
	return resMasiv
}


func MergeSort(masiv []int) []int {
	if len(masiv) < 2 {
		return masiv
	}
	if len(masiv) == 2 {
		if masiv[0] > masiv[1] {
			masiv[0], masiv[1] = masiv[1], masiv[0]
			return masiv
		} else {
			return masiv
		}
	}
	masiv1 := MergeSort(masiv[:len(masiv) / 2])
	masiv2 := MergeSort(masiv[len(masiv) / 2:])
	return merge(masiv1, masiv2)
}

func main() {
	var n int
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()
	input := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(input, &n)
	masiv := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &masiv[i])
	}
	sortMasiv := MergeSort(masiv)
	for _, num := range sortMasiv {
		fmt.Fprint(out, num, " ")
	}
	fmt.Fprintln(out)
}
