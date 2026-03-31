class ListNode {
    int val;
    ListNode next;

    public ListNode(int val) {
        this(val, null);
    }

    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class LinkedList {
    ListNode head;
    ListNode tail;

    public LinkedList() {
        this.head = new ListNode(-1);
        this.tail = this.head;
    }

    public int get(int index) {
        var curr = this.head.next;

        int i = 0;
        while(curr != null) {
            if(i == index) {
                return curr.val;
            }
            curr = curr.next;
            i++;
        }
        return -1;
    }

    public void insertHead(int val) {
        var newNode = new ListNode(val, this.head.next);
        this.head.next = newNode;

        if (newNode.next == null) {
            this.tail = newNode;
        }
    }

    public void insertTail(int val) {
        this.tail.next = new ListNode(val);
        this.tail = this.tail.next;
    }

    public boolean remove(int index) {
        int i = 0;
        var curr = this.head;
        while(curr != null && i < index) {
            curr = curr.next;
            i++;
        }
        if(curr != null && curr.next != null) {
            if(curr.next == this.tail) {
                this.tail = curr;
            }
            curr.next = curr.next.next;
            return true;
        }
        return false;
    }

    public ArrayList<Integer> getValues() {
        ArrayList<Integer> list = new ArrayList();
        var curr = this.head.next;
        while(curr != null) {
            list.add(curr.val);
            curr = curr.next;
        }
        return list;
    }
}
