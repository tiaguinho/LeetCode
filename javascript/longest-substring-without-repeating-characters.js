/**
 * Link of the problem
 * @link https://leetcode.com/problems/longest-substring-without-repeating-characters
 */

/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
  let lastLongest = 0;
  let substring = '';
  let i = 0;
  while (i < s.length) {
    var index = substring.indexOf(s[i]);
    if (index !== -1) {
      if (substring.length > lastLongest) {
        lastLongest = substring.length;
      }

      substring = substring.slice(index + 1);
    }

    substring += s[i];
    i++;
  }

  return substring.length > lastLongest ? substring.length : lastLongest;
};
