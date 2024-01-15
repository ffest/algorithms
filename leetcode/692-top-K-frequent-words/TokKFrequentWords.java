import java.util.ArrayList;
import java.util.Comparator;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Set;

class TokKFrequentWords {
    public List<String> topKFrequent(String[] words, int k) {
        Map<String, Integer> cache = new HashMap<>();
        for(String word : words){
            cache.put(word, cache.getOrDefault(word, 0)+1);
        }
        MyComparator comparator = new MyComparator();
        PriorityQueue<Map.Entry<String, Integer>> mh = new PriorityQueue<>(comparator);
        for (Map.Entry<String, Integer> entry : cache.entrySet()) {
            if (mh.size() < k) {
                mh.offer(entry);
            } else if (comparator.compare(mh.peek(), entry) < 0) {
                mh.poll();
                mh.offer(entry);
            }
        }

        List<String> result = new ArrayList<>();
        for(int i = 0; i < k; i++){
            result.add(0, mh.poll().getKey());//the "smaller" entry poll out ealier
        }
        return result;
    }

    static class MyComparator implements Comparator<Map.Entry<String, Integer>> {
        public int compare(Map.Entry<String, Integer> e1, Map.Entry<String, Integer> e2){
            if (e1.getValue() != e2.getValue()) {
                return e1.getValue() - e2.getValue();
            } else {
                return e2.getKey().compareTo(e1.getKey());
            }
        }
    }

    public static void main(String[] args) {
        TokKFrequentWords main = new TokKFrequentWords();
        List<String> result = main.topKFrequent(new String[]{"i","love","leetcode","i","love","coding"}, 1);
        System.out.println(result);
    }
}
