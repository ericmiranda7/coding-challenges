class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> hm = new HashMap<>();
        int[] res = new int[2];

        for (int i = 0; i < nums.length; i++) {
            hm.put(nums[i], i);
        }

        for (int i = 0; i < nums.length; i++) {
            int need = target - nums[i];

            if (hm.containsKey(need) && hm.get(need) != i) {
                int found = hm.get(need);
                res[0] = Math.min(i, found);
                res[1] = Math.max(i, found);
            }
        }

        return res;
    }
}
