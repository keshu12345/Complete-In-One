package voting;

public class Candidate {

    String name;
    int totalPoints=0;

    public Candidate(String name){
        this.name=name;
        this.totalPoints=0;
    }

    public void addPoint(int points){
        this.totalPoints+=points;
    }
}
