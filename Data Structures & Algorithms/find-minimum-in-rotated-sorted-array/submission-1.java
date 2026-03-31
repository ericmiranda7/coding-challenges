class Solution {
    public int findMin(int[] nums) {
        // binsearch
        // conditions:
            // rval < mid: binsearch(mid+1) (right)
            // else: binsearch(,mid-1) (left)

        int start = 0;
        int end = nums.length - 1;

        int min = nums[((end - start) / 2) + start];
        while (start <= end) {
            if (nums[start] < min && nums[start] <= nums[end]) {
                min = nums[start];
                break;
            }
            int mid = ((end - start) / 2) + start;
            
            if (nums[mid] < min) {
                min = nums[mid];
            }
            
            if (nums[end] <= nums[mid]) {
                start = mid + 1;
            } else {
                end = mid - 1;
            }
        }

        return min;
    }
}
