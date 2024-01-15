import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class OpenTheLock {
    public int openLock(String[] deadends, String target) {
        Set<String> deadendSet = new HashSet<>(List.of(deadends));
        Set<String> visited = new HashSet<>();

        List<String> queue = new ArrayList<>();
        queue.add("0000");
        int length = 0;
        while (!queue.isEmpty()) {
            int queueSize = queue.size();
            for (int i = 0; i < queueSize; i++) {
                String current = queue.get(0);
                queue.remove(0);
                if (deadendSet.contains(current)) {
                    continue;
                }
                if (current.equals(target)) {
                    return length;
                }
                for (int j = 0; j < 4; j++) {
                    String next = getNextOrPrev(current.substring(0, j), current.substring(j+1),
                            current.charAt(j), false);
                    if (!visited.contains(next)) {
                        queue.add(next);
                        visited.add(next);
                    }
                    String prev = getNextOrPrev(current.substring(0, j), current.substring(j+1),
                            current.charAt(j), true);
                    if (!visited.contains(prev)) {
                        queue.add(prev);
                        visited.add(prev);
                    }
                }
            }
            length++;
        }
        return -1;
    }

    private String getNextOrPrev(String prefix, String suffix, char c, boolean prev) {
        if (prev) {
            if (c == '0') {
                c = '9';
            } else {
                c--;
            }
        } else {
            if (c == '9') {
                c = '0';
            } else {
                c++;
            }
        }
        return prefix + c + suffix;
    }

    public static void main(String[] args) {
        OpenTheLock main = new OpenTheLock();
        int result = main.openLock(new String[]{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
                "8888");
        System.out.println(result);
    }
}
