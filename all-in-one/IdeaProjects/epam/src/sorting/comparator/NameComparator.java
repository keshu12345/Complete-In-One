package sorting.comparator;

import java.util.Comparator;


public class NameComparator implements Comparator<Employee> {
    @Override
    public int compare(Employee e1,Employee e2){
        return e2.getName().compareTo(e1.getName());
    }
}
