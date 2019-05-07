/**
 * Link of the problem
 * @link https://leetcode.com/problems/add-two-numbers/submissions/
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
var addTwoNumbers = function(l1, l2) {
  return sumTwoNumbers(l1, l2, 0);
};

const sumTwoNumbers = (l1, l2, diff) => {
  let value = diff;
  if (l1) value += l1.val;
  if (l2) value += l2.val;

  const l3 = new ListNode(value);
  if (l3.val >= 10) {
    l3.val = l3.val % 10;
    diff = 1;
  } else {
    diff = 0;
  }

  if (l1 && l1.next && l2 && l2.next) {
    l3.next = sumTwoNumbers(l1.next, l2.next, diff);
  } else if (l1 && l1.next) {
    l3.next = sumTwoNumbers(l1.next, null, diff);
  } else if (l2 && l2.next) {
    l3.next = sumTwoNumbers(null, l2.next, diff);
  } else if (diff == 1) {
    l3.next = new ListNode(diff);
  }

  return l3;
};
