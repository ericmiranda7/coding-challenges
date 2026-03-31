class DynamicArray {
    int capacity;
    int length;
    int[] arr;

    public DynamicArray(int capacity) {
        this.capacity = capacity;
        this.arr = new int[capacity];
        this.length = 0;
    }

    public int get(int i) {
        return arr[i];
    }

    public void set(int i, int n) {
        arr[i] = n;
    }

    public void pushback(int n) {
        if (length == capacity) resize();
        arr[length++] = n;
    }

    public int popback() {
        return arr[--length];
    }

    private void resize() {
        capacity *= 2;
        int[] newArr = new int[capacity];

        for (int i = 0; i < arr.length; i++) {
            newArr[i] = arr[i];
        }
        arr = newArr;
    }

    public int getSize() {
        return length;
    }

    public int getCapacity() {
        return capacity;
    }
}
