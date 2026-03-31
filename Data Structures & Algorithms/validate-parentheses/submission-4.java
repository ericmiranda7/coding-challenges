class Solution {
    Map<Character, Character> parMap = Map.of(
        ')', '(',
        '}', '{',
        ']', '['
    );

    public boolean isValid(String s) {
        // for each c
            // if c is op
                // add to stack
            // else
                // if popped != opp(c)
                    // return false

        List<Character> par = new ArrayList<>();
        Set<Character> ops = Set.of('(', '[', '{');

        for (char c : s.toCharArray()) {
            if (ops.contains(c)) par.add(c);
            else {
                if (par.size() == 0) return false;
                char p = par.remove(par.size() - 1);
                if (p != parMap.get(c)) return false;
            }
        }

        return par.size() == 0;
    }
}
