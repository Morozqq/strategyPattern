package main

import "fmt"

type GuitarStrategy interface {
	Play()
}

//Structure for play styles

type PlayFingerstyle struct {
}
type PlayChords struct {
}
type PlayBoth struct {
}

func (pFinger PlayFingerstyle) Play() {
	fmt.Println("He/She is playing FingerStyle")
}

func (pChords PlayChords) Play() {
	fmt.Println("He/She is playing Chords")
}

func (pBoth PlayBoth) Play() {
	fmt.Println("He/She is playing both styles")
}

type GuitarPlayer struct {
	playStyle GuitarStrategy
}

func NewGuitarPlayer(playStyle GuitarStrategy) GuitarPlayer {
	return GuitarPlayer{playStyle}
}
func (gp GuitarPlayer) PlayGuitar() {
	gp.playStyle.Play()
}
func main() {
	fingerstylePlayer := NewGuitarPlayer(PlayFingerstyle{})
	fingerstylePlayer.PlayGuitar()

	chordsPlayer := NewGuitarPlayer(PlayChords{})
	chordsPlayer.PlayGuitar()

	bothPlayer := NewGuitarPlayer(PlayBoth{})
	bothPlayer.PlayGuitar()
}
