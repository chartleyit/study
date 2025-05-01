#!/bin/env python

class ListNode():
    def __init__(self, val=0, next=None):
        self.value = val
        self.next = next

def createList(list: list[int]) -> ListNode:
    # error handling for empty list
    if list[0] == 0:
        return None
    head = ListNode(list[0])
    new_head = head

    # start at 1 because list initialized with 0
    for i in range(1, len(list)):
        new_head.next = ListNode(list[i])
        new_head = new_head.next
    
    return head

def iterateList(ll: ListNode):
    while ll:
        print(ll.value)
        ll.next

def addLinkedLists(ll1: ListNode, ll2: ListNode) -> ListNode:
    new_list = ListNode()
    new_list.next = None

if __name__ == "__main__":
    ll1 = createList([1,2,3]) # reverse list 321
    ll2 = createList([5,3,9]) # reverse list 935

    iterateList(ll1)
    iterateList(ll2)