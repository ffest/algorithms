import java.util.Arrays;

class ShortestDistanceToACharacter {
    public int[] shortestToChar(String s, char c) {
        int[] result = new int[s.length()];
        int prev = -1;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == c) {
                for (int j = prev+1; j < i; j++) {
                    if (result[j] == 0 || result[j] > i - j) {
                        result[j] = i - j;
                    }
                }
                result[i] = 0;
                prev = i;
            } else if (prev != -1) {
                result[i] = i - prev;
            }
        }
        return result;
    }

    public static void main(String[] args) {
        ShortestDistanceToACharacter main = new ShortestDistanceToACharacter();
        int[] result = main.shortestToChar("loveleetcode", 'e');
        System.out.println(Arrays.toString(result));
    }
}
