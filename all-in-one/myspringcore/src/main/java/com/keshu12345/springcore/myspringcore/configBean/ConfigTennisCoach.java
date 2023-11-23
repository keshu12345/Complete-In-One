package com.keshu12345.springcore.myspringcore.configBean;

import org.springframework.stereotype.Component;

@Component
public class ConfigTennisCoach implements  ConfigCoach{
    @Override
    public String dailyWorkout() {
        return " ConfigTennisCoach  working 2 hours per day";
    }
}
