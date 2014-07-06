
// Problem:
// Given an array of integers, find two numbers such that they add up to a specific target number.

// The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.

// You may assume that each input would have exactly one solution.

// Input: numbers={2, 7, 11, 15}, target=9
// Output: index1=1, index2=2

public class Solution {
    public int[] twoSum(int[] numbers, int target) {
        TreeMap<Integer,Integer> list = new TreeMap<Integer,Integer>();
        for(int i=0; i<numbers.length; i++){
            if(list.containsKey(numbers[i])){
            	//when numbers has 2 sames num,jugded if they are answer.if not,it doesnt matter
                if(numbers[i]+numbers[i] == target){
                    int[] res ={list.get(numbers[i])+1,i+1};
                    return res;
                }
            }else{
                list.put(numbers[i],i);
                
            }
        }
        int indext1=0,indext2=0;
        for(int j:list.keySet()){
            if(list.containsKey(target-j)){
                indext1=list.get(j);
                indext2=list.get(target-j);
                break;
            }
        } 
        if(indext1 < indext2){
            int[] res = {indext1+1,indext2+1};
            return res;
        }else{
            int[] res = {indext2+1,indext1+1};
            return res;
        }
    }
}