class Solution:
    def trap(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1

        maxl, maxr = height[l], height[r]

        trappable = 0
        while l != r:
            if height[l] <= height[r]:
                l += 1
                curr_ind = l
            else:
                r -= 1
                curr_ind = r
            
            current_trap = min(maxl, maxr) - height[curr_ind]
            trappable += 0 if current_trap < 0 else current_trap

            if height[l] > maxl:
                maxl = height[l]
            if height[r] > maxr:
                maxr = height[r]

        return trappable