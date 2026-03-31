// x a b c d e f
// b a x

type MyStack struct {
    q *Queue
}

func Constructor() MyStack {
	return MyStack{&Queue{}}
}

func (this *MyStack) Push(x int) {
	this.q.enq(x)
	for _ = range this.q.size() - 1 {
		this.q.enq(this.q.pop())
	}
}

func (this *MyStack) Pop() int {
	return this.q.pop()
}

func (this *MyStack) Top() int {
	return this.q.peek()
}

func (this *MyStack) Empty() bool {
	return this.q.isEmpty()
}

type Queue struct {
	arr []int
}

func (q *Queue) enq(val int) {
	q.arr = append(q.arr, val)
}

func (q *Queue) peek() int {
	return q.arr[0]
}

func (q *Queue) pop() int {
	r := q.arr[0]
	q.arr = q.arr[1:]
	return r
}

func (q Queue) size() int {
	return len(q.arr)
}

func (q Queue) isEmpty() bool {
	return len(q.arr) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
