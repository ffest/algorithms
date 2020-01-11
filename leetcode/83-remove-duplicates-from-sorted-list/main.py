# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        dummy = head
        while head and head.next:
            if head.val == head.next.val:
                head.next = head.next.next
            else:
                head = head.next
        return dummy


solution = Solution()
head = ListNode(1, ListNode(2, ListNode(2, ListNode(4, ListNode(4)))))
print(solution.deleteDuplicates(head))
