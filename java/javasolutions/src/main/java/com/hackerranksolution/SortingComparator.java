package com.hackerranksolution;

import java.util.Comparator;

public class SortingComparator {
    class Checker implements Comparator<Player> {
        public int compare(Player a, Player b) {
            if (a.score < b.score){
                return 1;
            } else if(a.score > b.score) {
                return -1;
            }
            return a.name.compareTo(b.name);
        }
    }

    class Player {
        int score;
        String name;
        public Player(String name, int score) {
            this.score = score;
            this.name = name;
        }
    }
}
