class Solution:
    # @return an integer
    def lengthOfLongestSubstring(self, s):
        c =[]
        max =0
        i = 0
        for x in s:
            if x in c:
                if i > max:
                    max =i
                p = c.index(x)
                i =len(c)-p
                c =c[(p+1):]
                c.append(x)
            else:
                i+=1
                c.append(x)
        if i > max:
            max =i
        return max
