type MyStack struct {
    q1 *Queue
    q2 *Queue
}

func Constructor() MyStack {
    return MyStack{MakeQueue(), MakeQueue()}
}

func (this *MyStack) Push(x int) {
    this.q1.enq(x)
    for !this.q2.isEmpty() {
        this.q1.enq(this.q2.deq())
    }
    this.q2 = this.q1
    this.q1 = MakeQueue()
}

func (this *MyStack) Pop() int {
    return this.q2.deq()
}

func (this *MyStack) Top() int {
    return this.q2.peek()
}

func (this *MyStack) Empty() bool {
    return this.q2.isEmpty()
}

type Queue struct {
    vals []int
}
func MakeQueue() *Queue {
    return &Queue{[]int{}}
}
func (q *Queue) enq(val int) {
    q.vals = append(q.vals, val)
}
func (q *Queue) deq() int {
    res := q.vals[0]
    q.vals = q.vals[1:]
    return res
}
func (q *Queue) size() int {
    return len(q.vals)
}
func (q *Queue) isEmpty() bool {
    return len(q.vals) == 0
}
func (q *Queue) peek() int {
    return q.vals[0]
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
