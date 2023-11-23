import java.util.*;

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello world!");

        //Input: strs = ["eat","tea","elbow","cat","below","ate"]
        //Output: [["cat"],["below","elbow"],["ate","eat","tea"]]

        String []words={"eat","tea","elbow","cat","below","ate"};

        List<List<String>>result=solve(words);
        System.out.println(result);

    }

    private static List<List<String>> solve(String[] words) {
        HashMap<String,List<String>> hashMap=new HashMap<>();

        List<List<String>>result=new ArrayList<>();

        for(String word:words) {

            char[]charArray=word.toCharArray();
            Arrays.sort(charArray);

            String sortedWord=String.valueOf(charArray);

            if(!hashMap.containsKey(sortedWord)) {
                hashMap.put(sortedWord,new ArrayList<>());
            }
            hashMap.get(sortedWord).add(word);

        }
        result.addAll(hashMap.values());


Collections.sort(result, new Comparator<List<String>>() {
    @Override
    public int compare(List<String> o1, List<String> o2) {
        return o1.get(0).compareTo(o2.get(0));
    }
});

        return result;
    }


}
