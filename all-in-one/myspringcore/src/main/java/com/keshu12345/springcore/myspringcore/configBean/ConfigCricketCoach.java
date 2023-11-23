package com.keshu12345.springcore.myspringcore.configBean;

import org.springframework.stereotype.Component;

@Component
public class ConfigCricketCoach implements ConfigCoach{
    @Override
    public String dailyWorkout() {
        return "Inside ConfigCricketCoach  working 1 hours per day";
    }
}
