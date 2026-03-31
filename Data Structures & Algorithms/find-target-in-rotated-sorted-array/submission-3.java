class Solution {
    public int search(int[] nums, int target) {
        // bin search
        // get midpoint
        // if target == midpoint return
        // find out if left half or right half
            // left half - start < mid
            // right half - end > mid
        // if lh
            // if target > mid
                // need to go right
                // start = mid + 1
            // if target < mid
                // need to go left or rh
                // if target > start
                    // end = mid - 1
                // if target < start
                    // start = mid + 1
        // if rh
            // if target > mid
                // need to go right
                // if target > end
                    // end = mid - 1
                // if target < end
                    // start = mid + 1
            // if target < mid
                // need to go left
                

        int start = 0;
        int end = nums.length - 1;

        while (start <= end) {
            int mid = (start + end) / 2;

            if (target == nums[mid]) return mid;

            if (nums[0] <= nums[mid]) {
                // left half
                if (target > nums[mid]) start = mid + 1;
                else {
                    if (target >= nums[start]) end = mid - 1;
                    else start = mid + 1;
                }
            } else {
                // right half
                if (target > nums[mid]) {
                    if (target > nums[end]) end = mid - 1;
                    else start = mid + 1;
                } else {
                    end = mid - 1;
                }
            }
        }

        return - 1;
    }
}