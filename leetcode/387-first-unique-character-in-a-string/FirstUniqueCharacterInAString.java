import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;

class FirstUniqueCharacterInAString {
    public int firstUniqCharNaive(String s) {
        int index = -1;
        for(char ch : s.toCharArray()) {
            if(s.indexOf(ch) == s.lastIndexOf(ch)) {
                return s.indexOf(ch);
            }
        }

        return index;
    }

    public int firstUniqChar(String s) {
        int[] cache = new int[26];
        for (char c : s.toCharArray()) {
            cache[c-'a']++;
        }

        for (int i = 0; i < s.length(); i++) {
            if (cache[s.charAt(i)-'a'] == 1) {
                return i;
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        FirstUniqueCharacterInAString main = new FirstUniqueCharacterInAString();

        int result = main.firstUniqChar("leetcode");
        System.out.println(result);
    }
}
