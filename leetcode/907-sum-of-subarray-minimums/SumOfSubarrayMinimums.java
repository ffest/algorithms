import java.util.Stack;

class SumOfSubarrayMinimums {
    public int sumSubarrayMins(int[] arr) {
        int count = 0;
        int mod = (int) 1e9 +7;

        int[] ple = new int[arr.length];
        int[] nle = new int[arr.length];

        Stack<Integer> stackPLE = new Stack<>();
        for (int i = 0; i < arr.length; i++) {
            while (!stackPLE.empty() && arr[stackPLE.peek()] > arr[i]) {
                stackPLE.pop();
            }
            if (stackPLE.empty()) {
                ple[i] = i + 1;
            } else {
                ple[i] = i - stackPLE.peek();
            }
            stackPLE.push(i);
        }
        Stack<Integer> stackNLE = new Stack<>();
        for (int i = arr.length-1; i >= 0; i--) {
            while (!stackNLE.empty() && arr[stackNLE.peek()] >= arr[i]) {
                stackNLE.pop();
            }
            if (stackNLE.empty()) {
                nle[i] = arr.length - i;
            } else {
                nle[i] = stackNLE.peek() - i;
            }
            stackNLE.push(i);
        }

        for (int i = 0; i < arr.length; i++) {
            long add = ((long) arr[i] * nle[i] * ple[i]) % mod;
            count += (int) add;
            count %= mod;
        }

        return count;
    }

    public static void main(String[] args) {
        SumOfSubarrayMinimums main = new SumOfSubarrayMinimums();
        int result = main.sumSubarrayMins(new int[]{3, 1, 2 ,4});
        System.out.println(result);
    }
}
