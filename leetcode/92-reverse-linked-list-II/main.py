# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next=None):
        self.val = x
        self.next = next


class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        result = ListNode(0, head)
        left = result
        for i in range(m-1):
            left = left.next
        middle, right = left.next, left.next.next

        for i in range(n-m):
            middle.next = right.next
            right.next = left.next
            left.next = right
            right = middle.next

        return result.next


solution = Solution()
head = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
print(solution.reverseBetween(head, 2, 4))
