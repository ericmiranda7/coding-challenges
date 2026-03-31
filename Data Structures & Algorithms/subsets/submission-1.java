class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> out = new ArrayList<>();
        List<Integer> currentSet = new ArrayList<>();

        generateSubsets(nums, 0, out, currentSet);

        return out;
    }

    void generateSubsets(int[] nums,
                        int i,
                        List<List<Integer>> out,
                        List<Integer> currentSet) {
        // base
        if (i == nums.length) {
            out.add(currentSet.stream().toList());
            return;
        }

        // dont include
        generateSubsets(nums, i+1, out, currentSet);

        // include num
        List<Integer> cpy = currentSet.stream().collect(Collectors.toList());
        cpy.add(nums[i]);
        generateSubsets(nums, i+1, out, cpy);
    }
}