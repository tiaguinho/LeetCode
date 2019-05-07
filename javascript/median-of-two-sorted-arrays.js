/**
 * Link of the problem
 * @link https://leetcode.com/problems/median-of-two-sorted-arrays
 */

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
  let a = 0;
  let b = 0;
  let res = [];
  const size = nums1.length + nums2.length;
  const middle = Math.ceil(size / 2);
  for (let i = 0; i <= middle; i++) {
    if (nums1[a] == null) {
      res.push(nums2[b]);
      b++;
    } else if (nums2[b] == null) {
      res.push(nums1[a]);
      a++;
    } else if (nums1[a] < nums2[b]) {
      res.push(nums1[a]);
      a++;
    } else {
      res.push(nums2[b]);
      b++;
    }

    if (i === middle) {
      if (size % 2 === 0) {
        return (res[middle - 1] + res[middle]) / 2;
      } else {
        return res[middle - 1];
      }
    }
  }
};
