class LRUCache {
    HashMap<Integer, Node> hmap = new HashMap<>();
    Node head;
    Node tail;
    int capacity;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        Node head = new Node(0, 0);
        Node tail = new Node(0, 0);
        head.next = tail;
        tail.prev = head;
        this.head = head;
        this.tail = tail;
    }

    public int get(int key) {
        if (hmap.containsKey(key)) {
            Node hit = hmap.get(key);
            remove(hit);
            insert(hit);
            return hit.value;
        }
        return -1;
    }

    void remove(Node node) {
        Node prev = node.prev;
        Node next = node.next;
        prev.next = next;
        next.prev = prev;
        hmap.remove(node.key);
    }

    void insert(Node node) {
        Node next = head.next;
        next.prev = node;
        node.next = next;
        node.prev = head;
        head.next = node;
        hmap.put(node.key, node);
    }

    public void put(int key, int value) {
        if (hmap.containsKey(key)) {
            remove(hmap.get(key));
        }

        Node newNode = new Node(key, value);
        insert(newNode);

        if (hmap.size() > capacity) {
            remove(tail.prev);
        }
    }
}

class Node {
    int key;
    int value;
    Node next;
    Node prev;

    Node(int key, int value) {
        this.key = key;
        this.value = value;
        this.next = null;
        this.prev = null;
    }
}