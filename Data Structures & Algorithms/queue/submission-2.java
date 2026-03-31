class Deque {
    Node head;
    Node tail;
    int size;

    public Deque() {
        this.head = new Node(-1, null, null);
        this.tail = this.head;
        this.size = 0;
    }

    public boolean isEmpty() {
        if (size == 0) return true;
        return false;
    }

    public void append(int value) {
        tail.right = new Node(value, tail, null);
        tail = tail.right;
        size += 1;
    }

    public void appendleft(int value) {
        head.right = new Node(value, head, head.right);
        if (head == tail) tail = head.right;    
        else head.right.right.left = head.right;
        size += 1;
    }

    public int pop() {
        if (size == 0) return -1;

        var tmp = tail;
        tail = tail.left;
        size -= 1;

        return tmp.value;
    }

    public int popleft() {
        if (size == 0) return -1;

        var tmp = head.right;
        if (head.right == tail) tail = head;
        head.right = head.right.right;

        size -= 1;

        return tmp.value;
    }
}

class Node {
    int value;
    Node right;
    Node left;

    public Node(int value, Node left, Node right) {
        this.value = value;
        this.right = right;
        this.left = left;
    }
}
