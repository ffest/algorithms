class FindMinimumInRotatedSortedArray {
    public int findMin(int[] nums) {
        int left = 0;
        int right = nums.length - 1;
        while (left < right) {
            int mid = (left + right) / 2;
            if (nums[mid] > nums[right]) {
                left = mid+1;
            } else {
                right = mid;
            }
        }
        return nums[left];
    }

    public static void main(String[] args) {
        FindMinimumInRotatedSortedArray main = new FindMinimumInRotatedSortedArray();

        int result = main.findMin(new int[]{4,5,6,7,0,1,2});
        System.out.println(result);
    }
}
