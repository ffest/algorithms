import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

class ValidParentheses {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();

        Map<Character, Character> closes = new HashMap<>();
        closes.put(')', '(');
        closes.put('}', '{');
        closes.put(']', '[');

        for (int i = 0; i < s.length(); i++) {
            Character current = s.charAt(i);
            if (closes.containsKey(current)) {
                if (stack.empty() || stack.pop() != closes.get(current)) {
                    return false;
                }
            } else {
                stack.push(current);
            }
        }
        return stack.empty();
    }

    public static void main(String[] args) {
        ValidParentheses main = new ValidParentheses();
        boolean result = main.isValid("{[]}");
        System.out.println(result);
    }
}
