package com.keshu12345.springcore.myspringcore.qualifiers;

public class TennisCoach implements QualifiersCoach{
    @Override
    public String dailyWorkout() {
        return "Tennis Coach give time 15 minutes for work out ):";
    }
}
