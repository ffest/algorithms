import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;

class ConcatenaitedWords {
    public List<String> findAllConcatenatedWordsInADict(String[] words) {
        Arrays.sort(words, (a,b) -> a.length() - b.length());

        List<String> result = new ArrayList<>();
        HashSet<String> usedWords = new HashSet<>();
        for (String word : words ) {
            if (canForm(word, usedWords)) {
                result.add(word);
            }
            usedWords.add(word);
        }

        return result;
    }

    private boolean canForm(String word, HashSet<String> usedWords) {
        if (usedWords.isEmpty()) {
            return false;
        }
        boolean[] dp = new boolean[word.length()+1];
        dp[0] = true;
        for (int i = 1; i <= word.length(); i++) {
            for (int j = 0; j < i; j++) {
                if (dp[j] && usedWords.contains(word.substring(j, i))) {
                    dp[i] = true;
                    break;
                }
            }
        }
        return dp[word.length()];
    }

    public static void main(String[] args) {
        ConcatenaitedWords main = new ConcatenaitedWords();

        List<String> result = main.findAllConcatenatedWordsInADict(new String[]{
                "cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"
        });
        System.out.println(result);
    }
}
