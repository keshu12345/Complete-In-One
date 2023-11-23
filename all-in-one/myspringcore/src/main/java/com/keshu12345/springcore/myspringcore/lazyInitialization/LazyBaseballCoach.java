package com.keshu12345.springcore.myspringcore.lazyInitialization;

import org.springframework.stereotype.Component;

@Component
public class LazyBaseballCoach implements  LazyCoach{
    public LazyBaseballCoach(){
        System.out.println("In constructor ::"+getClass().getSimpleName());
    }
    @Override
    public String dailyWorkout() {
        return "LazyBaseballCoach";
    }
}
