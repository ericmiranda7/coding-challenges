class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        curr_price_ind, j = 0, 1

        max_profit = 0
        while j < len(prices):
            profit = prices[j] - prices[curr_price_ind]
            if profit > max_profit:
                max_profit = profit

            if prices[j] < prices[curr_price_ind]:
                curr_price_ind = j
            
            j += 1
        
        return max_profit
