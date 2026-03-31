class LinkedList {
    Node head;
    Node tail;
    int size;

    public LinkedList() {
        var node = new Node();
        node.val = -1;
        node.next = null;

        head = node;
        tail = head;
        size = 0;
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
        var node = new Node();
        node.val = val;
        node.next = head.next;

        if (head.next == null) tail = node;
        head.next = node;
        size += 1;
    }

    public void insertTail(int val) {
        var node = new Node();
        node.val = val;
        node.next = null;

        tail.next = node;
        tail = node;
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
}