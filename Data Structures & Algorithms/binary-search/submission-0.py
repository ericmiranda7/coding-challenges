class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return self.binary_search(nums, target, 0, len(nums) - 1)

    def binary_search(self, num_arr, target, start, end):
        mid = (start + end) // 2
        if end < start:
            return -1
        if num_arr[mid] == target:
            return mid

        if target < num_arr[mid]:
            return self.binary_search(num_arr, target, start, mid - 1)
        else:
            return self.binary_search(num_arr, target, mid + 1, end)