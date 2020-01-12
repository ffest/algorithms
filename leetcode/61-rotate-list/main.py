# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next = None):
        self.val = x
        self.next = next


class Solution:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if head is None:
            return None

        length, tail = 1, head
        while tail.next:
            length += 1
            tail = tail.next
        tail.next = head
        steps = length - k % length
        while steps:
            steps -= 1
            tail = tail.next
        head = tail.next
        tail.next = None

        return head


solution = Solution()
head = ListNode(0, ListNode(1, ListNode(2)))
print(solution.rotateRight(head, 4).val)
