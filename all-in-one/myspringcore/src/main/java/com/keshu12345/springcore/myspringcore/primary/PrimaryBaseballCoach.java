package com.keshu12345.springcore.myspringcore.primary;

import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Component;

@Component
@Primary
public class PrimaryBaseballCoach implements PrimaryCoach{
    @Override
    public String dailyWorkout() {
        return "PrimaryBaseballCoach workout 45 minutes best practice so far";
    }
}
