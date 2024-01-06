import java.util.HashMap;
import java.util.Map;


class TwoSum {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> numToIndex = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (numToIndex.containsKey(target-nums[i])) {
                return new int[]{numToIndex.get(target-nums[i]), i};
            }
            numToIndex.put(nums[i], i);
        }
        return new int[]{};
    }

    public static void main(String[] args) {
        TwoSum twoSum = new TwoSum();
        int[] nums = {2, 7, 11, 15};
        int target = 9;

        int[] result = twoSum.twoSum(nums, target);

        if (result.length == 2) {
            System.out.println("Indices: [" + result[0] + ", " + result[1] + "]");
        } else {
            System.out.println("No solution found.");
        }
    }
}
