# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        prev: ListNode = None
        while head:
            tmp = head.next
            head.next = prev
            prev = head
            head = tmp

        return prev


solution = Solution()
list = ListNode(1)
next = ListNode(5)
next.next = ListNode(9)
list.next = next
print(solution.reverseList(list))
