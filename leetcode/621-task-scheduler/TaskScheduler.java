import java.util.HashMap;

class TaskScheduler {
    public int leastInterval(char[] tasks, int n) {
        int maxTasks = 0;
        int maxTaskCount = 0;
        HashMap<Character, Integer> cache = new HashMap<>();
        for (char task: tasks) {
            cache.put(task, cache.getOrDefault(task, 0)+1);
            if (cache.get(task) > maxTasks) {
                maxTaskCount = 1;
                maxTasks = cache.getOrDefault(task, 0);
            } else {
                maxTaskCount += 1;
            }
        }

        int idles = maxTasks - 1;
        int idleSize = n - maxTaskCount + 1;
        int maxTaskSpace = maxTasks * maxTaskCount;

        return Math.max(tasks.length, maxTaskSpace+idles*idleSize);
    }

    public static void main(String[] args) {
        TaskScheduler main = new TaskScheduler();

        int result = main.leastInterval(new char[]{'A','A','A','B','B','B'}, 0);
        System.out.println(result);
    }
}
