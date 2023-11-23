package com.keshu12345.springcore.myspringcore.circularDependency.circular;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class BeanB {
    private BeanA beanA;
    @Autowired
    public BeanB(BeanA beanA) {
        System.out.println("In Constructor ::"+getClass().getSimpleName());
        this.beanA = beanA;
    }
}
