class DynamicArray {

    int ogCapacity;
    int length;
    int[] arr;

    public DynamicArray(int capacity) {
        ogCapacity = capacity;
        arr = new int[ogCapacity];
    }

    public int get(int i) {
        return arr[i];
    }

    public void set(int i, int n) {
        if(i < length) {
            arr[i] = n;
            return;
        }
        return;
    }

    public void pushback(int n) {
        if(length == ogCapacity) {
            resize();
        }
        arr[length] = n;
        length++;
    }

    public int popback() {
        length--;
        return arr[length];

    }

    private void resize() {
        ogCapacity = 2 * ogCapacity;
        int[] newArr = new int[ogCapacity];

        for (int i=0; i < length; i++) {
            newArr[i] = arr[i];
        }
        arr = newArr;
    }

    public int getSize() {
        return length;
    }

    public int getCapacity() {
        return ogCapacity;
    }
}
