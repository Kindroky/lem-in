package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	timer := time.Now()
	fileContent := GettingFile()
	numberAntPoint, roomStartPoint, roomEndPoint, roomsPoint, relasPoint := VerifyFile(fileContent)
	tabSalles, tabFourmis, tabRelation := CreateAntEmpire(numberAntPoint, roomStartPoint, roomEndPoint, roomsPoint, relasPoint)

	allPaths := StartExploration(tabSalles[0])

	allGroups := IndiePaths(allPaths)
	bestGroups := FindGroupsWithMostPaths(allGroups)
	bestGroup := TrimGroup(FindShortestPath(bestGroups))
	/*PrintGroups(allGroups)
	fmt.Println("---------")
	PrintGroups(bestGroups)
	fmt.Println("---------")
	PrintGroups(append([][][]*Piece{}, bestGroup))*/
	Assignerfourmi(tabFourmis, bestGroup)
	/*for _, f := range tabFourmis {
		fmt.Println(*f)
	}*/

	fmt.Println(len(tabFourmis))
	for _, p := range tabSalles {
		fmt.Println(p.AffichagePiece())
	}
	for _, l := range tabRelation {
		fmt.Println(l.AffichageRelation())
	}
	fmt.Println()

	compteur := 0
	for !tabFourmis[len(tabFourmis)-1].Salle.End {
		compteur++
		for _, check := range tabRelation {
			check.Utilise = false
		}
		Avancefourmi(tabFourmis)
		fmt.Println()
	}
	fmt.Println("temp : ", time.Since(timer))
	fmt.Println("nombre de tours :", compteur)
}

func VerifyFile(content string) (*int, *string, *string, []string, []string) {
	numberAnt := VerifyAnt(content)
	roomStart, roomEnd := VerifyStartEnd(content)
	rooms := VerifyRoom(content)
	relas := VerifyRelation(content, rooms)

	return &numberAnt, &roomStart, &roomEnd, rooms, relas
}

func CreateAntEmpire(num *int, start *string, end *string, rooms []string, relas []string) ([]*Piece, []*Fourmi, []*Relation) {
	tabSalles := []*Piece{}
	tabFourmis := []*Fourmi{}
	tabRelation := []*Relation{}

	tabSalles = CreateSalle(tabSalles, start, end, rooms)
	tabFourmis = CreateFourmis(tabFourmis, tabSalles, num)
	tabRelation = CreateRelation(tabRelation, tabSalles, relas)

	return tabSalles, tabFourmis, tabRelation
}

func CreateSalle(tabSalles []*Piece, start *string, end *string, rooms []string) []*Piece {
	tabSalles = append(tabSalles, NewStart(*start))
	tabSalles = append(tabSalles, NewEnd(*end))
	for _, room := range rooms {
		if room != *start && room != *end {
			tabSalles = append(tabSalles, NewPiece(room))
		}
	}
	return tabSalles
}
func CreateFourmis(tabFourmis []*Fourmi, tabSalles []*Piece, num *int) []*Fourmi {
	for i := 0; i < *num; i++ {
		tabFourmis = append(tabFourmis, NewFourmi(tabSalles[0], tabFourmis))
	}
	return tabFourmis
}
func CreateRelation(tabRelation []*Relation, tabSalles []*Piece, relas []string) []*Relation {
	for _, relation := range relas {
		namesInRel := strings.Split(relation, "-")
		index1 := 0
		index2 := 0
		for index, salle := range tabSalles {
			if salle.Nom == namesInRel[0] {
				index1 = index
			} else if salle.Nom == namesInRel[1] {
				index2 = index
			}
		}
		tabRelation = append(tabRelation, tabSalles[index1].NewLiaison(tabSalles[index2]))
	}
	return tabRelation
}
