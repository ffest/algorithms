class BestTimeToBuyStockII {
    public int maxProfit(int[] prices) {
        int profit = 0;
        for (int i = 1; i < prices.length; i++) {
            if (prices[i] > prices[i-1]) {
                profit += prices[i]-prices[i-1];
            }
        }
        return profit;
    }

    public static void main(String[] args) {
        BestTimeToBuyStockII main = new BestTimeToBuyStockII();

        int result = main.maxProfit(new int[]{1, 2, 8, 4, 5, 6});
        System.out.println(result);
    }
}
