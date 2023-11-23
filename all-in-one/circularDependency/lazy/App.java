package com.keshu12345.springcore.myspringcore.circularDependency.lazy;

import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class App {
    public static void main(String[] args) {
        ApplicationContext applicationContext=new AnnotationConfigApplicationContext(AppConfig.class);
        BeanA beanA= (BeanA) applicationContext.getBean("beanA");
        String result=beanA.getBeanA();
        System.out.println(result);
    }
}
