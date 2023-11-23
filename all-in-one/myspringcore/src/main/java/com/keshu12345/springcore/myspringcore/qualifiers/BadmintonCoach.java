package com.keshu12345.springcore.myspringcore.qualifiers;

import org.springframework.stereotype.Component;

@Component
public class BadmintonCoach implements QualifiersCoach{
    @Override
    public String dailyWorkout() {
        return "Badminton Coach give time 30 minutes to workout";
    }
}
