/*
Поразрядная сортировка является одним из видов сортировки, которые работают практически за линейное от 
размера сортируемого массива время. Такая скорость достигается за счет того, что эта сортировка 
использует внутреннюю структуру сортируемых объектов. Изначально этот алгоритм использовался для сортировки перфокарт. 
Первая его компьютерная реализация была создана в университете MIT Гарольдом Сьюардом (Harold Н. Seward). 
Опишем алгоритм подробнее. Пусть задан массив строк s1 , ..., si причём все строки имеют одинаковую длину m. 
Работа алгоритма состоит из m фаз. На i -ой фазе строки сортируются па i -ой с конца букве. 
Происходит это следующим образом. Будем, для простоты, в этой задаче рассматривать строки из цифр от 0 до 9. 
Для каждой цифры создается «корзина» («bucket»), после чего строки si распределяются по «корзинам» в 
соответствии с i-ой цифрой с конца. Строки, у которых i-ая с конца цифра равна j попадают в j-ую корзину 
(например, строка 123 на первой фазе попадет в третью корзину, на второй — во вторую, на третьей — в первую). 
После этого элементы извлекаются из корзин в порядке увеличения номера корзины. 
Таким образом, после первой фазы строки отсортированы по последней цифре, после двух фаз — по двум последним, ..., 
после m фаз — по всем. При важно, чтобы элементы в корзинах сохраняли тот же порядок, что и в исходном массиве 
(до начала этой фазы). Например, если массив до первой фазы имеет вид: 111, 112, 211, 311, то элементы по корзинам 
распределятся следующим образом: в первой корзине будет. 111, 211, 311, а второй: 112. 
Напишите программу, детально показывающую работу этого алгоритма на заданном массиве.

Формат ввода
Первая строка входного файла содержит целое число n (1 ≤ n ≤ 1000) . 
Последующие n строк содержат каждая по одной строке si . Длины всех si , одинаковы и не превосходят 20. 
Все si состоят только из цифр от 0 до 9.

Формат вывода
В выходной файл выведите исходный массив строк в, 
состояние «корзин» после распределения элементов по ним для каждой фазы и отсортированный массив. 
Следуйте формату, приведенному в примере.

*/


package main

import (
	"fmt"
	"os"
	"strconv"
)

func BitwiseSort(masiv []string, disch int) map[string][]string {
	Bucket := map[string][]string{
		"0": make([]string, 0),
		"1": make([]string, 0),
		"2": make([]string, 0),
		"3": make([]string, 0),
		"4": make([]string, 0),
		"5": make([]string, 0),
		"6": make([]string, 0),
		"7": make([]string, 0),
		"8": make([]string, 0),
		"9": make([]string, 0),
	}
	for _, num := range masiv {
		ind := string(num[disch]) 
		Bucket[ind] = append(Bucket[ind], num)
	}
	return Bucket
}

func PrintBucket(Bucket map[string][]string) []string {
	var res []string
	for i := 0; i < 10; i++ {
		ind := strconv.Itoa(i)
		if len(Bucket[ind]) == 0 {
			fmt.Fprintln(os.Stdout, "Bucket " + ind + ": " + "empty")
			continue
		}
		res = append(res, Bucket[ind]...)
		fmt.Fprint(os.Stdout, "Bucket " + ind + ": ")
		printMass(Bucket[ind])
	}
	return res
}

func printMass(masiv []string) {
	for i, num := range masiv {
		if i != len(masiv) - 1 {
			num += ", "
		}
		fmt.Fprint(os.Stdout, num)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Fscan(os.Stdin, &n)
	masiv := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &masiv[i])
	}
	fmt.Fprintln(os.Stdout, "Initial array:")
	printMass(masiv)
	disch := len(masiv[0]) - 1
	for i := 1; i < len(masiv[0]) + 1; i++ {
		fmt.Fprintln(os.Stdout, "**********")
		fmt.Fprintln(os.Stdout, "Phase " + strconv.Itoa(i))
		bucket := BitwiseSort(masiv, disch)
		masiv = PrintBucket(bucket)
		disch--
	}
	fmt.Fprintln(os.Stdout, "**********")
	fmt.Fprintln(os.Stdout, "Sorted array:")
	printMass(masiv)
}
