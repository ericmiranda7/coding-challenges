class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        nums = Arrays.stream(nums).sorted().toArray();
        List<List<Integer>> res = new ArrayList();
        
        for(int i = 0; i< nums.length; i++) {
            if(i != 0 && nums[i - 1] == nums[i]) {
                continue;
            }

            int left = i +1;
            int right = nums.length - 1;

            while(left < right) {
                int currSum = nums[i] + nums[left] + nums[right];
                if(currSum == 0) {
                    res.add(List.of(nums[i], nums[left], nums[right]));
                    left++;
                    while(nums[left] == nums[left - 1] && left < right) {
                        left++;
                    }
                } else if(currSum < 0) {
                    left++;
                } else {
                    right--;
                }
            }
        }
        return res;
    }
}
