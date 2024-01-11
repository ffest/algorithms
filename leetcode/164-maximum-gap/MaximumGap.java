import java.util.Arrays;

class MaximumGap {
    public int maximumGap(int[] nums) {
        if (nums == null || nums.length < 2) {
            return 0;
        }

        int min = nums[0];
        int max = nums[0];
        for (int i : nums) {
            min = Math.min(min, i);
            max = Math.max(max, i);
        }

        int gapSize = (int)Math.ceil((double)(max - min) / (nums.length - 1));
        int[] bucketsMIN = new int[nums.length - 1];
        int[] bucketsMAX = new int[nums.length - 1];
        Arrays.fill(bucketsMIN, Integer.MAX_VALUE);
        Arrays.fill(bucketsMAX, Integer.MIN_VALUE);

        for (int i: nums) {
            if (i == min || i == max) {
                continue;
            }
            int bucketIdx = (i - min)/gapSize;
            bucketsMIN[bucketIdx] = Math.min(bucketsMIN[bucketIdx], i);
            bucketsMAX[bucketIdx] = Math.max(bucketsMAX[bucketIdx], i);
        }

        int maxGap = Integer.MIN_VALUE;
        int previous = min;

        for (int i = 0; i < nums.length-1; i++) {
            if (bucketsMIN[i] == Integer.MAX_VALUE && bucketsMAX[i] == Integer.MIN_VALUE) {
                continue;
            }
            maxGap = Math.max(maxGap, bucketsMIN[i] - previous);
            previous = bucketsMAX[i];
        }

        maxGap = Math.max(maxGap, max - previous);
        return maxGap;
    }

    public static void main(String[] args) {
        MaximumGap main = new MaximumGap();

        int result = main.maximumGap(new int[]{3,6,9,1});
        System.out.println(result);
    }
}
