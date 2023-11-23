package sorting;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

public class ListOfEmployee {


    public static void main(String []args) {

        List<List<String>> list = new ArrayList<>(List.of( List.of("below", "elbow"), List.of("cat"),List.of("tea", "eat", "ate")));
        Collections.sort(list, new Comparator<List<String>>() {
            @Override
            public int compare(List<String> list1, List<String> list2) {
                return Integer.compare(list2.size(),list1.size());

            }

        });
//        list.sort((list1,list2)->Integer.compare(list1.size(),list2.size()));

        // Sort the list of lists lexicographically
//        list.forEach(Collections::sort);
        list.sort(Comparator.comparing(list1 -> list1.get(0)));

        System.out.println(list.toString());
    }
}
