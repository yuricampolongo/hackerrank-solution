package com.hackerranksolution;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.junit.Test;

public class CountTripletsTest {
    
    @Test
    public void testFirstEntry(){
        assertEquals(6, CountTriplets.countTriplets(Arrays.asList(1l,3l,9l,9l,27l,81l), 3));
    }

    @Test
    public void testSecondEntry(){
        assertEquals(4, CountTriplets.countTriplets(Arrays.asList(1l, 5l, 5l, 25l, 125l), 5));
    }

    @Test
    public void testRatio1Entry(){
        assertEquals(1, CountTriplets.countTriplets(Arrays.asList(1l,1l,1l), 1));
    }

    @Test
    public void testLargeInput(){
        List<Long> input = new ArrayList<>();
        for(int i = 0; i < 100000; i++){
            input.add(1237l);
        }
        assertEquals(166_661_666_700_000l, CountTriplets.countTriplets(input, 1));
    }
}
