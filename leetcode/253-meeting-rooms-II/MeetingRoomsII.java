import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

class MeetingRoomsII {
    public int minMeetingRooms(int[][] intervals) {
        ArrayList<Integer> starts = new ArrayList<>();
        ArrayList<Integer> ends = new ArrayList<>();
        for (int[] interval : intervals) {
            starts.add(interval[0]);
            ends.add(interval[1]);
        }
        Collections.sort(starts);
        Collections.sort(ends);
        int rooms = 0;
        int tail = 0;
        for (Integer start : starts) {
            if (start < ends.get(tail)) {
                rooms++;
            } else {
                tail++;
            }
        }
        return rooms;
    }

    public static void main(String[] args) {
        MeetingRoomsII main = new MeetingRoomsII();

        int result = main.minMeetingRooms(new int[][]{
                {0, 30},
                {5, 10},
                {15, 20},
                {6, 11},
        });
        System.out.println(result);
    }
}
