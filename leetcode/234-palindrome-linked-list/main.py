# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next


class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        if head is None or head.next is None:
            return True

        fast = slow = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        prev = None
        while slow:
            tmp = slow.next
            slow.next = prev
            prev = slow
            slow = tmp

        while prev:
            if prev.val != head.val:
                return False
            prev = prev.next
            head = head.next

        return True





solution = Solution()
head = ListNode(1, ListNode(2, ListNode(3, ListNode(2, ListNode(1)))))
print(solution.isPalindrome(head))
