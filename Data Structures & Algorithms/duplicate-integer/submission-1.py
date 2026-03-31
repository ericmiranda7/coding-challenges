class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        ht = dict()

        for n in nums:
            if n in ht:
                return True
            ht[n] = True
        
        return False