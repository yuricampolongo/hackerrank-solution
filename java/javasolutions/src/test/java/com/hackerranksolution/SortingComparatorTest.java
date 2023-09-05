package com.hackerranksolution;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import org.junit.Test;

import com.hackerranksolution.SortingComparator.Player;

public class SortingComparatorTest {
    
    private SortingComparator sortingComparator = new SortingComparator();

    @Test
    public void testCompare(){
        List<Player> players = new ArrayList<>();
        players.add(sortingComparator.new Player("Smith",20));
        players.add(sortingComparator.new Player("Jones",15));
        players.add(sortingComparator.new Player("Jones",20));

        Collections.sort(players, sortingComparator.new Checker());

        assertEquals("Jones", players.get(0).name);
        assertEquals(20, players.get(0).score);
        assertEquals("Smith", players.get(1).name);
        assertEquals(20, players.get(1).score);
        assertEquals("Jones", players.get(2).name);
        assertEquals(15, players.get(2).score);
    }

}


