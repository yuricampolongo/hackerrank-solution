package com.hackerranksolution;

import java.util.Collections;
import java.util.List;
import java.util.PriorityQueue;

public class FraudulentActivityNotifications {
    public static int activityNotifications(List<Integer> expenditure, int d) {
        int numOfNotif = 0;
        double median;
        PriorityQueue<Integer> max = new PriorityQueue<>();
        PriorityQueue<Integer> min = new PriorityQueue<>(Collections.reverseOrder());
        if (expenditure.get(0) < expenditure.get(1)) {
            min.add(expenditure.get(0));
            max.add(expenditure.get(1));
        } else {
            max.add(expenditure.get(0));
            min.add(expenditure.get(1));
        }

        for (int i = 2; i < expenditure.size() - 1; i++) {
            if (expenditure.get(i) < max.peek())
                min.add(expenditure.get(i));
            else
                max.add(expenditure.get(i));
            balance(min, max);
            if (max.size() + min.size() == d) {
                if (min.size() == max.size())
                    median = ((double) min.peek() + (double) max.peek()) / 2;
                else if (min.size() > max.size())
                    median = min.peek();
                else
                    median = max.peek();
                if (expenditure.get(i + 1) >= 2 * median)
                    numOfNotif++;
                if (max.contains(expenditure.get(i - d + 1)))
                    max.remove(expenditure.get(i - d + 1));
                else
                    min.remove(expenditure.get(i - d + 1));
            }
        }
        return numOfNotif;
    }

    public static void balance(PriorityQueue<Integer> min, PriorityQueue<Integer> max) {
        if (max.size() - min.size() > 1) {
            min.add(max.peek());
            max.remove(max.peek());
        }
        if (min.size() - max.size() > 1) {
            max.add(min.peek());
            min.remove(min.peek());
        }
    }
}
