import java.util.Arrays;
import java.util.Comparator;

class MeetingRooms {
    public boolean canAttendMeetings(int[][] nums) {
        Arrays.sort(nums, Comparator.comparingInt(interval -> interval[0]));

        int finish = nums[0][1];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i][0] < finish) {
                return false;
            }
            finish = nums[i][1];
        }
        return true;
    }

    public static void main(String[] args) {
        MeetingRooms main = new MeetingRooms();

        boolean result = main.canAttendMeetings(new int[][]{
                {0, 30},
                {5, 10},
                {15, 20},
        });
        System.out.println(result);
    }
}
