# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        dummy = ListNode(0, head)
        first, second = dummy, dummy
        for i in range(n):
            first = first.next
        while first.next:
            second = second.next
            first = first.next
        second.next = second.next.next

        return dummy.next


solution = Solution()
head = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
print(solution.removeNthFromEnd(head, 4))
