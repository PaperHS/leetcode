class Solution:
    # @return a float
    def findMedianSortedArrays(self, A, B):
        C = self.MergeQue(A,B)
        ret =0.0
        if len(C)%2 == 1:
            return C[len(C)//2]
        else:
            return (C[len(C)//2]+C[len(C)//2-1])/float(2)
        
        
    def MergeQue(self,left,right):
        if len(left)==0:
            return right
        if len(right)==0:
            return left
        result = []
        i=0
        j=0
        while i < len(left) and j < len(right):
            if left[i] <= right[j]:
                result.append(left[i])
                i+=1
            else:
                result.append(right[j])
                j+=1
        result+=left[i:]
        result+=right[j:]
        return result
