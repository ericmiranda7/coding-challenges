class Solution:
    def __init__(self):
        self.arr = []

    def encode(self, strs: List[str]) -> str:
        res_str = ""
        for s in strs:
            self.arr.append(len(s))
            res_str += s

        return res_str

    def decode(self, s: str) -> List[str]:
        res_arr = []
        i = 0
        for word_len in self.arr:
            res_arr.append(s[i:i+word_len])
            i += word_len
        return res_arr
