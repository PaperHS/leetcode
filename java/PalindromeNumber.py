class Solution:
    # @return a boolean
    def isPalindrome(self, x):
        s = str(x)
        l = len(s)
        s1=s[:l//2]
        if l%2 ==0:
            s2=s[l//2:]
        else:
            s2=s[l//2+1:]
        if s1==s2[::-1]:
            return True
        else:
            return False
