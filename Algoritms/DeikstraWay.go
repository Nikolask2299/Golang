/*
Дан ориентированный взвешенный граф. Найдите кратчайший путь от одной заданной вершины до другой.

Формат ввода
В первой строке содержатся три числа: N, S и F (1 ≤ N ≤ 100, 1 ≤ S, F ≤ N), где 
N — количество вершин графа, S — начальная вершина, а F — конечная. В следующих N строках вводится по N чисел,
 не превосходящих 100, – матрица смежности графа, где -1 означает, что ребра между вершинами нет, 
 а любое неотрицательное число — наличие ребра данного веса. На главной диагонали матрицы записаны нули.

Формат вывода
Последовательно выведите все вершины одного (любого) из кратчайших путей, или -1,
 если пути между указанными вершинами не существует
*/

package main

import (
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Item struct {
	value int
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil 
    item.index = -1
    *pq = old[0 : n-1]
    return item
}

type Vertex struct {
	key int
	connections map[int]int
}

type Graph struct {
	vertex_dict map[int]Vertex
}

func (g *Graph) add_edge(f, t, weight int) {
	if weight <= 0 {
		return
	}
	if g.get_vertex(f).key == 0{
		g.add_vertex(f)
	}
	if g.get_vertex(t).key == 0{
		g.add_vertex(t)
	}
	g.vertex_dict[f].add_connect(t, weight)
}

func (g *Graph) get_vertex(key int) Vertex {
	if g.vertex_dict[key].key == 0 {
		return Vertex{}
	}
	return g.vertex_dict[key]
}

func (g *Graph) add_vertex(key int) {
	var new_vertex Vertex
	new_vertex.key = key
	new_vertex.connections = make(map[int]int)
	g.vertex_dict[key] = new_vertex
}

func (v Vertex) add_connect(vertex int, weight int) {
	if weight <= 0 {
		return
	}
	v.connections[vertex] = weight
}


func AlgoritmDeicstra(dist map[int]int, s, f int, g Graph) []int {
	visited := make(map[int]bool)
	pq := make(PriorityQueue, 1)
	marshr := make(map[int]int) 
	marshr[s] = -1
	dist[s] = 0
	pq[0] = &Item{
		value: s,
		priority: 0,
		index: 0,
	}
	heap.Init(&pq)
	for len(pq) > 0 {
		elem := heap.Pop(&pq).(*Item)
		cur_dist, cur_vert := elem.priority, elem.value
		if visited[elem.value] {
			continue
		}
		if cur_dist > dist[cur_vert] {
			continue
		}
		visited[elem.value] = true
		for neig, weight := range g.vertex_dict[cur_vert].connections {
			distance := cur_dist + weight
			if distance < dist[neig] {
				dist[neig] = distance
				items := &Item{
					value: neig,
					priority: distance,
				}
				marshr[neig] = cur_vert
				heap.Push(&pq, items)
			}

		}
	}
	var resmas []int
	var down int
	if dist[f] == x {
		resmas = []int{-1}
		return resmas
	}
	down = f
	for down != -1 {
		resmas = append(resmas, down)
		down = marshr[down]
	}
	slices.Reverse(resmas)
	return resmas
}

const (
	x = 1_000_000_000
)

func main() {
	var n, s, f int
	var graph Graph
	graph.vertex_dict = make(map[int]Vertex)
	dist_dict := make(map[int]int)
	file, _ := os.Open("input.txt")
	defer file.Close()
	fmt.Fscan(file, &n, &s, &f)
	for i := 1; i < n + 1; i++ {
		dist_dict[i] = x
		for j := 1; j < n + 1; j++ {
			var vert int
			fmt.Fscan(file, &vert)
			graph.add_edge(i, j, vert)
		}
	}	
	
	res := AlgoritmDeicstra(dist_dict, s, f, graph)
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}