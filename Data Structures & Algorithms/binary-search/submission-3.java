class Solution {
    public int search(int[] nums, int target) {
        return binSearch(nums, 0, nums.length - 1, target);
    }

    int binSearch(int[] nums, int start, int end, int target) {
        // base
        if (start > end) return - 1;

        int mid = (start + end) / 2;
        
        if (target > nums[mid]) start = mid + 1;
        else if (target < nums[mid]) end = mid - 1;
        else return mid;

        return binSearch(nums, start, end, target);
    }
}
