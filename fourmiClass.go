package main

import "fmt"

type Fourmi struct {
	Salle  *Piece //room it is located in
	Numero int    // name of the ant
	Chemin []*Piece
}

//check if we can move an ant to a specific new room and move it if possible
func (f *Fourmi) Deplacer(nouvSalle *Piece) bool {
	dispo := f.GetSallesDispo() //get all available rooms
	found := false              //set found to false because he hasn't yet found the nouvsalle room
	for _, p := range dispo {
		if p == nouvSalle { //if it founds the room, found is set to true and the loop is broken
			found = true
			break
		}
	}
	if !found { //if not found, print error message and return false
		fmt.Printf("L%dX%s ", f.Numero, nouvSalle.Nom)
		return false
	}
	tunnel := f.Salle.GetLiaison(nouvSalle) //get the tunnel connected from the room of the ant to nouvsalle
	if tunnel.Utilise {                     //if already used return false
		return false
	}
	tunnel.Use()                                   //otherwise, pass it as used
	f.Salle = nouvSalle                            //move the ant to the nouvsalle
	fmt.Printf("L%d-%s ", f.Numero, nouvSalle.Nom) //print the movement
	return true                                    //movement successful
}

//GetSalles but from a room an ant is located in
func (f *Fourmi) GetSallesDispo() []*Piece {
	return f.Salle.GetSalles()
}

//display info related to an ant
func (f *Fourmi) AffichageFourmi() string {
	return fmt.Sprintf("L%d : %s", f.Numero, f.Salle.Nom)
}

//create a new ant
func NewFourmi(salle *Piece, tabFourmi []*Fourmi) *Fourmi {
	return &Fourmi{salle, len(tabFourmi) + 1, []*Piece{}}
}
