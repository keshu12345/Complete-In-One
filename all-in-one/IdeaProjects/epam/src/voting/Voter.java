package voting;

import java.util.List;

public class Voter {

    List<Candidate>votedCandidates;

    public Voter(List<Candidate>votedCandidates) throws IllegalAccessException {
        if(votedCandidates.size()>3){
            throw new IllegalAccessException("A voter can vote for up to three candidates");
        }
        this.votedCandidates=votedCandidates;
    }

}
