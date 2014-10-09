#Reverse digits of an integer.
#Example1: x = 123, return 321
#Example2: x = -123, return -321
class Solution:
    # @return an integer
    def reverse(self, x):
        if x>=0:
            i=1
        else:
            i=-1
        s=str(abs(x))
        return i*int(s[::-1])
