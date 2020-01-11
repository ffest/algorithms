# Definition for singly-linked list.
class ListNode:
    def __init__(self, x, next=None):
        self.val = x
        self.next = next


def overlapping_lists(l1: ListNode, l2: ListNode) -> ListNode:
    root1, root2 = detectCycle(l1), detectCycle(l2)
    if not root1 and not root2:
        return overlapping_lists_without_cycles(l1, l2)
    elif (root1 and not root2) or (root2 and not root1):
        return None
    else:
        tmp = root2
        while tmp:
            tmp = tmp.next
            if tmp is root1 or tmp is root2:
                break
        return root2 if tmp is root1 else None

# TODO: write this func 7.4
def overlapping_lists_without_cycles(l1: ListNode, l2: ListNode) -> ListNode:
    return l1


def detectCycle(head: ListNode) -> ListNode:
    fast = slow = head
    while fast and fast.next and fast.next.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            slow = head
            while slow != fast:
                slow = slow.next
                fast = fast.next
            return slow
    return None

cycle = ListNode(10, ListNode(11, ListNode(12)))
cycle.next = cycle
l1 = ListNode(0, ListNode(1, ListNode(2, cycle)))
l2 = ListNode(5, ListNode(6, cycle))
print(overlapping_lists(l1, l2))
