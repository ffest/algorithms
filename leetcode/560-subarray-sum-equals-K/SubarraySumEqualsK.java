import java.util.HashMap;

class SubarraySumEqualsK {
    public int subarraySum(int[] nums, int k) {
        int count = 0, sum = 0;
        HashMap<Integer, Integer> prefixSums = new HashMap<>();
        prefixSums.put(0, 1);

        for (int num : nums) {
            sum += num;
            if (prefixSums.containsKey(sum - k)) {
                count += prefixSums.get(sum - k);
            }
            prefixSums.put(sum, prefixSums.getOrDefault(sum, 0) + 1);
        }

        return count;
    }

    public static void main(String[] args) {
        SubarraySumEqualsK main = new SubarraySumEqualsK();

        int result = main.subarraySum(new int[]{1,2,3}, 3);
        System.out.println(result);
    }
}
