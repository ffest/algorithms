class HouseRobber {
    public int rob(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        int[] dp = new int[nums.length+1];
        dp[0] = 0;
        dp[1] = nums[0];
        for (int i = 1; i < nums.length; i++) {
            dp[i+1] = Math.max(dp[i], dp[i-1]+nums[i]);
        }
        return dp[nums.length];
    }

    public static void main(String[] args) {
        HouseRobber main = new HouseRobber();

        int result = main.rob(new int[]{4,5,6,7,0,1,2});
        System.out.println(result);
    }
}
