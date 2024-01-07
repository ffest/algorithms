class BestTimeToBuyStockIII {
    public int maxProfit(int[] prices) {
        int firstBuy = Integer.MIN_VALUE;
        int secondBuy = Integer.MIN_VALUE;
        int firstSell = 0;
        int secondSell = 0;
        for (int price : prices) {
            secondSell = Math.max(secondSell, secondBuy + price);
            secondBuy = Math.max(secondBuy, firstSell - price);
            firstSell = Math.max(firstSell, firstBuy + price);
            firstBuy = Math.max(firstBuy, -price);
        }
        return secondSell;
    }

    public int maxProfitDP(int[] prices) {
        if (prices.length == 0) return 0;
        int[][] dp = new int[3][prices.length];
        for (int k = 1; k <= 2; k++) {
            int min = prices[0];
            for (int i = 1; i < prices.length; i++) {
                min = Math.min(min, prices[i] - dp[k-1][i-1]);
                dp[k][i] = Math.max(dp[k][i-1], prices[i] - min);
            }
        }
        return dp[2][prices.length-1];
    }

    public static void main(String[] args) {
        BestTimeToBuyStockIII main = new BestTimeToBuyStockIII();

        int result = main.maxProfitDP(new int[]{1, 2, 8, 4, 5, 6});
        System.out.println(result);
    }
}
