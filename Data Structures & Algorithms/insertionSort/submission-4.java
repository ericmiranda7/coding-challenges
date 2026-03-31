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
        List<List<Pair>> ans = new ArrayList();
        for (int i = 0; i < pairs.size(); i++) {
            int j = i - 1;
            while (j >= 0 && pairs.get(j + 1).key < pairs.get(j).key) {
                var temp = pairs.get(j + 1);
                pairs.set(j + 1, pairs.get(j));
                pairs.set(j, temp);
                j--;
            }
            ans.add(new ArrayList<>(pairs));
        }
        return ans;
    }
}
