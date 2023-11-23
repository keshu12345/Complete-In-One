package com.keshu12345.springcore.myspringcore.circularDependency.setter;

import org.springframework.stereotype.Component;

@Component
public class BeanB {
    private BeanA beanA;
    BeanB() {
        System.out.println("BeanB constructor called !!!");
    }

    public void setBeanB(BeanA beanA){
        this.beanA=beanA;
        System.out.println("In Constructor ::"+getClass().getSimpleName());
    }

    public String getBeanB(){
        return "inside "+getClass().getSimpleName();
    }
}
