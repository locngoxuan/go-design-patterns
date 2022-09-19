package main

import (
	"fmt"
	"log"
	"strings"
)

type Memento struct {
	No     int
	Next   *Memento
	Player int
	Move
}

func (s *Memento) FindLastMove() *Memento {
	if s.Next == nil {
		return s
	}
	return s.Next.FindLastMove()
}

func (s *Memento) FindAnyMove(i int) *Memento {
	if s.No == i {
		return s
	}
	return s.Next.FindAnyMove(i)
}

type Move struct {
	XSrc int
	YSrc int
	XDst int
	YDst int
}

type ScoreSheet struct {
	CurrentMove int
	memento     *Memento
}

func (s *ScoreSheet) AddMove(playerId int, m Move) {
	lm := s.memento.FindAnyMove(s.CurrentMove)
	lm.Next = &Memento{
		No:     lm.No + 1,
		Next:   nil,
		Player: playerId,
		Move:   m,
	}
	s.CurrentMove++
}

func (s ScoreSheet) Snapshot() string {
	var builder strings.Builder
	m := s.memento
	builder.WriteString("\n------- START -------")
	builder.WriteString("\n")
	for m.Next != nil && m.Next.No <= s.CurrentMove {
		builder.WriteString(fmt.Sprintf("player: %v | from: x = %v, y = %v | to : x = %v, y = %v",
			m.Next.Player, m.Next.XSrc, m.Next.YSrc, m.Next.XDst, m.Next.YDst))
		builder.WriteString("\n")
		m = m.Next
	}
	builder.WriteString("------- END -------")
	return builder.String()
}

func (s *ScoreSheet) Undo() {
	s.CurrentMove--
}

const (
	Player int = 0
	Bot
)

func main() {
	sheet := &ScoreSheet{
		CurrentMove: 0,
		memento: &Memento{
			No:     0,
			Next:   nil,
			Player: -1,
			Move:   Move{},
		},
	}
	sheet.AddMove(Player, Move{
		XSrc: 3,
		YSrc: 1,
		XDst: 3,
		YDst: 2,
	})
	sheet.AddMove(Bot, Move{
		XSrc: 3,
		YSrc: 7,
		XDst: 3,
		YDst: 6,
	})
	sheet.AddMove(Player, Move{
		XSrc: 4,
		YSrc: 1,
		XDst: 4,
		YDst: 3,
	})
	sheet.AddMove(Bot, Move{
		XSrc: 4,
		YSrc: 7,
		XDst: 4,
		YDst: 6,
	})
	log.Printf("snapshot after 4 moves: %v", sheet.Snapshot())
	sheet.Undo()
	sheet.Undo()
	log.Printf("snapshot after 2 moves: %v", sheet.Snapshot())
	sheet.AddMove(Player, Move{
		XSrc: 4,
		YSrc: 1,
		XDst: 4,
		YDst: 2,
	})
	sheet.AddMove(Bot, Move{
		XSrc: 4,
		YSrc: 7,
		XDst: 4,
		YDst: 5,
	})
	log.Printf("snapshot after 4 moves (2 new moves): %v", sheet.Snapshot())
}
