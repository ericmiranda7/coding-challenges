type KthLargest struct {
    arr []int
    k int
}


func Constructor(k int, nums []int) KthLargest {
    arr := []int{}
    arr = append(arr, nums...)
    return KthLargest{arr, k}
}


func (this *KthLargest) Add(val int) int {
    this.arr = append(this.arr, val)
    sort.Ints(this.arr)
    return this.arr[len(this.arr) - this.k]
}
