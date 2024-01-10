import java.util.HashMap;
import java.util.Map;


class WordSearch {
    public boolean exist(char[][] board, String word) {
        if (board.length == 0 || board[0].length == 0) {
            return false;
        }
        char[] wordChars = word.toCharArray();
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (dfs(i, j, wordChars, board, 0)) {
                    return true;
                }
            }
        }

        return false;
    }

    private boolean dfs(int i, int j, char[] wordChars, char[][] board, int k) {
        if (k == wordChars.length) {
            return true;
        }
        if (i < 0 || j < 0 || i == board.length || j == board[0].length || board[i][j] != wordChars[k]) {
            return false;
        }
        char tmp = board[i][j];
        board[i][j] = '#';
        boolean dfs = dfs(i+1, j, wordChars, board, k+1) ||
                dfs(i-1, j, wordChars, board, k+1) ||
                dfs(i, j+1, wordChars, board, k+1) ||
                dfs(i, j-1, wordChars, board, k+1);
        board[i][j] = tmp;

        return dfs;
    }

    public static void main(String[] args) {
        WordSearch twoSum = new WordSearch();
        char[][] board = {{'C', 'A','A'}, {'A', 'A', 'A'}, {'B', 'C', 'B'}};

        boolean result = twoSum.exist(board, "AAB");
        System.out.println(result);
    }
}
