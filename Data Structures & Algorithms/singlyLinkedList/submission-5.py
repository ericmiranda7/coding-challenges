class LinkedList:
    
    def __init__(self):
        self.head = None
        self.size = 0

    
    def get(self, index: int) -> int:
        if index > self.size - 1:
            return -1
        
        tmp = self.head
        for i in range(index):
            tmp = tmp.next_ptr
        return tmp.value

    def insertHead(self, val: int) -> None:
        new_node = Node(val)
        new_node.next_ptr = self.head
        self.head = new_node
        self.size += 1
        

    def insertTail(self, val: int) -> None:
        tmp = self.head
        while tmp != None and tmp.next_ptr != None:
            tmp = tmp.next_ptr
        
        new_node = Node(val)
        if tmp == None:
            self.head = new_node
        else:
            tmp.next_ptr = new_node
        self.size += 1
        

    def remove(self, index: int) -> bool:
        if index > self.size - 1:
            return False
        
        prev_node = self.head
        for i in range(index - 1):
            prev_node = prev_node.next_ptr

        if index == 0:
            if self.size > 1:
                self.head = self.head.next_ptr
            else:
                self.head = None
        else:
            prev_node.next_ptr = prev_node.next_ptr.next_ptr
        
        self.size -= 1
        return True
        

    def getValues(self) -> List[int]:
        if self.size == 0:
            return []
        res = []
        tmp = self.head
        while tmp != None and tmp.next_ptr != None:
            res.append(tmp.value)
            tmp = tmp.next_ptr
        res.append(tmp.value)
        return res
        
class Node:
    def __init__(self, value):
        self.next_ptr = None
        self.value = value