class CoinChange {
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount+1];
        for (int i = 1; i <= amount; i++) {
            int min = -1;
            for (int coin : coins) {
                if (i < coin || dp[i-coin] == -1) {
                    continue;
                }
                if (min == -1 || dp[i-coin]+1 < min) {
                    min = dp[i-coin]+1;
                }
            }
            dp[i] = min;
        }
        return dp[amount];
    }

    public static void main(String[] args) {
        CoinChange main = new CoinChange();

        int result = main.coinChange(new int[]{1,2,5}, 8);
        System.out.println(result);
    }
}
