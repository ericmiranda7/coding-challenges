/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

class Solution {
    public ListNode reverseList(ListNode head) {
        // head ptr, tail ptr
        // while head != tail
        // bkp = head.next
        // head.next = tail.next
        // tail.next = head
        // head = bkp

        if (head == null) {
            return null;
        }

        ListNode tail = head;
        while (tail.next != null) {
            tail = tail.next;
        }

        while (head != tail) {
            ListNode bkp = head.next;
            head.next = tail.next;
            tail.next = head;
            head = bkp;
        }

        return tail;
    }
}
