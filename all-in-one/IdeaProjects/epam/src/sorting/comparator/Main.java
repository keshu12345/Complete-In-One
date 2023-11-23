package sorting.comparator;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

public class Main {

    public static void main(String[] args) {
        Employee e1 = new Employee("keshav", 101, 30);
        Employee e2 = new Employee("Ashish", 102, 35);
        Employee e3 = new Employee("Akash", 103, 20);
        Employee e4 = new Employee("Himesh", 104, 40);
        Employee e5 = new Employee("Dhairya", 105, 19);
        Employee e6 = new Employee("Dhanush", 106, 29);
        List<List<Employee>> list = new ArrayList<>(List.of(List.of(e1, e2), List.of(e3, e4, e5), List.of(e6)));

//        System.out.println(list.toString());
        list.sort((list1, list2) -> Integer.compare(list1.size(), list2.size()));
        list.sort(Comparator.comparingInt(List::size));
        System.out.println(list.toString());
        list.sort((list1, list2) -> Integer.compare(list2.size(), list1.size()));
        System.out.println(list.toString());
//        list.forEach(list1 -> list1.sort(Comparator.comparing(Employee::getName)));
//        for (List<Employee> l : list) {
//            for (Employee employee : l) {
//                System.out.println(employee);
//            }

            // Sort the inner lists based on employee name in descending order
            list.forEach(innerList -> innerList.sort((eo1, eo2) -> eo2.getName().compareTo(eo1.getName())));

            // Print the sorted list
//            for (List<Employee> innerList : list) {
//                for (Employee employee : innerList) {
//                    System.out.println(employee);
//                }
//            }
        }
    }

