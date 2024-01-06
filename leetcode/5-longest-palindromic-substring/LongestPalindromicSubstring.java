class LongestPalindromicSubstring {
    public String longestPalindrome(String s) {
        String longestPalindrome = "";
        for (int i = 0; i < s.length(); i++){
            String palindrome = expand(i, i, s);
            longestPalindrome = palindrome.length() > longestPalindrome.length() ? palindrome : longestPalindrome;
            palindrome = expand(i, i+1, s);
            longestPalindrome = palindrome.length() > longestPalindrome.length() ? palindrome : longestPalindrome;
        }
        return longestPalindrome;
    }

    private String expand(int left, int right, String s) {
        String palindrome = "";
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            palindrome = s.substring(left, right+1);
            left--;
            right++;
        }
        return palindrome;
    }

    public static void main(String[] args) {
        LongestPalindromicSubstring main = new LongestPalindromicSubstring();

        String result = main.longestPalindrome("abcabcbb");
        System.out.println(result);
    }
}
