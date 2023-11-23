package com.keshu12345.springcore.myspringcore.lazyInitialization;

import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;

@Component
@Lazy
public class LazyCricketCoach implements LazyCoach{
    public LazyCricketCoach() {
        System.out.println("in Constructor ::"+getClass().getSimpleName());
    }

    @Override
    public String dailyWorkout() {
        return "LazyCricketCoach";
    }
}
