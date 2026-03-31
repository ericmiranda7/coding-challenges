type MyStack struct {
    q1 *list.List
    q2 *list.List
}

func Constructor() MyStack {
    return MyStack{list.New(), list.New()}
}

func (this *MyStack) Push(x int) {
    this.q1.PushBack(x)

    for this.q2.Len() > 0 {
        this.q1.PushBack(this.q2.Remove(this.q2.Front()))
    }

    tmp := this.q1
    this.q1 = this.q2
    this.q2 = tmp
}

func (this *MyStack) Pop() int {
    return this.q2.Remove(this.q2.Front()).(int)
}

func (this *MyStack) Top() int {
    return this.q2.Front().Value.(int)
}

func (this *MyStack) Empty() bool {
    return this.q2.Len() == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
