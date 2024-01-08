import java.util.Arrays;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Set;

class FindKthSmallestPairDistance {
    public int smallestDistancePair(int[] nums, int k) {
        Arrays.sort(nums);
        int minDist = 0;
        int maxDist = nums[nums.length-1]-nums[0];

        while (minDist < maxDist) {
            int midDist = (maxDist + minDist) / 2;
            int count = 0;

            int left = 0;
            int right = 0;
            while (right < nums.length) {
                if (nums[right] - nums[left] > midDist) {
                    left++;
                } else {
                    count += right-left;
                    right++;
                }
            }

            if (count < k) {
                minDist = midDist + 1;
            } else {
                maxDist = midDist;
            }
        }

        return minDist;
    }

    public static void main(String[] args) {
        FindKthSmallestPairDistance main = new FindKthSmallestPairDistance();
        int result = main.smallestDistancePair(new int[]{9, 10, 7, 10, 6, 1, 5, 4, 9, 8}, 18);
        System.out.println(result);
    }
}
