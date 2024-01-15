class PalindromicSubstrings {
    public int countSubstrings(String s) {
        int count = 0;
        for (int i = 0; i < s.length(); i++) {
            count += expand(i, i, s);
            count += expand(i, i+1, s);
        }
        return count;
    }

    private int expand(int i, int j, String s) {
        int count = 0;
        while (i >= 0 && j < s.length() && s.charAt(i) == s.charAt(j)) {
            count++;
            i--;
            j++;
        }
        return count;
    }

    public static void main(String[] args) {
        PalindromicSubstrings main = new PalindromicSubstrings();

        int result = main.countSubstrings("abcac");
        System.out.println(result);
    }
}
