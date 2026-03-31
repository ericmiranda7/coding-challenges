type MinHeap struct {
    arr []int
}

func NewMinHeap() *MinHeap {
    fmt.Println("returned heap")
    return &MinHeap{[]int{-1}}
}

func (mh *MinHeap) Push(val int) {
    fmt.Println("pushing")
    mh.arr = append(mh.arr, val)

    for percIndx := len(mh.arr) - 1; percIndx > 0; {
        parentIndx := percIndx / 2
        if parentIndx <= 0 {
            break
        }

        if mh.arr[percIndx] < mh.arr[parentIndx] {
            // percolate up
            mh.arr[percIndx], mh.arr[parentIndx] = mh.arr[parentIndx], mh.arr[percIndx]
        } else {
            // done percolating to the top
            break
        }
        percIndx = parentIndx
    }
}

func (mh *MinHeap) Pop() int {
    if len(mh.arr) <= 1 {
        return -1
    }
    res := mh.arr[1]
    fmt.Println("popping", res)

    mh.arr[1] = mh.arr[len(mh.arr)-1]
    mh.arr = mh.arr[0:len(mh.arr)-1]
    fmt.Println("after pop", mh.arr)

    for i := 1; i < len(mh.arr); {
        var minChildIndx int
        lchild := 2*i
        rchild := lchild + 1

        if lchild < len(mh.arr) && rchild < len(mh.arr) {
            if mh.arr[lchild] <= mh.arr[rchild] {
                minChildIndx = lchild
            } else {
                minChildIndx = rchild
            }
        } else if lchild < len(mh.arr) {
            minChildIndx = lchild
        } else if rchild < len(mh.arr) {
            minChildIndx = rchild
        } else {
            break
        }
        fmt.Println("minchildINdx", minChildIndx)

        if mh.arr[i] > mh.arr[minChildIndx] {
            mh.arr[i], mh.arr[minChildIndx] = mh.arr[minChildIndx], mh.arr[i]
            i = minChildIndx
        } else {
            break
        }
    }

    return res
}

func (mh *MinHeap) Top() int {
    if len(mh.arr) <= 1 {
        return -1
    }
    fmt.Println("topped", mh.arr)
    return mh.arr[1]
}

func (mh *MinHeap) Heapify(nums []int) {
    mh.arr = append(mh.arr, nums...)

    for i := len(mh.arr) - 1; i > 0; {
        var minChildIndx int
        lchild := 2*i
        rchild := lchild + 1

        if lchild < len(mh.arr) && rchild < len(mh.arr) {
            if mh.arr[lchild] <= mh.arr[rchild] {
                minChildIndx = lchild
            } else {
                minChildIndx = rchild
            }
        } else if lchild < len(mh.arr) {
            minChildIndx = lchild
        } else if rchild < len(mh.arr) {
            minChildIndx = rchild
        } else {
            i--
            continue
        }

        if mh.arr[i] > mh.arr[minChildIndx] {
            mh.arr[i], mh.arr[minChildIndx] = mh.arr[minChildIndx], mh.arr[i]
            mh.minify(minChildIndx)
        }

        i--
    }

    fmt.Println("heapify", mh.arr)
}

func (mh *MinHeap) minify(i int) {
    if i >= len(mh.arr) {
        return
    }

    var minChildIndx int
    lchild := 2*i
    rchild := lchild + 1

    if lchild < len(mh.arr) && rchild < len(mh.arr) {
        if mh.arr[lchild] <= mh.arr[rchild] {
            minChildIndx = lchild
        } else {
            minChildIndx = rchild
        }
    } else if lchild < len(mh.arr) {
        minChildIndx = lchild
    } else if rchild < len(mh.arr) {
        minChildIndx = rchild
    } else {
        return
    }
    
    if mh.arr[i] > mh.arr[minChildIndx] {
        mh.arr[i], mh.arr[minChildIndx] = mh.arr[minChildIndx], mh.arr[i]
        mh.minify(minChildIndx)
    }
}
