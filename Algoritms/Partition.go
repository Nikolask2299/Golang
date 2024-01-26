/*
Базовым алгоритмом для быстрой сортировки является алгоритм partition, который разбивает набор элементов на две части относительно заданного предиката.
По сути элементы массива просто меняются местами так, что левее некоторой точки в нем после этой операции лежат элементы, удовлетворяющие заданному предикату, а справа — не удовлетворяющие ему.
Например, при сортировке можно использовать предикат «меньше опорного», что при оптимальном выборе опорного элемента может разбить массив на две примерно равные части.
Напишите алгоритм partition в качестве первого шага для написания быстрой сортировки.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func partition(masiv []int, indl, indr, pivot int) (int, int) {
	if  len(masiv) == 0 {
		return -1, 1
	}
	for {
		if indl == indr {
			if masiv[indl] >= pivot {
				indl--
			}
			if masiv[indr] < pivot {
				indr++
			}
			break
		}
		if masiv[indl] < pivot {
			indl++
			continue
		} 
		if masiv[indr] >= pivot {
			indr--
			continue
		} 
		masiv[indl], masiv[indr] = masiv[indr], masiv[indl]
	}
	
	return indl, indr
}

func main() {
	var n, elemop int
	file, err := os.Open("input.txt")
  	if err != nil {
    	return
  	}
  	defer file.Close()
  	input := bufio.NewReader(file)
  	out := bufio.NewWriter(os.Stdout)
  	defer out.Flush()
  	fmt.Fscan(input, &n)
  	if n == 0 {
    	fmt.Fscan(input, &elemop)
	}
  	masiv := make([]int, n)
  	for i := 0; i < n; i++ {
    	fmt.Fscan(input, &masiv[i])
  	}
  	fmt.Fscan(input, &elemop)
  	indexmin, indexmax := partition(masiv, 0, len(masiv) - 1, elemop)
	if indexmin < 0 {
		indexmin = 0
	} else {
		indexmin = len(masiv[:indexmin+1])
	}	
	if indexmax > (len(masiv) - 1){
		indexmax = 0
	} else {
		indexmax =	len(masiv[indexmax:])
	}

  	fmt.Fprintln(out, indexmin)
  	fmt.Fprintln(out, indexmax)
}