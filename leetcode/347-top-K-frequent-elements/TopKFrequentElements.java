import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;

class TopKFrequentElements {
    public int[] topKFrequentMinHeap(int[] nums, int k) {
        Map<Integer, Integer> cache = new HashMap<>();
        for (int num : nums) {
            cache.put(num, cache.getOrDefault(num, 0)+1);
        }

        PriorityQueue<Map.Entry<Integer, Integer>> minHeap = new PriorityQueue<>(Comparator.comparingInt(Map.Entry::getValue));
        for(Map.Entry<Integer,Integer> entry: cache.entrySet()){
            if (minHeap.size() < k) {
                minHeap.offer(entry);
            } else if (minHeap.peek().getValue() < entry.getValue()){
                minHeap.poll();
                minHeap.offer(entry);
            }
        }
        List<Integer> res = new ArrayList<>();
        while (!minHeap.isEmpty()) {
            res.add(minHeap.poll().getKey());
        }

        return res.stream().mapToInt(Integer::intValue).toArray();
    }

    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> cache = new HashMap<>();
        for (int num : nums) {
            cache.put(num, cache.getOrDefault(num, 0)+1);
        }

        List<Integer>[] buckets = new List[nums.length +1];
        for (int key : cache.keySet()) {
            if (buckets[cache.get(key)] == null) {
                buckets[cache.get(key)] = new ArrayList<>();
            }
            buckets[cache.get(key)].add(key);
        }

        List<Integer> result = new ArrayList();
        for (int i = buckets.length - 1; i >= 0; i--) {
            if (buckets[i] == null) {
                continue;
            }
            for (int j : buckets[i]) {
                if (result.size() < k) {
                    result.add(j);
                }
            }
        }

        return result.stream().mapToInt(Integer::intValue).toArray();
    }

    public static void main(String[] args) {
        TopKFrequentElements main = new TopKFrequentElements();

        int[] result = main.topKFrequent(new int[]{1,1,2,5,5,5}, 2);
        System.out.println(Arrays.toString(result));
    }
}
