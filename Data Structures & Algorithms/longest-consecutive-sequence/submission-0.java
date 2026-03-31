class Solution {
    public int longestConsecutive(int[] nums) {
        Set<Integer> set = new HashSet();
        int longest = 0;
        for(int i: nums) {
            set.add(i);
        }

        for(int i: nums) {
            int currLongest = 0;
            int start = i;
            while(!set.contains(start - 1) || set.contains(start) ) {
                start++;
                currLongest++;
            }
            longest = Math.max(longest, currLongest);
        }
        return longest;
    }
}
