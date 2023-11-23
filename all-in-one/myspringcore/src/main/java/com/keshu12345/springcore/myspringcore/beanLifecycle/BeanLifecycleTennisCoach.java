package com.keshu12345.springcore.myspringcore.beanLifecycle;

import com.keshu12345.springcore.myspringcore.beanScopes.BeanScopeCoach;
import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;
@Component
@Lazy
public class BeanLifecycleTennisCoach implements BeanLifecycleCoach {

    public BeanLifecycleTennisCoach(){
        System.out.println("In constructor ::"+getClass().getSimpleName());
    }
    @Override
    public String dailyWorkout() {
        return "LazyTennisCoach";
    }
}

