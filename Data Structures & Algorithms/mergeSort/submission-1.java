// Definition for a pair.
// class Pair {
//     public int key;
//     public String value;
//
//     public Pair(int key, String value) {
//         this.key = key;
//         this.value = value;
//     }
// }
class Solution {
    public List<Pair> mergeSort(List<Pair> pairs) {
        // todo(): use .get().keys
        if (pairs.isEmpty()) return new ArrayList<>();
        return mergeIt(pairs, 0, pairs.size() - 1);
    }

    public List<Pair> mergeIt(List<Pair> pairs, int start, int end) {
        if (start - end == 0) return List.of(pairs.get(start));

        int mid = ((end - start) / 2) + start;
        var leftList = mergeIt(pairs, start, mid);
        var rightList = mergeIt(pairs, mid+1, end);

        List<Pair> res = new ArrayList<>();
        int l = 0;
        int r = 0;

        while (l < leftList.size() && r < rightList.size()) {
            if (leftList.get(l).key <= rightList.get(r).key) {
                res.add(leftList.get(l++));
            } else {
                res.add(rightList.get(r++));
            }
        }

        if (r >= rightList.size()) res.addAll(leftList.subList(l, leftList.size()));
        else res.addAll(rightList.subList(r, rightList.size()));

        return res;
    }
}
