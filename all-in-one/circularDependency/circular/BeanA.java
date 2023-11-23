package com.keshu12345.springcore.myspringcore.circularDependency.circular;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;

@Component
public class BeanA {
    private BeanB beanB;
    @Autowired
    public BeanA(BeanB beanB) {
        System.out.println("In Constructor ::"+getClass().getSimpleName());
        this.beanB = beanB;
    }
}
