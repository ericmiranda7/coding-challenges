class Solution {
    public boolean isValidSudoku(char[][] board) {
        Map<Integer, Set<Character>> rows = new HashMap();
        Map<Integer, Set<Character>> cols = new HashMap();
        Map<Integer, Set<Character>> box = new HashMap();

        for(int row = 0; row < 9; row++) {
            for(int col = 0; col < 9; col++) {
                if(board[row][col] == '.') continue;

                int boxKey = row / 3 * 3 + col / 3;
                if(
                    !rows.computeIfAbsent(row, i -> new HashSet()).add(board[row][col]) ||
                    !cols.computeIfAbsent(col, i -> new HashSet()).add(board[row][col]) ||
                    !box.computeIfAbsent(boxKey, i -> new HashSet()).add(board[row][col])
                ) {
                    return false;
                }
            }
        }
        return true;
    }
}
