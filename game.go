package main

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/voting"
)

type Game struct {
	players []*Player
}

func (g *Game) run() {
	g.startGame()

	for !g.isOver() {

		g.printPlayers()

		g.startNight()

		i := g.startWerewolfVote()

		if g.players[i].alive && g.players[i].Name() != WEREWOLF {
			fmt.Println("Player ", i, " has died this night. It was a ", g.players[i].Name())
			g.players[i].alive = false
		} else {
			fmt.Println("There are no victims this night.")
		}

		g.printPlayers()

		if g.isOver() {
			break
		}

		i = g.starDailyVote()

		if g.players[i].alive {
			g.players[i].alive = false
			fmt.Println("Player ", i, " has died by public execution. It was a ", g.players[i].Name())
		}
	}

	g.printPlayers()
	fmt.Println("The game is over")

	//night
	//wake werewolves
	//werewolves vote on victim
	//day
	//announce werewolf victim
	//check win conditions
	//vote for player to be eliminated
	//announce death of player
	//check win conditions
	//Repeat
}

func (g *Game) startGame() {
	fmt.Println("The Game has started.")
	event := events.GameStartedEvent{}
	for i := 0; i < len(g.players); i++ {
		if g.players[i].alive {
			g.players[i].OnGameStarted(event)
		}
	}
}

func (g *Game) startNight() {
	fmt.Println("The night starts, Everyone goes to sleep.")
	event := events.NightStartedEvent{}
	for i := 0; i < len(g.players); i++ {
		if g.players[i].alive {
			g.players[i].OnNightStarted(event)
		}
	}
}

func (g *Game) startWerewolfVote() int {

	fmt.Println("The werewolves wake up and choose a victim.")
	ballotBox := &voting.BallotBox{}
	event := events.WerewolfVoteEvent{ballotBox}
	for i := 0; i < len(g.players); i++ {
		if g.players[i].alive {
			g.players[i].OnWerewolfVote(event)
		}
	}

	return ballotBox.Result()
}

func (g *Game) starDailyVote() int {
	fmt.Println("Vote for a player to be executed.")
	ballotBox := &voting.BallotBox{}
	event := events.DailyVoteEvent{ballotBox}
	for i := 0; i < len(g.players); i++ {
		if g.players[i].alive {
			g.players[i].OnDailyVote(event)
		}
	}

	return ballotBox.Result()
}

func (g *Game) isOver() bool {

	stillHumans := false
	stillWerewolves := false

	for i := 0; i < len(g.players); i++ {
		if g.players[i].alive {
			if g.players[i].Name() == WEREWOLF {
				stillWerewolves = true
			} else {
				stillHumans = true
			}
		}
	}

	return !(stillHumans && stillWerewolves)
}

func (g *Game) printPlayers() {
	fmt.Println()
	for i := 0; i < len(g.players); i++ {
		if g.players[i].alive {
			fmt.Println("Player ", i, " alliance:", g.players[i].Name())
		}
	}
}
