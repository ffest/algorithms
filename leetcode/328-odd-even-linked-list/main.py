# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next


class Solution:
    def oddEvenList(self, head: ListNode) -> ListNode:
        if head is None:
            return head

        even_head, odd_head = ListNode(0), ListNode(0)
        even, odd = even_head, odd_head
        step = 0
        while head:
            if step % 2 == 0:
                even.next = head
                even = even.next
            else:
                odd.next = head
                odd = odd.next
            step += 1
            head = head.next
        even.next = odd_head.next
        odd.next = None
        return even_head.next


head = ListNode(0, ListNode(1, ListNode(2, ListNode(3, ListNode(4)))))
solution = Solution()
print(solution.oddEvenList(head).val)
