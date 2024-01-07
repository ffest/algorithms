import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Set;

class WordLadder {
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Set<String> set = new HashSet<>(wordList);
        if (!set.contains(endWord)) {
            return 0;
        }
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);

        int length = 1;
        while (!queue.isEmpty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                String current = queue.poll();
                for (int j = 0; j < current.length(); j++){
                    for (char letter = 'a'; letter <= 'z'; letter++) {
                        StringBuilder newWord = new StringBuilder(current);
                        newWord.setCharAt(j, letter);
                        if (newWord.toString().equals(endWord)) {
                            return length+1;
                        }
                        if (!set.contains(newWord.toString()) || newWord.toString().equals(current)) {
                            continue;
                        }
                        set.remove(newWord.toString());
                        queue.offer(newWord.toString());
                    }
                }
            }
            length++;
        }

        return 0;
    }

    public static void main(String[] args) {
        WordLadder main = new WordLadder();

        int result = main.ladderLength("hit", "cog", List.of("hot","dot","dog","lot","log","cog"));
        System.out.println(result);
    }
}
