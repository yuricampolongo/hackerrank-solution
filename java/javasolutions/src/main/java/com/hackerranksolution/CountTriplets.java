package com.hackerranksolution;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class CountTriplets {

    static long countTriplets(List<Long> arr, long r) {
        Map<Long, Long> leftMap = new HashMap<>();
        Map<Long, Long> rightMap = new HashMap<>();
        
        // Populate rightMap with frequency of elements
        for (long num : arr) {
            rightMap.put(num, rightMap.getOrDefault(num, 0L) + 1);
        }
        
        long count = 0;
        
        // Iterate through the array and count triplets
        for (int i = 0; i < arr.size(); i++) {
            long midTerm = arr.get(i);
            long countLeft = 0;
            long countRight = 0;
            
            rightMap.put(midTerm, rightMap.get(midTerm) - 1);
            
            if (leftMap.containsKey(midTerm / r) && midTerm % r == 0) {
                countLeft = leftMap.get(midTerm / r);
            }
            
            if (rightMap.containsKey(midTerm * r)) {
                countRight = rightMap.get(midTerm * r);
            }
            
            count += countLeft * countRight;
            
            leftMap.put(midTerm, leftMap.getOrDefault(midTerm, 0L) + 1);
        }
        
        return count;
    }
}
