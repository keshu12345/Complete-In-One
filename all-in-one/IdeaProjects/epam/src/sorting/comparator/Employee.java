package sorting.comparator;

import java.security.PrivateKey;

public class Employee {
    private String name;
    private int empId;
    private int age;

    Employee(String name ,int empId,int age) {
        this.age = age;
        this.name = name;
        this.empId = empId;
    }

    public String getName(){
        return name;
    }
    public void setName(String name){
        this.name=name;
    }

    public int getAge(){
        return age;
    }
    public  void setAge(int age){
        this.age=age;
    }
    public int getEmpId(){
        return empId;
    }

    public void setEmpId(int empId){
        this.empId=empId;
    }


}
