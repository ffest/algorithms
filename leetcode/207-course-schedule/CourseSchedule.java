import java.util.ArrayList;

class CourseSchedule {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        if (numCourses < 2 || prerequisites.length == 0) {
            return true;
        }
        ArrayList[] adjMatrix = new ArrayList[numCourses];
        for (int i = 0; i < numCourses; i++) {
            adjMatrix[i] = new ArrayList<>();
        }

        boolean[] visited = new boolean[numCourses];
        boolean[] using = new boolean[numCourses];
        for (int i = 0; i < prerequisites.length; i++) {
            int want = prerequisites[i][0];
            int need = prerequisites[i][1];
            adjMatrix[want].add(need);
        }

        for (int i = 0; i < numCourses; i++){
            if (visited[i]) {
                continue;
            }
            if (hasCycles(i, adjMatrix, visited, using)) {
                return false;
            }
        }

        return true;
    }

    private boolean hasCycles(int subj, ArrayList[] adjMatrix, boolean[] visited, boolean[] using) {
        if (using[subj]) {
            return true;
        }

        ArrayList needs = adjMatrix[subj];
        if (needs.isEmpty()) {
            visited[subj] = true;
            return false;
        }


        for (int i = 0; i < needs.size(); i++){
            using[subj] = true;
            if (hasCycles((int) needs.get(i), adjMatrix, visited, using)) {
                return true;
            }
        }
        visited[subj] = true;
        using[subj] = false;
        return false;
    }

    public static void main(String[] args) {
        CourseSchedule main = new CourseSchedule();

        boolean result = main.canFinish(3, new int[][]{
                {1, 0},
                {2, 0}
        });
        System.out.println(result);
    }
}
