"""
" Link of the problem
" https://leetcode.com/problems/jewels-and-stones
"""

class Solution:
    def numJewelsInStones(self, J: str, S: str) -> int:
        total = 0
        for l in J:
            have = [stone for stone in S if stone == l]
            total = total + len(have)
            
        return total
