package com.keshu12345.springcore.myspringcore.controller;

import com.keshu12345.springcore.myspringcore.beanScopes.BeanScopeCoach;
import com.keshu12345.springcore.myspringcore.configBean.ConfigCoach;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ConfigBeanController {

    private ConfigCoach coach;

    @Autowired
    public ConfigBeanController( @Qualifier("swimCoach") ConfigCoach coach) {
        this.coach = coach;
    }
    @GetMapping("/config-bean")
    public String getWorkout(){
        return coach.dailyWorkout();
    }
}
