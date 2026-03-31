type KthLargest struct {
    arr []int
    k int
}


func Constructor(k int, nums []int) KthLargest {
    arr := []int{}
    arr = append(arr, -1)

    kl := KthLargest{arr, k}

    for _, v := range nums {
        kl.addVal(v)
    }

    return kl
}

func (this *KthLargest) addVal(val int) {
    this.arr = append(this.arr, val)
    this.percolateUp(len(this.arr)-1)
    if len(this.arr) > this.k + 1 {
        this.remove()
    }
}

func (this *KthLargest) remove() {
    this.arr[1] = this.arr[len(this.arr)-1]
    this.arr = this.arr[:len(this.arr)-1]
    this.percolateDown(1)
}

func (this *KthLargest) percolateDown(i int) {
    if i >= len(this.arr) {
        return
    }

    // find min child
    var minchild int
    lchild := 2*i
    rchild := lchild + 1

    if rchild < len(this.arr) {
        // both valid
        if this.arr[lchild] < this.arr[rchild] {
            minchild = lchild
        } else {
            minchild = rchild
        }
    } else if lchild < len(this.arr) {
        minchild = lchild
    } else {
        return
    }

    if this.arr[minchild] < this.arr[i] {
        this.arr[minchild], this.arr[i] = this.arr[i], this.arr[minchild]
        this.percolateDown(minchild)
    }
}

func (this *KthLargest) percolateUp(i int) {
    parentIndx := i / 2
    if parentIndx < 1 {
        return
    }

    if this.arr[parentIndx] > this.arr[i] {
        this.arr[parentIndx], this.arr[i] = this.arr[i], this.arr[parentIndx]
        this.percolateUp(parentIndx)
    } else {
        return
    }
}


func (this *KthLargest) Add(val int) int {
    this.addVal(val)
    return this.arr[1]
}
