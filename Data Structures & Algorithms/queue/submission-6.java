class Node {
    int val;
    Node prev;
    Node next;

    public Node(int val) {
        this(val, null, null);
    }

    public Node(int val, Node prev, Node next) {
        this.val = val;
        this.prev = prev;
        this.next = next;
    }
}

class Deque {
    Node head;
    Node tail;

    public Deque() {
        head = new Node(-1);
        tail = new Node(-1);

        head.next = tail;
        tail.prev = head;
    }

    public boolean isEmpty() {
        return head.next == tail;
    }

    public void append(int value) {
       var newNode = new Node(value);
       newNode.next = tail;
       newNode.prev = tail.prev;
       tail.prev = newNode;
       tail.prev.prev.next = newNode;
    }

    public void appendleft(int value) {
        var newNode = new Node(value);
        newNode.prev = head;
        newNode.next = head.next;
        head.next = newNode;
        head.next.next.prev = newNode;
    }

    public int pop() {
        if(head.next == tail) {
            return -1;
        }

        var pop = tail.prev;
        pop.prev.next = tail;
        tail.prev = pop.prev;

        return pop.val;
    }

    public int popleft() {
        if(head.next == tail) return -1;

        var pop = head.next;

        pop.next.prev = head;
        head.next = pop.next;
        
        return pop.val;
    }
}
