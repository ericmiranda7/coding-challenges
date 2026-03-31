type MinStack struct {
    s []int
    mv int
}

func Constructor() MinStack {
    return MinStack{[]int{}, math.MaxInt}
}

func (this *MinStack) Push(val int) {
    this.s = append(this.s, val - this.mv)
    if val < this.mv {
        this.mv = val
    }
}

func (this *MinStack) Pop() {
    topVal := this.s[len(this.s)-1]
    if topVal < 0 {
        this.mv = this.mv - topVal
    }
    this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
    topVal := this.s[len(this.s)-1]
    var cv int
    if topVal < 0 {
        cv = this.mv
    } else {
        cv = topVal + this.mv
    }
    return cv
}

func (this *MinStack) GetMin() int {
    return this.mv
}
