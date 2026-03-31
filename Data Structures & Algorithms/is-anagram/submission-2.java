class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;

        int[] sa = new int[26];
        int[] ta = new int[26];

        for (char c : s.toCharArray()) {
            sa[(int) c - 97] += 1;
        }
        for (char c : t.toCharArray()) {
            ta[(int) c - 97] += 1;
        }

        for (int i = 0; i < sa.length; i++) {
            if (sa[i] != ta[i]) return false;
        }

        return true;
    }
}
