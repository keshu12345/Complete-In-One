package com.keshu12345.springcore.myspringcore.constructorInjection;

import com.keshu12345.springcore.util.Coach;
import org.springframework.stereotype.Component;

@Component
public class CricketCoach implements Coach {
    @Override
    public String dailyWorkout() {
        return "Daily Cricket Exercise 15 minutes";
    }
}
