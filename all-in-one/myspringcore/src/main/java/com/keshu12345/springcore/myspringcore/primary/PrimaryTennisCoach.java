package com.keshu12345.springcore.myspringcore.primary;

import org.springframework.stereotype.Component;

@Component
public class PrimaryTennisCoach implements PrimaryCoach{
    @Override
    public String dailyWorkout() {
        return "PrimaryTennisCoach workout for 20 minutes";
    }
}
