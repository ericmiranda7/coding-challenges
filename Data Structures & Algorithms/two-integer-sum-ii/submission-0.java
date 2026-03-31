class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int left = 0;
        int right = numbers.length - 1;

        int tot = 0;
        while(left < right) {
            tot = numbers[left] + numbers[right];
            if(tot == target) {
                return new int[]{left+1, right+1};
            }

            if(target < tot) {
                right--;
            } else {
                left++;
            }
        }
        return new int[]{0,0};
    }
}
