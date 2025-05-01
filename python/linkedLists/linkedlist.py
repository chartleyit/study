#!/bin/env python

class Node():
    def __init__(self, data):
        self.next = None
        self.data = data

class LinkedList():
    def __init__(self):
        self.head = None

    def print_list(self):
        curr_node = self.head
        while curr_node:
            print(curr_node.data)
            curr_node = curr_node.next

    def append(self, data):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
            return
        last_node = self.head
        while last_node.next:
            last_node = last_node.next
        last_node.next = new_node
    
    def prepend(self, data):
        new_node = Node(data)
        new_node.next = self.head
        self.head = new_node
    
    def len(self):
        curr_node = self.head
        count = 0
        while curr_node:
            count += 1
            curr_node = curr_node.next
        return(count)
        

if __name__ == '__main__':
    ll1 = LinkedList()
    for x in [3,4,5]:
        ll1.append(x)

    print("len:", ll1.len())

    for y in [9,8,7]:
        ll1.prepend(y)

    print("len:", ll1.len())

    ll1.print_list()
