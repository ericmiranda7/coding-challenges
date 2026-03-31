class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
            
        store = [0] * 26

        for i in range(len(s)):
            store[ord(s[i]) - ord('a')] += 1
            store[ord(t[i]) - ord('a')] -= 1
            print(store)
        
        for n in store:
            if n != 0:
                return False

        return True