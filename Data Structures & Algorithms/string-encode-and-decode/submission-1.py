class Solution:
    def encode(self, strs: List[str]) -> str:
        encoded_str = ""
        for s in strs:
            encoded_str = encoded_str + str(len(s)) + "#" + s

        return encoded_str

    def decode(self, s: str) -> List[str]:
        res_arr = []

        i = 0
        while i < len(s):
            len_in_str = ""
            while s[i] != "#":
                len_in_str += s[i]
                i += 1
            str_len = int(len_in_str)
            
            word = res_arr.append(s[i+1:i+1+str_len])
            i = i+1+str_len

        return res_arr
            
            