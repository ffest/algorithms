import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class FindAllDuplicatesInAnArray {
    public List<Integer> findDuplicates(int[] nums) {
        List<Integer> result = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            int index = Math.abs(nums[i])-1;
            if (nums[index] < 0) {
                result.add(index+1);
            }
            nums[index] = -nums[index];
        }
        return result;
    }

    public static void main(String[] args) {
        FindAllDuplicatesInAnArray main = new FindAllDuplicatesInAnArray();

        List<Integer> result = main.findDuplicates(new int[]{4,3,2,7,8,2,3,1});
        System.out.println(result);
    }
}
