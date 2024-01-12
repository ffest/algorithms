import java.util.Arrays;
import java.util.Comparator;

class LongestIncreasingSubsequence {
    public int lengthOfLIS(int[] nums) {
        int[] tails = new int[nums.length];
        int size = 0;
        for (int num : nums) {
            int i = 0, j = size;
            while (i != j) {
                int mid = (i+j)/2;
                if (tails[mid] < num) {
                    i = mid+1;
                } else {
                    j = mid;
                }
            }
            tails[i] = num;
            if (i == size) {
                size++;
            }
        }
        return size;
    }

    public int lengthOfLISDP(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int[] dp = new int[nums.length];
        Arrays.fill(dp, 1);

        int size = 1;
        for (int i = 1; i < nums.length; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i]> nums[j]) {
                    dp[i] = Math.max(dp[j]+1, dp[i]);
                }
                size = Math.max(size, dp[i]);
            }
        }

        return size;
    }

    public static void main(String[] args) {
        LongestIncreasingSubsequence main = new LongestIncreasingSubsequence();

        int result = main.lengthOfLIS(new int[]{10,9,2,5,3,7,101,18});
        System.out.println(result);
    }
}
