package com.keshu12345.springcore.myspringcore.primary;

import org.springframework.stereotype.Component;

@Component
public class PrimaryBadmintonCoach implements  PrimaryCoach{
    @Override
    public String dailyWorkout() {
        return "PrimaryBadmintonCoach workout for 30 minutes";
    }
}
