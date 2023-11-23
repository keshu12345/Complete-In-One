package voting;

import java.util.*;

public class VoterCounter {
   public List<Candidate> findWinner(List<Voter>voters){
       Map<String, Candidate>candidateMap=new HashMap<>();

       for(Voter voter:voters){

           for(int i=0;i<voter.votedCandidates.size();i++){
               Candidate candidate= voter.votedCandidates.get(i);
               int points=3-i;
               candidate.addPoint(points);
               candidateMap.put(candidate.name,candidate);

           }
       }

       List<Candidate>candidateList=new ArrayList<>(candidateMap.values());
       Collections.sort(candidateList,Comparator.comparingInt(c->c.totalPoints));


       return candidateList;
       }

}
