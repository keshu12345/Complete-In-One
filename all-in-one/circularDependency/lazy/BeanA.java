package com.keshu12345.springcore.myspringcore.circularDependency.lazy;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class BeanA {
    private BeanB beanB;
    @Autowired
    public BeanA(BeanB beanB) {
        System.out.println("In Constructor ::"+getClass().getSimpleName());
        this.beanB = beanB;
    }
    public  String getBeanA(){
        return "inside "+getClass().getSimpleName();
    }
}
