package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Piece struct {
	Nom                   string
	PosX, PosY            int
	Liaisons              []*Relation
	Start, End, Occuppied bool
	Visited               bool
}

// get all the linked unoccupied rooms
func (p *Piece) GetSalles() []*Piece {
	tabDispo := []*Piece{}
	for _, l := range p.Liaisons {
		if l.Salle1 != p {
			tabDispo = append(tabDispo, l.Salle1)
		}
		if l.Salle2 != p {
			tabDispo = append(tabDispo, l.Salle2)
		}
	}
	return tabDispo
}

// get all the tunnels to adjacent rooms
func (p *Piece) GetLiaison(piece *Piece) *Relation {
	for _, l := range p.Liaisons {
		if l.Salle1 == piece || l.Salle2 == piece {
			return l
		}
	}
	return nil
}

// create a tunnel
func (p *Piece) NewLiaison(nouvPiece *Piece) *Relation {
	relation := &Relation{false, p, nouvPiece}
	nouvPiece.Liaisons = append(nouvPiece.Liaisons, relation)
	p.Liaisons = append(p.Liaisons, relation)
	return relation
}

func (p *Piece) AffichagePiece() string {
	/*strFinal := fmt.Sprintf("%s -> ", p.Nom)
	for _, s := range p.GetSalles() {
		strFinal += fmt.Sprintf("<- %s ", s.Nom)
	}
	return strFinal*/
	strFinal := ""
	if p.Start {
		strFinal += "##start\n"
	} else if p.End {
		strFinal += "##end\n"
	}
	return strFinal + fmt.Sprintf("%s %d %d", p.Nom, p.PosX, p.PosY)
}

func NewEnd(strSalle string) *Piece {
	salle := NewPiece(strSalle)
	salle.End = true
	return salle
}

func NewStart(strSalle string) *Piece {
	salle := NewPiece(strSalle)
	salle.Start = true
	salle.Occuppied = true
	return salle
}

func NewPiece(strSalle string) *Piece {
	tabStr := strings.Split(strSalle, " ")
	X, _ := strconv.Atoi(tabStr[1])
	Y, _ := strconv.Atoi(tabStr[2])
	return &Piece{tabStr[0], X, Y, []*Relation{}, false, false, false, false}
}
