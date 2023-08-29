package com.hackerranksolution;

import org.junit.Test;

import static org.junit.Assert.*;

public class SherlockAndTheAnagramsTest {

    @Test public void testKkkk() {
        assertEquals(10, Result.sherlockAndAnagrams("kkkk"));
    }

    @Test public void testMom() {
        assertEquals(2, Result.sherlockAndAnagrams("mom"));
    }

    @Test public void testAbba() {
        assertEquals(4, Result.sherlockAndAnagrams("abba"));
    }

    @Test public void testAbcd() {
        assertEquals(0, Result.sherlockAndAnagrams("abcd"));
    }

    @Test public void testIfailuhkqq() {
        assertEquals(3, Result.sherlockAndAnagrams("ifailuhkqq"));
    }

    @Test public void testCdcd() {
        assertEquals(5, Result.sherlockAndAnagrams("cdcd"));
    }
}