/*
Реализуйте быструю сортировку, используя алгоритм из предыдущей задачи.
На каждом шаге выбирайте опорный элемент и выполняйте partition относительно него. Затем рекурсивно запуститесь от двух частей, на которые разбился исходный массив.
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func insertSort(masiv []int) {
	for i := 1; i < len(masiv); i++ {
		valies := masiv[i]
		for i > 0 && masiv[i-1] > valies {
			masiv[i] = masiv[i-1]
			i--
		}
		masiv[i] = valies
	}
}

func Quicksort(masiv []int, l, r int) {
	if l == r {
		return
	}
	if len(masiv[l:r + 1]) < 20 {
		insertSort(masiv[l : r + 1])
		return
	}
	indleg, indrig := partition(masiv, l, r)
	Quicksort(masiv, l, indleg)
	Quicksort(masiv, indrig, r)
}

func partition(masiv []int, indl, indr int) (int, int) {
	indexpiv := indl + rand.Intn(indr - indl) + 1
	pivot := masiv[indexpiv]
	masiv[indexpiv], masiv[indl] = masiv[indl], masiv[indexpiv]
	pivotLow, pivotHigh := indl, indr
	i := indl

	for i <= pivotHigh {
		if masiv[i] < pivot {
			masiv[pivotLow], masiv[i] = masiv[i], masiv[pivotLow]
			pivotLow++
			i++
		} else if masiv[i] > pivot {
			masiv[pivotHigh], masiv[i] = masiv[i], masiv[pivotHigh]
			pivotHigh--
		} else {
			i++
		}
	}
	return pivotLow, pivotHigh
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
	Quicksort(masiv, 0, (len(masiv) - 1))
	for _, num := range masiv {
		fmt.Fprint(out, num, " ")
	}
	fmt.Fprintln(out)
}