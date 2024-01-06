class AddTwoNumbers {
    public static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode tail = new ListNode(0);
        ListNode head = tail;
        int num = 0;

        while (l1 != null || l2 != null) {
            num /= 10;
            if (l1 != null) {
                num += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                num += l2.val;
                l2 = l2.next;
            }

            head.next = new ListNode(num % 10);
            head = head.next;
        }

        if (num / 10 == 1) {
            head.next = new ListNode(1);
        }
        return tail.next;
    }

    public static void main(String[] args) {
        ListNode list1 = new ListNode(2, new ListNode(4, new ListNode(3)));
        ListNode list2 = new ListNode(5, new ListNode(6, new ListNode(4)));
        AddTwoNumbers main = new AddTwoNumbers();

        ListNode result = main.addTwoNumbers(list1, list2);
        System.out.println(result);
    }
}
