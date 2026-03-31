class Solution {
    public int removeDuplicates(int[] nums) {
        int unique = 1;
        for(int j = 1; j < nums.length; j++) {
            if(nums[j] != nums[j-1]) {
                nums[unique++] = nums[j];
            }
        }
        return unique;
    }
}