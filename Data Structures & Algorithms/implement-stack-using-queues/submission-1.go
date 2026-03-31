type MyStack struct {
    q1 *Queue
    q2 *Queue
}

func Constructor() MyStack {
    q1 := MakeQueue()
    q2 := MakeQueue()
    return MyStack{q1, q2}
}

func (this *MyStack) Push(x int) {
    if this.q1.size() == 0 {
        this.q1.enq(x)
        for this.q2.size() > 0 {
            this.q1.enq(this.q2.deq())
        }
    } else {
        this.q2.enq(x)
        for this.q1.size() > 0 {
            this.q2.enq(this.q1.deq())
        }
    }
}

func (this *MyStack) Pop() int {
    if this.q1.size() == 0 {
        return this.q2.deq()
    } else {
        return this.q1.deq()
    }
}

func (this *MyStack) Top() int {
    if this.q1.size() == 0 {
        return this.q2.peek()
    } else {
        return this.q1.peek()
    }
}

func (this *MyStack) Empty() bool {
    return this.q1.size() + this.q2.size() == 0
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
    val := q.vals[0]
    q.vals = q.vals[1:]
    return val
}
func (q *Queue) peek() int {
    return q.vals[0]
}
func (q *Queue) size() int {
    return len(q.vals)
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
