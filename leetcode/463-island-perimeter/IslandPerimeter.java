class IslandPerimeter {
    public int islandPerimeter(int[][] grid) {
        int lands = 0, neighbours = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] != 1) {
                    continue;
                }
                lands++;
                if (i != grid.length-1 && grid[i+1][j] == 1) {
                    neighbours++;
                }
                if (j != grid[0].length-1 && grid[i][j+1] == 1) {
                    neighbours++;
                }
            }
        }
        return 4*lands - 2*neighbours;
    }

    public static void main(String[] args) {
        IslandPerimeter main = new IslandPerimeter();

        int result = main.islandPerimeter(new int[][]{
                {0,1,0,0},
                {1,1,1,0},
                {0,1,0,0},
                {1,1,0,0},
        });
        System.out.println(result);
    }
}
