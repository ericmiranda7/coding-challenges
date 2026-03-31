class LinkedList {
    int size;
    Node head;
    Node tail;


    public LinkedList() {
        this.size = 0;
        this.head = new Node(-1, null);
        this.tail = head;
    }

    public int get(int index) {
        if (index >= size) return -1;

        var curr = head.next;
        while (index > 0) {
            curr = curr.next;
            index -= 1;
        }

        return curr.val;
    }

    public void insertHead(int val) {
        var newNode = new Node(val, head.next);

        if (head == tail) tail = newNode;
        head.next = newNode;

        size += 1;
    }

    public void insertTail(int val) {
        var newNode = new Node(val);

        tail.next = newNode;
        tail = newNode;

        size += 1;
    }

    public boolean remove(int index) {
        if (index >= size) return false;

        var curr = head;
        while (index > 0) {
            curr = curr.next;
            index -= 1;
        }
        if (curr.next == tail) tail = curr;
        curr.next = curr.next.next;

        size -= 1;
        return true;
    }

    public ArrayList<Integer> getValues() {
        ArrayList<Integer> res = new ArrayList<>();

        var curr = head.next;
        while (curr != null) {
            res.add(curr.val);
            curr = curr.next;
        }

        return res;
    }
}

class Node {
    int val;
    Node next;

    public Node(int val, Node next) {
        this.val = val;
        this.next = next;
    }

    public Node(int val) {
        this(val, null);
    }
}