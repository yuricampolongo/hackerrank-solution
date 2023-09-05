package com.hackerranksolution;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class FraudulentActivityNotifications {
    public static int activityNotifications(List<Integer> expenditure, int d) {
        int totalNotifications = 0;

        for(int i = d; i< expenditure.size(); i++) {
            List<Integer> subList = expenditure.subList(i-d, i);
            Integer currentExpenditure = expenditure.get(i);
            double median = getMedian(subList);
            if(currentExpenditure >= median * 2){
                totalNotifications++;
            }
        }

        return totalNotifications;
    }

    // Quick Select Algorithm
    private static double getMedian(List<Integer> expenditures){
        if(expenditures.size() % 2 != 0){ // get the min middle element
            return quickSelect(expenditures, 0, expenditures.size()-1,expenditures.size()/2);
        } else {
            int median1 = quickSelect(new ArrayList<>(expenditures), 0, expenditures.size()-1,expenditures.size()/2);
            int median2 = quickSelect(new ArrayList<>(expenditures), 0, expenditures.size()-1,expenditures.size()/2 - 1);
            return (median1 + median2) / 2.0;
        }
    }


    private static int quickSelect(List<Integer> expenditures, int left, int right, int k){
        if (left == right){
            return expenditures.get(left);
        }

        int pivotIndex = partition(expenditures, left, right);
        if(k == pivotIndex){
            return expenditures.get(k);
        } else if (k < pivotIndex){
            return quickSelect(expenditures, left, pivotIndex - 1, k);
        } else {
            return quickSelect(expenditures, pivotIndex + 1, right, k);
        }
    }

    private static int partition(List<Integer> expenditures, int left, int right){
        int pivotValue = expenditures.get(left + right / 2);

        int i = left;
        for(int j = left; j < right; j++){
            if(expenditures.get(j) <= pivotValue){
                Collections.swap(expenditures, i, j);
                i++;
            }
        }

        Collections.swap(expenditures, i, right);
        return i;
    }
}

