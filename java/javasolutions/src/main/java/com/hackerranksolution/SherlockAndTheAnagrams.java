package com.hackerranksolution;

import java.util.*;

class Result {
    public static int sherlockAndAnagrams(String s) {
        int count = 0;
        Map<String, Integer> map = new HashMap<>();
        for(int i = 0; i < s.length(); i++){
            for(int j = i+1; j <= s.length(); j++){
                String sub = s.substring(i, j);
                String key = reorder(sub);
                if(map.containsKey(key)){
                    count += map.get(key);
                    map.put(key, map.get(key)+1);
                }else{
                    map.put(key, 1);
                }
            }
        }
        return count;
    }

    public static String reorder(String value){
        String[] chars = value.split("");
        Arrays.sort(chars, 0, chars.length);
        return String.join("", chars);
    }

}

public class SherlockAndTheAnagrams {
}
