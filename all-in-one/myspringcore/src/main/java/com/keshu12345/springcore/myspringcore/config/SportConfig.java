package com.keshu12345.springcore.myspringcore.config;

import com.keshu12345.springcore.myspringcore.configBean.ConfigCoach;
import com.keshu12345.springcore.myspringcore.configBean.SwimCoach;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

@Configuration
//@ComponentScan
public class SportConfig {
    @Bean
    public ConfigCoach swimCoach(){
        return new SwimCoach();
    }
}
