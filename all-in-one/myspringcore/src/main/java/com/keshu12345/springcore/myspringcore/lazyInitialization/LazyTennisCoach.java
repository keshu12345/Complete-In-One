package com.keshu12345.springcore.myspringcore.lazyInitialization;

import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;

@Component
@Lazy
public class LazyTennisCoach implements LazyCoach{

    public LazyTennisCoach(){
        System.out.println("In constructor ::"+getClass().getSimpleName());
    }
    @Override
    public String dailyWorkout() {
        return "LazyTennisCoach";
    }
}
