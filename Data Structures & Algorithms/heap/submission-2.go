type MinHeap struct {
    arr []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{[]int{}}
}

func (mh *MinHeap) Push(val int) {
    mh.arr = append(mh.arr, val)

	fmt.Println("sres")

	newNodeIndx := len(mh.arr) - 1
	for newNodeIndx > 0 {
		pIndx := (newNodeIndx-1) / 2
		if mh.arr[newNodeIndx] < mh.arr[pIndx] {
			mh.arr[newNodeIndx], mh.arr[pIndx] = mh.arr[pIndx], mh.arr[newNodeIndx]
			newNodeIndx = pIndx
		} else {
			break
		}
	}

	fmt.Println("push success")
}

func (mh *MinHeap) Pop() int {
    if len(mh.arr) == 0 {
        return -1
    }

    tmp := mh.arr[0]
    mh.arr[0] = mh.arr[len(mh.arr)-1]
    mh.arr = mh.arr[:len(mh.arr)-1]

    // bubble down
    itr := 0

    for {
		lchild := 2*itr + 1
		rchild := 2*itr + 2
		minIndx := lchild
		if lchild < len(mh.arr) && rchild < len(mh.arr) {
			if mh.arr[lchild] < mh.arr[rchild] {
				minIndx = lchild
			} else {
				minIndx = rchild
			}
		} else if lchild < len(mh.arr) {
			minIndx = lchild
		} else if rchild < len(mh.arr) {
			minIndx = rchild
		} else {
			break
		}

		if mh.arr[itr] > mh.arr[minIndx] {
			mh.arr[itr], mh.arr[minIndx] = mh.arr[minIndx], mh.arr[itr]
			itr = minIndx
		} else {
			break
		}
    }

    return tmp
}

func (mh *MinHeap) Top() int {
	if len(mh.arr) == 0 {
		return -1
	}
    return mh.arr[0]
}

func (mh *MinHeap) Heapify(nums []int) {
    for _, num := range nums {
        mh.Push(num)
    }
}
