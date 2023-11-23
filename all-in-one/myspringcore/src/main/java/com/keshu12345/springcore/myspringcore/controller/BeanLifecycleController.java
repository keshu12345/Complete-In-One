package com.keshu12345.springcore.myspringcore.controller;

import com.keshu12345.springcore.myspringcore.beanLifecycle.BeanLifecycleCoach;
import com.keshu12345.springcore.myspringcore.lazyInitialization.LazyCoach;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class BeanLifecycleController {

    private BeanLifecycleCoach beanLifecycleCoach;

    public BeanLifecycleController(@Qualifier("beanLifecycleCricketCoach") BeanLifecycleCoach beanLifecycleCoach) {
        this.beanLifecycleCoach = beanLifecycleCoach;
    }

    @GetMapping("/beanCycle")
    public String getDailyWorkout(){
        return beanLifecycleCoach.dailyWorkout();
    }
}
