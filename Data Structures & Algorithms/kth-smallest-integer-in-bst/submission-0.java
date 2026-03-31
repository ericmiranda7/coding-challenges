/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

class Solution {
    public int kthSmallest(TreeNode root, int k) {
        int[] res = new int[2];

        trav(root, k, res);

        return res[1];
    }

    void trav(TreeNode root, int k, int[] res) {
        if (root == null) return;

        trav(root.left, k, res);
        res[0] += 1;
        if (res[0] == k) {
            System.out.println("done for " + root.val);
            res[1] = root.val;
            return;
        }
        trav(root.right, k, res);

        return;
    }

    
}
