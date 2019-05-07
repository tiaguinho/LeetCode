"""
" Link to the problem
" https://leetcode.com/problems/to-lower-case/
"""

class Solution:
    def toLowerCase(self, str: str) -> str:
        lower = ""
        for s in str:
            if s >= 'A' and s <= 'Z':
                lower += str(s + 32)
            else:
                lower += str(s)
        
        return lower
