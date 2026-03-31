class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        // resultlist
        // backIt(nums, i, resultList)
        var res = new ArrayList<List<Integer>>();
        backIt(nums, 0, new ArrayList<Integer>(), res);
        return res;
    }

    public void backIt(int[] nums, int i, List<Integer> pickList, List<List<Integer>> res) {
        // if (i >= len(nums)) resList.add(picklist)

        // backIt(nums, i+1, pickList, res)
        // cp = copy(pickList)
        // cp.add(nums[i])
        // backIt(nums, i+1, cp, res)

        if (i >= nums.length) {
            res.add(new ArrayList<>(pickList));
            return;
        }

        pickList.add(nums[i]);
        backIt(nums, i+1, pickList, res);
        pickList.remove(pickList.size() - 1);
        backIt(nums, i+1, pickList, res);
    }
}
