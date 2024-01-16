class CarPooling {
    public boolean carPooling(int[][] trips, int capacity) {
        int[] lengthOfTrip = new int[1001];
        for (int[] trip : trips) {
            lengthOfTrip[trip[1]] += trip[0];
            lengthOfTrip[trip[2]] -= trip[0];
        }

        int carLoad = 0;
        for (int i = 0; i < 1001; i++) {
            carLoad += lengthOfTrip[i];
            if (carLoad > capacity) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        CarPooling main = new CarPooling();
        boolean result = main.carPooling(new int[][]{
                {3, 2, 7},
                {3, 7, 9},
                {8, 3, 9},
        }, 11);
        System.out.println(result);
    }
}
