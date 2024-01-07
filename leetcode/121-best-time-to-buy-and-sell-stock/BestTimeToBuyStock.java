import java.util.Map;

class BestTimeToBuyStock {
    public int maxProfit(int[] prices) {
        if (prices.length < 2) {
            return 0;
        }
        int min = prices[0];
        int maxProfit = 0;
        for (int i = 1; i < prices.length; i++) {
            maxProfit = Math.max(maxProfit, prices[i]-min);
            min = Math.min(min, prices[i]);
        }
        return maxProfit;
    }

    public static void main(String[] args) {
        BestTimeToBuyStock main = new BestTimeToBuyStock();

        int result = main.maxProfit(new int[]{1, 2, 3, 4, 5, 6});
        System.out.println(result);
    }
}
