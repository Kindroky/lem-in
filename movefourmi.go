package main

func Assignerfourmi(tabFourmi []*Fourmi, groupe [][]*Piece) {
	compteurniveau := 0
	compteurfourmi := 0

	for compteurfourmi < len(tabFourmi) {
		for _, chemin := range groupe {
			if compteurniveau < len(chemin) && compteurfourmi < len(tabFourmi) {
				tabFourmi[compteurfourmi].Chemin = chemin
				compteurfourmi++
			}
		}
		compteurniveau++
	}
}

func Avancefourmi(tabFourmi []*Fourmi) {
	for _, fourmi := range tabFourmi {
		for index, salle := range fourmi.Chemin {
			if fourmi.Salle.Nom == salle.Nom && !fourmi.Salle.End {
				fourmi.Deplacer(fourmi.Chemin[index+1])
				break
			}
		}
	}
}
