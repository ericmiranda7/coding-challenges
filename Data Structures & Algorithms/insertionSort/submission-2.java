// Definition for a pair
// class Pair {
//     int key;
//     String value;
//
//     Pair(int key, String value) {
//         this.key = key;
//         this.value = value;
//     }
// }
public class Solution {
    public List<List<Pair>> insertionSort(List<Pair> pairs) {
        if (pairs.size() == 0) return new ArrayList();
        List<List<Pair>> res = new ArrayList<>();
        res.add(new ArrayList<>(pairs));

        for (int i = 1; i < pairs.size(); i++) {
            Pair p = pairs.get(i);

            int j = i-1;
            while ( j >= 0 && p.key < pairs.get(j).key) {
                // swap
                var tmp = pairs.get(j);
                pairs.set(j, pairs.get(j+1));
                pairs.set(j+1, tmp);
                j--;
            }

            res.add(new ArrayList<>(pairs));
        }

        return res;
    }
}
