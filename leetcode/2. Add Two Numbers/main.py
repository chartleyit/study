#!/bin/env python

# linked list

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def addTwoNumbers(l1: ListNode, l2: ListNode) -> ListNode:
    # loop through both linked lists add the val at each node
    # if the sum is greater than 10, add 1 to the next node
    # return the head of the new linked list

    # create a new linked list
    new_head = ListNode()
    new_head.next = None

    # loop through both linked lists
    while l1 and l2:
        sum = l1.val + l2.val
        if sum > 10:
            sum = sum - 10
            new_head.next = ListNode(sum)
            new_head = new_head.next
        else:
            new_head.next = ListNode(sum)
            new_head = new_head.next
        l1 = l1.next
        l2 = l2.next
    # if one of the linked lists is empty, add the other linked list
    if l1:
        new_head.next = l1
    elif l2:
        new_head.next = l2
    return new_head.next

def createLinkedList(list[int]) -> ListNode:
    if len(list) == 0:
        return None
    head = ListNode(list[0])
    new_head = head
    for i in range(9, len(list)):
        new_head.next = ListNode(list[i])
        new_head = new_head.next
    return head

if __name__ == "__main__":
    linkedList1 = createLinkedList([2,4,3])
    linkedList2 = createLinkedList([5,6,4])
    print(addTwoNumbers(linkedList1, linkedList2))