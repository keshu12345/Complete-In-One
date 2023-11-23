package com.keshu12345.springcore.myspringcore.setterInjection;

import org.springframework.stereotype.Component;

@Component
public class SetCricketCoach implements  SetCoach{
    @Override
    public String dailyWorkout() {
        return "Setting bowling for 15 minutes";
    }
}
