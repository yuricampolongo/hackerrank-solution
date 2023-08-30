package com.hackerranksolution;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class FrequencyQueries {
    // Complete the freqQuery function below.
    static List<Integer> freqQuery(List<List<Integer>> queries) {
        HashMap<Integer, Integer> frequencies = new HashMap<>();
        int[] quantities = new int[queries.size() + 1];
        List<Integer> result = new ArrayList<>();
        int frequency, value;
        for (List<Integer> query : queries) {
            value = query.get(1);
            switch (query.get(0)) {
            case 1:
                frequency = frequencies.getOrDefault(value, 0);
                frequencies.put(value, frequency + 1);
                quantities[frequency]--;
                quantities[frequency + 1]++;
                break;
            case 2:
                frequency = frequencies.getOrDefault(value, 0);
                if (frequency == 0)
                    break;
                if (frequency == 1)
                    frequencies.remove(value);
                else
                    frequencies.put(value, frequency - 1);
                // process qt
                quantities[frequency]--;
                quantities[frequency - 1]++;
                break;
            case 3:
                if (value >= quantities.length)
                    result.add(0);
                else
                    result.add(quantities[value] > 0 ? 1 : 0);
                break;
    
            }
    
        }
    
        return result;
    }
}
