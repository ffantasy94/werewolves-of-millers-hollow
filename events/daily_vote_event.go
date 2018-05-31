package events

import (
	"github.com/hunghoangmagrabbit/werewolves-of-millers-hollow/voting"
)

type DailyVoteEvent struct {
	*voting.BallotBox
}

func NewDailyVoteEvent() *DailyVoteEvent {
	return &DailyVoteEvent{voting.NewBallotBox()}
}
