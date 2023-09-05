package com.hackerranksolution;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicLong;

public class CountInversions {

    public static long countInversions(List<Integer> arr) {
        AtomicLong totalSwaps = new AtomicLong(0);
        mergeSort(arr, totalSwaps);
        return totalSwaps.get();
    }

    public static void mergeSort(List<Integer> input, AtomicLong totalSwaps) {
        int inputSize = input.size();
        if (inputSize < 2) {
            return;
        }

        int middlePos = inputSize / 2;
        List<Integer> leftEls = new ArrayList<>();
        List<Integer> rightEls = new ArrayList<>();

        for (int i = 0; i < middlePos; i++) {
            leftEls.add(input.get(i));
        }
        for (int i = middlePos; i < inputSize; i++) {
            rightEls.add(i - middlePos, input.get(i));
        }
        mergeSort(leftEls, totalSwaps);
        mergeSort(rightEls, totalSwaps);

        merge(input, leftEls, rightEls, totalSwaps);
    }

    public static void merge(List<Integer> input, List<Integer> leftEls, List<Integer> rightEls, AtomicLong totalSwaps) {
        int i = 0, j = 0, k = 0;

        while (i < leftEls.size() && j < rightEls.size()) {
            if (leftEls.get(i) <= rightEls.get(j)) {
                input.set(k++, leftEls.get(i++));
            } else {
                input.set(k++, rightEls.get(j++));
                totalSwaps.getAndAdd(leftEls.size() - i);
            }
        }
        while (i < leftEls.size()) {
            input.set(k++, leftEls.get(i++));
        }
        while (j < rightEls.size()) {
            input.set(k++, rightEls.get(j++));
        }
    }

}
