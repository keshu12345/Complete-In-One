package com.keshu12345.springcore.myspringcore.beanScopes;

import com.keshu12345.springcore.myspringcore.lazyInitialization.LazyCoach;
import org.springframework.stereotype.Component;

@Component
public class BeanScopeBaseballCoach implements BeanScopeCoach {
    public BeanScopeBaseballCoach(){
        System.out.println("In constructor ::"+getClass().getSimpleName());
    }
    @Override
    public String dailyWorkout() {
        return "LazyBaseballCoach";
    }
}

