package com.keshu12345.springcore.myspringcore.controller;

import com.keshu12345.springcore.myspringcore.beanScopes.BeanScopeCoach;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class BeanSCopeController {

    private BeanScopeCoach beanScopeCoach1;
    private BeanScopeCoach beanScopeCoach2;

    public BeanSCopeController(@Qualifier("beanScopeCricketCoach") BeanScopeCoach beanScopeCoach1,
                               @Qualifier("beanScopeCricketCoach") BeanScopeCoach beanScopeCoach2) {
        this.beanScopeCoach1 = beanScopeCoach1;
        this.beanScopeCoach2=beanScopeCoach2;
    }


    @GetMapping("/bean")
    public String getDailyWorkout() {
        return beanScopeCoach1.dailyWorkout();
    }

    @GetMapping("/check")
    public String check() {
        return "Comparing beans: myCoach == anotherCoach, " + (beanScopeCoach1 == beanScopeCoach2);
    }
}
