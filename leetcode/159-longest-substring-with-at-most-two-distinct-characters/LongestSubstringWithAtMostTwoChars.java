import java.util.HashMap;

class LongestSubstringWithAtMostTwoChars {
    public int lengthOfLongestSubstringTwoDistinct(String s) {
        int j = 0;
        int max = 0;

        HashMap<Character, Integer> map = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            char currentChar = s.charAt(i);
            map.put(currentChar, map.getOrDefault(s.charAt(i), 0)+1);
            while (map.size() > 2) {
                int currentNumber = map.get(s.charAt(j));
                map.put(s.charAt(j), currentNumber-1);
                if (currentNumber == 1) {
                    map.remove(s.charAt(j));
                }
                j++;
            }
            if (i - j + 1 > max) {
                max = i - j + 1;
            }
        }
        return max;
    }

    public static void main(String[] args) {
        LongestSubstringWithAtMostTwoChars main = new LongestSubstringWithAtMostTwoChars();

        int result = main.lengthOfLongestSubstringTwoDistinct("ccaabbb");
        System.out.println(result);
    }
}
