package com.hackerranksolution;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class SpecialStringAgainTest {
    
    @Test
    public void testMnonopoo(){
        String test = "mnonopoo";
        long substrCount = SpecialStringAgain.substrCount(test.length(), test);
        assertEquals(12, substrCount);
    }

    @Test
    public void testOoo(){
        String test = "ooo";
        long substrCount = SpecialStringAgain.substrCount(test.length(), test);
        assertEquals(6, substrCount);
    }

     @Test
    public void testAbcbaba(){
        String test = "abcbaba";
        long substrCount = SpecialStringAgain.substrCount(test.length(), test);
        assertEquals(10, substrCount);
    }
}
