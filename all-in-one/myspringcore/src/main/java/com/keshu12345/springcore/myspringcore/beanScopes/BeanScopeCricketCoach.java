package com.keshu12345.springcore.myspringcore.beanScopes;

import com.keshu12345.springcore.myspringcore.lazyInitialization.LazyCoach;
import org.springframework.beans.factory.config.ConfigurableBeanFactory;
import org.springframework.context.annotation.Lazy;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

@Component
@Lazy
@Scope(ConfigurableBeanFactory.SCOPE_PROTOTYPE)
public class BeanScopeCricketCoach implements BeanScopeCoach {
    public BeanScopeCricketCoach() {
        System.out.println("in Constructor ::"+getClass().getSimpleName());
    }

    @Override
    public String dailyWorkout() {
        return "LazyCricketCoach";
    }
}
