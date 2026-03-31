class Node {
    int key;
    int value;
    Node next;
    Node prev;

    public Node(int key, int value) {
        this.key = key;
        this.value = value;
    }
}

class LRUCache {
    int capacity;
    Map<Integer, Node> cache = new HashMap();
    Node left;
    Node right;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.left = new Node(0, 0);
        this.right = new Node(0, 0);
        this.left.next = this.right;
        this.right.prev = this.left;
    }
    
    public int get(int key) {
        if(cache.containsKey(key)) {
            remove(cache.get(key));
            insert(cache.get(key));
            return cache.get(key).value;
        }
        return -1;
    }
    
    public void put(int key, int value) {
        if(cache.containsKey(key)) {
            remove(cache.get(key));
        }
        Node newNode = new Node(key, value);
        cache.put(key, newNode);
        insert(newNode);

        if(cache.size() > capacity) {
            Node lru = this.left.next;
            remove(lru);
            cache.remove(lru.key);
        }
    }

    void insert(Node node) {
        var mru = this.right.prev;
        mru.next = node;
        node.prev = mru;
        node.next = this.right;
        this.right.prev = node;
    }

    void remove(Node node) {
        node.next.prev = node.prev;
        node.prev.next = node.next;
    }
}
