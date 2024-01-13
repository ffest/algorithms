import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

class FrogJump {
    public boolean canCross(int[] stones) {
        HashMap<Integer, Set<Integer>> cache = new HashMap<>();
        for (int stone : stones) {
            cache.put(stone, new HashSet<>());
        }
        cache.get(stones[0]).add(1);

        for (int stone : stones) {
            for (int jump : cache.get(stone)) {
                int reach = jump + stone;
                if (reach == stones[stones.length-1]) {
                    return true;
                }
                if (cache.containsKey(reach)) {
                    cache.get(reach).add(jump+1);
                    cache.get(reach).add(jump);
                    if (jump > 1) {
                        cache.get(reach).add(jump-1);
                    }
                }
            }
        }

        return false;
    }

    public static void main(String[] args) {
        FrogJump main = new FrogJump();

        boolean result = main.canCross(new int[]{0, 1, 3, 5, 6, 8, 12, 17});
        System.out.println(result);
    }
}
