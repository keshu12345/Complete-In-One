package sorting.comparator;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

public class Client {

 public static void main(String[] args) {

     Employee e1 = new Employee("keshav", 101, 30);
     Employee e2 = new Employee("Ashish", 102, 35);
     Employee e3 = new Employee("Akash", 103, 20);
     Employee e4 = new Employee("Himesh", 104, 40);
     Employee e5 = new Employee("Dhairya", 105, 19);
     Employee e6 = new Employee("Dhanush", 106, 29);
     List<Employee> list = new ArrayList<>(List.of(e1, e2, e3, e4, e5, e6));
     Collections.sort(list,new NameComparator());

     for(Employee employee:list) {
         System.out.println(employee.getName()+" "+employee.getAge()+" "+employee.getEmpId());
     }
     System.out.println("----------------------------------");
     Collections.sort(list,new AgeComparator());
     for(Employee employee:list) {
         System.out.println(employee.getName()+" "+employee.getAge()+" "+employee.getEmpId());
     }

     Collections.sort(list, new Comparator<Employee>() {
         @Override
         public int compare(Employee o1, Employee o2) {
             return Integer.compare(o2.getEmpId(),o1.getEmpId());
         }
     });
     System.out.println("-------------Sort asceding order--------------------");
     for(Employee employee:list) {
         System.out.println(employee.getName()+" "+employee.getAge()+" "+employee.getEmpId());
     }


//     list.sort((emp1, emp2) -> Integer.compare(emp2.getEmpId(), emp1.getEmpId()));
     Collections.sort(list,(o1,o2)->Integer.compare(o2.getEmpId(),o1.getEmpId()));
     System.out.println("----------------------------------");
     for(Employee employee:list) {
         System.out.println(employee.getName()+" "+employee.getAge()+" "+employee.getEmpId());
     }
     list.sort(Comparator.comparingInt(Employee::getEmpId));
     System.out.println("----------------------------------");
     for(Employee employee:list) {
         System.out.println(employee.getName()+" "+employee.getAge()+" "+employee.getEmpId());
     }


 }

}
