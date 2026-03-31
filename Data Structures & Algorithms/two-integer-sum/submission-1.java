class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> set = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int need = target - nums[i];
            if (set.containsKey(need)) return new int[]{Math.min(i, set.get(need)), Math.max(i, set.get(need))};
            set.put(nums[i], i);
        }

        return new int[2];
    }
}
