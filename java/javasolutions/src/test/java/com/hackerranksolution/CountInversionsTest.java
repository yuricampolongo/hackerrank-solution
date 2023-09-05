package com.hackerranksolution;

import static org.junit.Assert.assertEquals;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import org.junit.Test;

public class CountInversionsTest {
    
    @Test
    public void unorderedArray(){
        List<Integer> asList = Arrays.asList(2,1,3,1,2);
        long countInversions = CountInversions.countInversions(asList);
        assertEquals(4, countInversions);
    }

    @Test
    public void orderedArray(){
        List<Integer> asList = Arrays.asList(1,1,1,2,3);
        long countInversions = CountInversions.countInversions(asList);
        assertEquals(0, countInversions);
    }

     @Test
    public void orderedArray6Iterations(){
        List<Integer> asList = Arrays.asList(7,5,3,1);
        long countInversions = CountInversions.countInversions(asList);
        assertEquals(6, countInversions);
    }

    @Test
    public void bigOrderedArray(){
        List<Integer> input = IntStream.rangeClosed(1, 65911).boxed().collect(Collectors.toList());
        long countInversions = CountInversions.countInversions(input);
        assertEquals(0, countInversions);
    }

    @Test
    public void bigInversedArray(){
        List<Integer> input = IntStream.rangeClosed(1, 100000).boxed().collect(Collectors.toList());
        Collections.reverse(input);
        long countInversions = CountInversions.countInversions(input);
        assertEquals(4_999_950_000l, countInversions);
    }

}
