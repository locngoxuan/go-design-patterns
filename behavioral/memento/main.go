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

func (s *Memento) FindMove(i int) *Memento {
	if s.No == i {
		return s
	}
	return s.Next.FindMove(i)
}

type Move struct {
	ChessPiece string
	Col        string
	Row        string
}

type ScoreSheet struct {
	CurrentMove int
	memento     *Memento
}

func (s *ScoreSheet) AddMove(playerId int, m Move) {
	lm := s.memento.FindMove(s.CurrentMove)
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
		builder.WriteString(fmt.Sprintf("player: %v | %v%v%v",
			m.Next.Player, m.Next.ChessPiece, m.Next.Col, m.Next.Row))
		builder.WriteString("\n")
		m = m.Next
	}
	builder.WriteString("------- END -------")
	return builder.String()
}

func (s *ScoreSheet) Undo() {
	s.CurrentMove--
}

func (s *ScoreSheet) Redo() {
	s.CurrentMove++
}

const (
	Player int = iota
	Bot
)

func NewMove(chessPiece, col, row string) Move {
	return Move{
		ChessPiece: chessPiece,
		Col:        col,
		Row:        row,
	}
}

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
	sheet.AddMove(Player, NewMove("", "e", "4"))
	sheet.AddMove(Bot, NewMove("", "e", "5"))
	sheet.AddMove(Player, NewMove("", "d", "3"))
	sheet.AddMove(Bot, NewMove("Q", "h", "4"))
	log.Printf("latest snapshot: %v", sheet.Snapshot())
	sheet.Undo()
	sheet.Undo()
	log.Printf("latest snapshot: %v", sheet.Snapshot())
	sheet.AddMove(Player, NewMove("", "g", "3"))
	sheet.AddMove(Bot, NewMove("Q", "g", "5"))
	log.Printf("latest snapshot: %v", sheet.Snapshot())
	sheet.Undo()
	sheet.Undo()
	sheet.Redo()
	log.Printf("latest snapshot: %v", sheet.Snapshot())
}
