package com.keshu12345.springcore.myspringcore.circularDependency.lazy;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;

@Component
public class BeanB {
    private BeanA beanA;
    @Autowired
    public BeanB(@Lazy  BeanA beanA) {
        System.out.println("In Constructor ::"+getClass().getSimpleName());
        this.beanA = beanA;
    }
}
