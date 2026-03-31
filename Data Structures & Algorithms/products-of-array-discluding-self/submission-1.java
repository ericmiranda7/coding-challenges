class Solution {
    public int[] productExceptSelf(int[] nums) {
       // 1   2     4    6

       // 1   1     2    8
       // 48  24    6    8

    
       int[] ans = new int[nums.length];
       ans[0] = 1;
       for(int i = 1; i < nums.length; i++) {
            ans[i] =  ans[i-1] * nums[i-1] ;
       }

       int product = 1;
       for(int i = nums.length - 1; i >=0; i--) {
        ans[i] *= product;
        product *= nums[i];
       }

    //    System.out.println(Arrays.toString(ans));

       return ans;
    }
}  
