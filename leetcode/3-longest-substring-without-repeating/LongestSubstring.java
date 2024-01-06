class LongestSubstring {
    public int lengthOfLongestSubstring(String s) {
        boolean[] cache = new boolean[128];
        int max = 0;
        int length = 0;
        int tail = 0;
        for (int i = 0; i < s.length(); i++){
            while (cache[s.charAt(i)]) {
                cache[s.charAt(tail)] = false;
                length--;
                tail++;
            }
            cache[s.charAt(i)] = true;
            length++;
            if (max < length) {
                max = length;
            }
        }

        return max;
    }

    public static void main(String[] args) {
        LongestSubstring main = new LongestSubstring();

        int result = main.lengthOfLongestSubstring("abcabcbb");
        System.out.println(result);
    }
}
