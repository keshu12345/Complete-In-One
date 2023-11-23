package com.keshu12345.springcore.myspringcore.beanLifecycle;

import com.keshu12345.springcore.myspringcore.beanScopes.BeanScopeCoach;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import org.springframework.beans.factory.config.ConfigurableBeanFactory;
import org.springframework.context.annotation.Lazy;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

@Component
public class BeanLifecycleCricketCoach implements BeanLifecycleCoach {
    public BeanLifecycleCricketCoach() {
        System.out.println("in Constructor ::"+getClass().getSimpleName());
    }

    @Override
    public String dailyWorkout() {
        return "LazyCricketCoach";
    }

    // define your init method
    @PostConstruct
    public void doMyStartupStuff(){
        System.out.println("in doMyStartupStuff()"+getClass().getSimpleName());
    }

    // define your destroy method
    @PreDestroy
    public void doMyCleanUpStuff(){
        System.out.println("In doMyCleanUpStuff()"+getClass().getSimpleName());
    }




}
