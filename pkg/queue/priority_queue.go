package queue

import (
	"container/heap"
	"github.com/Nicolas-ggd/go-notification/pkg/storage/models/request"
	"sync"
)

type NotificationHeap request.NotificationRequest

// An PriorityQueue is a min-heap of notifications.
type PriorityQueue struct {
	notification []*NotificationHeap
	mutex        sync.Mutex
}

func (pq PriorityQueue) Len() int { return len(pq.notification) }

func (pq PriorityQueue) Less(i, j int) bool {
	priorities := map[string]int{
		"error":   1,
		"warning": 2,
		"info":    3,
	}

	return priorities[pq.notification[i].Type] > priorities[pq.notification[j].Type]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.notification[i], pq.notification[j] = pq.notification[j], pq.notification[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	item := x.(*NotificationHeap)
	pq.notification = append(pq.notification, item)
	heap.Fix(pq, pq.Len()-1)
}

func (pq *PriorityQueue) Pop() interface{} {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	old := pq.notification
	n := len(old)
	item := old[n-1]
	pq.notification = old[0 : n-1]
	heap.Init(pq)
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		notification: []*NotificationHeap{},
	}

	heap.Init(pq)
	return pq
}
