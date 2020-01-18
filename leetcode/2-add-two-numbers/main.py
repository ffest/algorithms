# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        result = result_iter = ListNode(0)

        current = 0
        while l1 or l2:
            if l1:
                current += l1.val
                l1 = l1.next
            if l2:
                current += l2.val
                l2 = l2.next

            result_iter.next = ListNode(current % 10)
            result_iter = result_iter.next
            current //= 10
        if current == 1:
            result_iter.next = ListNode(1)

        return result.next


solution = Solution()
l1 = ListNode(2, ListNode(4, ListNode(3)))
l2 = ListNode(5, ListNode(6, ListNode(4)))
print(solution.addTwoNumbers(l1, l2))
