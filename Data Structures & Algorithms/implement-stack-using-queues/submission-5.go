// inp a x b c d f
// f d c b x a -- top

// q1 d c b x a
// q2 f

type MyStack struct {
    q1 *Queue
    q2 *Queue
}

func Constructor() MyStack {
    return MyStack{&Queue{[]int{}}, &Queue{[]int{}}}
}

func (this *MyStack) Push(x int) {
    for _ = range this.q1.size() {
        this.q2.push(this.q1.pop())
    }

    this.q1.push(x)
    this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
    if this.q2.size() > 0 {
        return this.q2.pop()
    }

    return this.q1.pop()
}

func (this *MyStack) Top() int {
    if this.q2.size() > 0 {
        return this.q2.peek()
    }

    return this.q1.peek()
}

func (this *MyStack) Empty() bool {
    return this.q1.size() + this.q2.size() == 0
}

type Queue struct {
    arr []int
}

func (q *Queue) push(val int) {
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

func (q *Queue) isEmpty() bool {
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
