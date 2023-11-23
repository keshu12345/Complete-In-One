package com.keshu12345.springcore.myspringcore.beanLifecycle;

import com.keshu12345.springcore.myspringcore.beanScopes.BeanScopeCoach;
import org.springframework.stereotype.Component;

@Component
public class BeanLifecycleBaseballCoach implements BeanLifecycleCoach {
    public BeanLifecycleBaseballCoach(){
        System.out.println("In constructor ::"+getClass().getSimpleName());
    }
    @Override
    public String dailyWorkout() {
        return "LazyBaseballCoach";
    }
}

