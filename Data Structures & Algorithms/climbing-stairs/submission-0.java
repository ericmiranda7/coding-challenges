class Solution {
    public int climbStairs(int n) {
        if (n == 2 || n == 1) return n;

        return climbStairs(n-1) + climbStairs(n-2);
    }
}
