package com.keshu12345.springcore.myspringcore.beanScopes;

import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;
@Component
@Lazy
public class BeanScopeTennisCoach implements BeanScopeCoach {

    public BeanScopeTennisCoach(){
        System.out.println("In constructor ::"+getClass().getSimpleName());
    }
    @Override
    public String dailyWorkout() {
        return "LazyTennisCoach";
    }
}

