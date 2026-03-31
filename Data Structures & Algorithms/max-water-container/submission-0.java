class Solution {
    public int maxArea(int[] heights) {
        int max = 0;

        int left = 0;
        int right = heights.length - 1;

        int curr = 0;
        while(left < right) {
            curr = (right - left) * Math.min(heights[left], heights[right]);
            max = Math.max(max, curr);
            if(heights[left] <= heights[right]){
                left++;
            } else {
                right--;
            }
        }
        return max;
    }
}
