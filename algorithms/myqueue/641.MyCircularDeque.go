package myqueue

// MyCircularDeque structure declare
type MyCircularDeque struct {
	front int   // 指向队列头部第1个有效数据的位置
	rear  int   // 指向队列尾部（即最后1个有效数据）的下一个位置，即下一个从队尾入队元素的位置
	cap   int   // 队列容量
	nums  []int // 队列元素
}

// NewMyCircularDeque Initialize structure. Set the size of the deque to be k.
func NewMyCircularDeque(k int) *MyCircularDeque {
	capacity := k + 1
	return &MyCircularDeque{
		cap:  capacity,
		nums: make([]int, capacity, capacity), // 申请一个len 和 cap都为k+1的数组
	}
}

// InsertFront Adds an item at the front of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) InsertFront(x int) bool {
	if q.IsFull() {
		return false
	}

	// front指针前移一位，取模是关键，指针前移需要先加q.cap再整体取模，因为可能出现负数
	// 而后移是加法不需要加q.cap
	q.front = (q.front - 1 + q.cap) % q.cap
	q.nums[q.front] = x
	return true
}

// InsertLast Adds an item at the rear of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) InsertLast(x int) bool {
	if q.IsFull() {
		return false
	}

	// 先赋值再移动指针，因为rear存的是最后1个有效数据的下一个位置
	q.nums[q.rear] = x
	// rear指针后移一位，取模是关键
	q.rear = (q.rear + 1) % q.cap
	return true
}

// DeleteFront Deletes an item from the front of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}

	// front指针后移一位
	q.front = (q.front + 1) % q.cap
	return true
}

// DeleteLast Deletes an item from the rear of Deque. Return true if the operation is successful.
func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}

	// rear指针前移一位
	q.rear = (q.rear - 1 + q.cap) % q.cap
	return true
}

// GetFront Get the front item from the deque.
func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}

	return q.nums[q.front]
}

// GetRear Get the rear item from the deque.
func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.nums[(q.rear-1+q.cap)%q.cap]
}

// IsEmpty Checks whether the circular deque is empty or not.
func (q *MyCircularDeque) IsEmpty() bool {
	return q.front == q.rear
}

// IsFull Checks whether the circular deque is full or not.
func (q *MyCircularDeque) IsFull() bool {
	return (q.rear+1)%q.cap == q.front
}
