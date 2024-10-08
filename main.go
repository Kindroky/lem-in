package main

import (
	"fmt"
	"strings"
)

func main() {
	fileContent := GettingFile()
	numberAntPoint, roomStartPoint, roomEndPoint, roomsPoint, relasPoint := VerifyFile(fileContent)
	tabSalles, tabFourmis, tabRelation := CreateAntEmpire(numberAntPoint, roomStartPoint, roomEndPoint, roomsPoint, relasPoint)
	fmt.Println(tabFourmis, tabRelation)

	allPaths := StartExploration(tabSalles[0])
	/*for _, l := range allPaths {
		for _, s := range l {
			fmt.Print(s.Nom)
		}
		fmt.Println()
	}*/
	allGroups := IndiePaths(allPaths)
	allGroups = SortPath(allGroups)
	bestGroups := FindGroupsWithMostPaths(allGroups)
	bestGroup := FindShortestPath(bestGroups)
	PrintGroups(allGroups)
	fmt.Println("---------")
	PrintGroups(bestGroups)
	fmt.Println("---------")
	PrintGroups(append([][][]*Piece{}, bestGroup))
	Assignerfourmi(tabFourmis, bestGroup)
	fmt.Println(tabFourmis[2])
	for !tabFourmis[len(tabFourmis)-1].Salle.End {
		for _, check := range tabRelation {
			check.Utilise = false
		}
		Avancefourmi(tabFourmis)
		fmt.Println()
	}
	//fourmis()
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

/*func fourmis() {
	tabSalles := []*Piece{}
	tabFourmis := []*Fourmi{}
	tabRelation := []*Relation{}

	tabSalles = append(tabSalles, NewStart("start 1 1"))
	tabSalles = append(tabSalles, NewEnd("end 10 5"))
	tabSalles = append(tabSalles, NewPiece("cuisine 4 4"))
	tabSalles = append(tabSalles, NewPiece("salon 8 2"))

	tabRelation = append(tabRelation, tabSalles[0].NewLiaison(tabSalles[2])) // S---\
	tabRelation = append(tabRelation, tabSalles[2].NewLiaison(tabSalles[1])) //  \   s
	tabRelation = append(tabRelation, tabSalles[3].NewLiaison(tabSalles[0])) //   \ / \
	tabRelation = append(tabRelation, tabSalles[2].NewLiaison(tabSalles[3])) //    c   \
	tabRelation = append(tabRelation, tabSalles[3].NewLiaison(tabSalles[1])) //     \___E

	tabFourmis = append(tabFourmis, NewFourmi(tabSalles[0], tabFourmis))
	tabFourmis = append(tabFourmis, NewFourmi(tabSalles[0], tabFourmis))
	tabFourmis = append(tabFourmis, NewFourmi(tabSalles[0], tabFourmis))

	for _, s := range tabSalles { //afficher toute les salles
		fmt.Println(s.AffichagePiece())
	}
	fmt.Print("\n")
	for _, r := range tabRelation { //afficher toute les relations
		fmt.Println(r.AffichageRelation())
	}
	fmt.Print("\n")
	for _, f := range tabFourmis { //afficher toute les fourmis
		fmt.Println(f.AffichageFourmi())
	}
	fmt.Print("\nChoix Fourmi 1 :\n") //afichier les choix de la premiere fourmi
	for _, s := range tabFourmis[0].GetSallesDispo() {
		fmt.Println(s.AffichagePiece())
	}
	tabFourmis[0].Deplacer(tabSalles[1]) //deplacement vers salle inconnue
	tabFourmis[0].Deplacer(tabSalles[2]) //deplacement vers salle connue
}*/
