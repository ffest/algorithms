import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Set;

class MaximumProductSubarray {
    public int maxProduct(int[] nums) {
        int result = nums[0];
        int min = 1;
        int max = 1;

        for (int num : nums) {
            if (num < 0) {
                int tmp = min;
                min = max;
                max = tmp;
            }
            min = Math.min(num, num * min);
            max = Math.max(num, num * max);

            if (max > result) {
                result = max;
            }
        }

        return result;
    }

    public static void main(String[] args) {
        MaximumProductSubarray main = new MaximumProductSubarray();

        int result = main.maxProduct(new int[]{-2});
        System.out.println(result);
    }
}
