package voting;

import java.util.ArrayList;
import java.util.List;

public class Client {

public static void main(String[] args) throws IllegalAccessException {

    Candidate candidateA=new Candidate("Candidate A");
    Candidate candidateB=new Candidate("Candidate B");
    Candidate candidateC=new Candidate("Candidate C");
    Candidate candidateD=new Candidate("Candidate D");

    List<Voter> voters=new ArrayList<>();
    voters.add(new Voter(List.of(candidateA,candidateB,candidateC)));
    voters.add(new Voter(List.of(candidateB,candidateC,candidateD)));
    voters.add(new Voter(List.of(candidateA,candidateC,candidateD)));

    VoterCounter voterCounter=new VoterCounter();

    List<Candidate> winners=voterCounter.findWinner(voters);

    for(Candidate winner:winners){
        System.out.println(winner.name+" "+winner.totalPoints);
    }

    }
}
