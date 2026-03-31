class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        self.genPar(n, "", 0, 0, res)
        return res
    
    def genPar(self, n, stack, op, cl, res):
        if op == n and cl == n:
            res.append(stack)
        
        a, b = None, None
        if op < n:
            self.genPar(n, stack + '(', op + 1, cl, res)
        if cl < n and cl < op:
            self.genPar(n, stack + ')', op, cl + 1, res)

        return a, b