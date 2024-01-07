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

    public static void main(String[] args) {
        BestTimeToBuyStockIII main = new BestTimeToBuyStockIII();

        int result = main.maxProfit(new int[]{1, 2, 8, 4, 5, 6});
        System.out.println(result);
    }
}
