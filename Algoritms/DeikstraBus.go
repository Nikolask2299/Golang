/*
Между некоторыми деревнями края Васюки ходят автобусы.
Поскольку пассажиропотоки здесь не очень большие, то автобусы ходят всего несколько раз в день.
Марии Ивановне требуется добраться из деревни d в деревню v как можно быстрее 
(считается, что в момент времени 0 она находится в деревне d).

Формат ввода
Сначала вводится число N – общее число деревень (1 <= N <= 100), 
затем номера деревень d и v, за ними следует количество автобусных рейсов R (0 <= R <= 10000). 
Далее идут описания автобусных рейсов. Каждый рейс задается номером деревни отправления, временем отправления, 
деревней назначения и временем прибытия (все времена – целые от 0 до 10000). Если в момент t пассажир приезжает 
в какую-то деревню, то уехать из нее он может в любой момент времени, начиная с t.

Формат вывода
Выведите минимальное время, когда Мария Ивановна может оказаться в деревне v. 
Если она не сможет с помощью указанных автобусных рейсов добраться из d в v, выведите -1.
*/

package main

import (
	"container/heap"
	"fmt"
	"os"
)

type PriorityQueue []*SheduleVilage

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	graph := x.(*SheduleVilage)
	graph.index = n
	*pq = append(*pq, graph)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	graph := old[n-1]
	old[n-1] = nil   
	graph.index = -1 
	*pq = old[0 : n-1]
	return graph
}

func (pq *PriorityQueue) update(item *SheduleVilage, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

const (
	x = 1_000_000_000
)

type SheduleVilage struct {
	Shedule []Bus
	priority int
	timetrip int
	isChecked bool
	inHeap bool
	index int
}

type Bus struct{
	point int
	depart int
	arrival int
	timetrip int
}


func Deikstra(Shedule []SheduleVilage, d, v int) int {
	Shedule[d].priority, Shedule[d].index = 0, 0
	pq := make(PriorityQueue, 1, len(Shedule))
	pq[0] = &Shedule[d]
	heap.Init(&pq)
	for pq.Len() > 0 {
		Shed := heap.Pop(&pq).(*SheduleVilage)
		Shed.inHeap = false
		Shed.isChecked = true
		for _, bus := range Shed.Shedule {
			if bus.depart < Shed.timetrip {
				continue
			} 
			if Shedule[bus.point].isChecked {
				continue
			}
			timed := Shed.priority + (bus.depart - Shed.timetrip) + bus.timetrip
			if timed < Shedule[bus.point].priority {
				Shedule[bus.point].priority = timed
				Shedule[bus.point].timetrip = bus.arrival
			}
			if Shedule[bus.point].inHeap {
				pq.update(&Shedule[bus.point], Shedule[bus.point].priority)
				pq.update(&Shedule[bus.point], Shedule[bus.point].timetrip)
			} else {
				Shedule[bus.point].inHeap = true
				heap.Push(&pq, &Shedule[bus.point])
			}
		}
	}
	if Shedule[v].isChecked {
		return Shedule[v].priority
	}
	return -1
}

func main() {
	var n, d, v, r int
	file, _ := os.Open("input.txt")
	fmt.Fscan(file, &n)
	fmt.Fscan(file, &d, &v)
	fmt.Fscan(file, &r)
	Shedule := make([]SheduleVilage, n + 1)
	for i := 1; i < n + 1; i++ {
		shed := &SheduleVilage{priority: x}
		Shedule[i] = *shed
	}
	var a, ta, b, tb int
	for i := 0; i < r; i++ {
		fmt.Fscan(file, &a, &ta, &b, &tb)
		Shedule[a].Shedule = append(Shedule[a].Shedule, Bus{point: b, depart: ta, arrival: tb, timetrip: tb - ta})		
	}
	res := Deikstra(Shedule, d, v)
	fmt.Println(res)		
}