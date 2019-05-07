/**
 * Link of the problem
 * @link https://leetcode.com/problems/merge-two-sorted-lists
 */

/**
 * Definition for singly-linked list.
 */
function ListNode(val) {
  this.val = val;
  this.next = null;
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
  let node = null;
  if (!l1 && !l2) {
    return node;
  } else if (!l1) {
    node = new ListNode(l2.val);
    node.next = mergeTwoLists(l1, l2.next);
  } else if (!l2) {
    node = new ListNode(l1.val);
    node.next = mergeTwoLists(l1.next, l2);
  } else if (l1.val > l2.val) {
    node = new ListNode(l2.val);
    node.next = mergeTwoLists(l1, l2.next);
  } else {
    node = new ListNode(l1.val);
    node.next = mergeTwoLists(l1.next, l2);
  }

  return node;
};
