package com.keshu12345.springcore.myspringcore.qualifiers;

import org.springframework.stereotype.Component;

@Component
public class BaseballCoach implements  QualifiersCoach{
    @Override
    public String dailyWorkout() {
        return " BaseBall Coach give time to 20 minutes for workout";
    }
}
