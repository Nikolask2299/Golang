package main

import "fmt"

func sortenTim(spis []int) []int {
	if len(spis) == 1 {
		return spis
	}
	var left_spis []int
	var right_spis []int
	mod := len(spis) / 2
	left_spis = append(left_spis, spis[:mod]...)
	right_spis = append(right_spis, spis[mod:]...)   
	left_spis = sortenTim(left_spis)
	right_spis = sortenTim(right_spis)
	var all_spis []int
	if left_spis[0] < right_spis[0] {
		all_spis = append(left_spis, right_spis...)
	} else {
		all_spis = append(right_spis, left_spis...)
	}
	if len(all_spis) == 2 {
		return all_spis
	}
	for i := 1; i < len(all_spis); i++ {
		value := all_spis[i]
			for i > 0 && all_spis[i - 1] > value {
				all_spis[i] = all_spis[i - 1]
				i--
			}
			all_spis[i] = value
		}
		return all_spis
	}

func main(){
	fmt.Println("Введите размер списка:")
	var n int
	fmt.Scanln(&n)
	var spis []int
	fmt.Println("Введите элементы списка")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num) 
		spis = append(spis, num)		
	}
	fmt.Println(spis)
	fmt.Println("Отсортированный список")
	sortspis := sortenTim(spis)
	fmt.Println(sortspis)	
}