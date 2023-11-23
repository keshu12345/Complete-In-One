package com.keshu12345.springcore.myspringcore.controller;

import com.keshu12345.springcore.myspringcore.primary.PrimaryCoach;
import com.keshu12345.springcore.myspringcore.qualifiers.QualifiersCoach;
import com.keshu12345.springcore.myspringcore.setterInjection.SetCoach;
import com.keshu12345.springcore.util.Coach;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class MySpringCoreController {

    // Field Injection this not recommended due to hard unit testing
//    @Autowired
//    private Coach coach;
    private Coach coach;
    private SetCoach setCoach;

    private QualifiersCoach qualifiersCoach;
    private PrimaryCoach primaryCoach;
    // it is optional if only have one constructor
    @Autowired
    public  MySpringCoreController(Coach coach,
                                   @Qualifier("badmintonCoach") QualifiersCoach qualifiersCoach,
                                   PrimaryCoach primaryCoach){
        this.coach=coach;
        this.qualifiersCoach=qualifiersCoach;
        this.primaryCoach=primaryCoach;

    }
    @Autowired
    public void setCoachDoSomeStuff(SetCoach theSetCoach){
        setCoach=theSetCoach;

    }

    @GetMapping("/core")
    public String getWorkout(){
        return coach.dailyWorkout();
    }

    @GetMapping("/setter")
    public String getWorkoutSetter(){
        return setCoach.dailyWorkout();
    }

    @GetMapping("/conQ")
    public String getWorkoutQualifier(){
        return qualifiersCoach.dailyWorkout();
    }


    @GetMapping("/conPrimary")
    public String getWorkoutPrimary(){
        return primaryCoach.dailyWorkout();
    }
}
