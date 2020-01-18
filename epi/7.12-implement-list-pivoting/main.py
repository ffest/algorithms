# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next=None):
        self.val = x
        self.next = next


def list_pivoting(l: ListNode, x: int) -> ListNode:
    less_head = less = ListNode(0)
    equal_head = equal = ListNode(0)
    greater_head = greater = ListNode(0)

    while l.next:
        if l.val < x:
            less.next = l
            less = less.next
        elif l.val > x:
            greater.next = l
            greater = greater.next
        else:
            equal.next = l
            equal = equal.next
        l = l.next
    less.next = equal_head.next
    equal.next = greater_head.next
    return less_head.next

l = ListNode(3, ListNode(2, ListNode(2, ListNode(11, ListNode(7, ListNode(5, ListNode(11)))))))
print(list_pivoting(l, 7))
