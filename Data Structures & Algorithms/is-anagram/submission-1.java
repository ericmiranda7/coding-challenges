class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;

        int[] freq1 = new int[128];
        int[] freq2 = new int[128];
        Arrays.fill(freq1, 0);
        Arrays.fill(freq2, 0);

        for (int i = 0; i < s.length(); i++) {
            freq1[Character.getNumericValue(s.charAt(i))] += 1;
            freq2[Character.getNumericValue(t.charAt(i))] += 1;
        }

        for (int i = 0; i < t.length(); i++) {
            if (freq1[Character.getNumericValue(s.charAt(i))]
                    != freq2[Character.getNumericValue(s.charAt(i))]) return false;
        }

        return true;
    }
}
