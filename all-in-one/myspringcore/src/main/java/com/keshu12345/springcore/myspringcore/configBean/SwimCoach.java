package com.keshu12345.springcore.myspringcore.configBean;

public class SwimCoach implements ConfigCoach{
    public SwimCoach() {
        System.out.println("in constructor ::"+getClass().getSimpleName());
    }

    @Override
    public String dailyWorkout() {
        return " SwimCoach  working 5 hours per day";
    }
}
