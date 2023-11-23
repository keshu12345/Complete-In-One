package com.keshu12345.springcore.myspringcore.circularDependency.setter;

import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class App {
    public static void main(String[] args) {
        ApplicationContext applicationContext=new AnnotationConfigApplicationContext(AppConfig.class);

        BeanA beanA=(BeanA)applicationContext.getBean("beanA");
        System.out.println( beanA.getBeanA());
        BeanB beanB=(BeanB) applicationContext.getBean("beanB");
        System.out.println(beanB.getBeanB());
    }
}
