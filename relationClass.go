package main

import "fmt"

type Relation struct {
	Utilise bool
	Salle1  *Piece
	Salle2  *Piece
}

// when a tunnel is used it passes the boolean as true
func (r *Relation) Use() {
	r.Utilise = true
}

//display links between rooms
func (r *Relation) AffichageRelation() string {
	return fmt.Sprintf("%s <-> %s", r.Salle1.Nom, r.Salle2.Nom)
}
