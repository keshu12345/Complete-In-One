package com.keshu12345.springcore.myspringcore.controller;

import com.keshu12345.springcore.myspringcore.lazyInitialization.LazyCoach;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class LazyController {
    private LazyCoach lazyCoach;

    public LazyController(@Qualifier("lazyCricketCoach") LazyCoach lazyCoach) {
        System.out.println("In Constructor ::"+getClass().getSimpleName());
        this.lazyCoach = lazyCoach;
    }
    @GetMapping("/lazy")
    public String getDailyWorkout(){
        return lazyCoach.dailyWorkout();
    }
}
