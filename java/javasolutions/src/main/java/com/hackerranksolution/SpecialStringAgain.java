package com.hackerranksolution;

public class SpecialStringAgain {

    static long substrCount(int n, String s) {
        long counter = n;
        int i = 0, j = 0;

        while(j < n && i < n){
            String sub = s.substring(j, i+1);
            if(countUniqueCharacters(sub) > 2){
                i = ++j;
                continue;
            }
            String reversed = new StringBuilder(sub).reverse().toString();
            if(sub.length() > 1 && sub.equals(reversed)){
                counter++;
            }
            i++;
            if(i == n) {
                i = ++j;
            }
        }

        return counter;
    }

    public static int countUniqueCharacters(String input) {
        String unique = input.replaceAll("(.)(?=.*?\\1)", "");
        return unique.length();
    }
}
