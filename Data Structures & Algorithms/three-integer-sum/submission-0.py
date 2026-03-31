class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        
        res_arr = []
        for i in range(len(nums)):
            if i > 0 and nums[i-1] == nums[i]:
                continue
            
            lptr, rptr = i+1, len(nums)-1
            while lptr < rptr:
                if rptr < len(nums) - 1 and nums[rptr+1] == nums[rptr]:
                    rptr -= 1
                    continue
                
                summation = nums[i] + nums[lptr] + nums[rptr]
                if summation == 0:
                    res_arr.append([nums[i], nums[lptr], nums[rptr]])
                    rptr -= 1
                    lptr += 1
                elif summation > 0:
                    rptr -= 1
                elif summation < 0:
                    lptr += 1

        return res_arr