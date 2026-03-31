class DynamicArray {

    int capacity;
    int[] arr;
    int size;

    public DynamicArray(int capacity) {
        this.capacity = capacity;
        this.arr = new int[capacity];
        this.size = 0;
    }

    public int get(int i) {
        return this.arr[i];
    }

    public void set(int i, int n) {
        this.arr[i] = n;
    }

    public void pushback(int n) {
        if(this.size == this.capacity) {
            this.resize();
        }
        this.arr[this.size++] = n;
    }

    public int popback() {
        this.size--;
        return this.arr[this.size];
    }

    private void resize() {
        this.capacity = 2 * this.capacity;

        var tempArr = this.arr;
        this.arr = new int[this.capacity];

        for(int i = 0; i < tempArr.length; i++) {
            this.arr[i] = tempArr[i];
        }
    }

    public int getSize() {
        return this.size;
    }

    public int getCapacity() {
        return this.capacity;
    }
}
