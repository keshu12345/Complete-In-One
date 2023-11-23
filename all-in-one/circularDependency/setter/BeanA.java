package com.keshu12345.springcore.myspringcore.circularDependency.setter;

import org.springframework.stereotype.Component;

@Component
public class BeanA {
    private BeanB beanB;
    BeanA() {
        System.out.println("BeanA constructor called !!!");
    }

    public void setBean(BeanB beanObj){
        System.out.println("In Constructor ::"+getClass().getSimpleName());
        this.beanB=beanObj;
    }

    public String getBeanA(){
        return "inside "+getClass().getSimpleName();
    }
}
