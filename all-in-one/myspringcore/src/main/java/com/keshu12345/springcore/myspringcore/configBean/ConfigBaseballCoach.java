package com.keshu12345.springcore.myspringcore.configBean;

import org.springframework.stereotype.Component;

@Component
public class ConfigBaseballCoach implements ConfigCoach{
    @Override
    public String dailyWorkout() {
        return "Inside ConfigBaseballCoach  working 3 hours per day";
    }
}
