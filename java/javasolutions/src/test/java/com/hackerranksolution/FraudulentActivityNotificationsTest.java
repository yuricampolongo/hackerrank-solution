package com.hackerranksolution;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class FraudulentActivityNotificationsTest {
    
    @Test
    public void testFirstInput(){
        List<Integer> activities = new ArrayList<Integer>();
        activities.addAll(List.of(10,20,30,40,50));
        int activityNotifications = FraudulentActivityNotifications.activityNotifications(activities, 3);

        assertEquals(1, activityNotifications);
    }

    @Test
    public void testSecondInput(){
        List<Integer> activities = new ArrayList<Integer>();
        activities.addAll(List.of(1,2,3,4,4));
        int activityNotifications = FraudulentActivityNotifications.activityNotifications(activities, 4);

        assertEquals(0, activityNotifications);
    }

     @Test
    public void testThirdInput(){
        List<Integer> activities = new ArrayList<Integer>();
        activities.addAll(List.of(2, 3, 4, 2, 3, 6, 8, 4, 5));
        int activityNotifications = FraudulentActivityNotifications.activityNotifications(activities, 5);

        assertEquals(2, activityNotifications);
    }
}
