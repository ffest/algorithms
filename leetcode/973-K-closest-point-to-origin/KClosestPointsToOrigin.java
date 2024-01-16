import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;

class KClosestPointsToOrigin {
    public int[][] kClosest(int[][] points, int k) {
        if (points.length <= k) {
            return points;
        }

        PriorityQueue<int[]> pq = new PriorityQueue<>(k, (a, b) -> (b[0]*b[0] + b[1]*b[1]) - (a[0]*a[0] + a[1]*a[1]));
        for (int[] point: points) {
            pq.add(point);
            if (pq.size() > k) {
                pq.poll();
            }
        }

        return pq.toArray(new int[][]{});
    }

    public static void main(String[] args) {
        KClosestPointsToOrigin main = new KClosestPointsToOrigin();
        int[][] result = main.kClosest(new int[][]{
                {3, 3},
                {5, -1},
                {-2, 4}
        }, 2);
        System.out.println(Arrays.deepToString(result));
    }
}
