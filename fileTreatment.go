package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GettingFile() string {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR : Invalid number of args")
		os.Exit(1)
	}
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent := ""
	buf := bufio.NewScanner(file)
	for buf.Scan() {
		fileContent += buf.Text() + "\n"
	}
	if fileContent == "" {
		fmt.Println("ERROR : Empty file")
		os.Exit(0)
	}
	return fileContent
}

func VerifyAnt(content string) int {
	countAnt := 0
	numberAnt := ""
	numberAntInt := 0
	contentSlice := strings.Split(content, "\n")
	isNumber := true
	for _, lineStr := range contentSlice {
		for _, ru := range lineStr {
			if ru < '0' || ru > '9' {
				isNumber = false
				break
			}
		}
		if isNumber && lineStr != "0" {
			countAnt++
			numberAnt = lineStr
		}
	}
	if countAnt != 1 {
		fmt.Println("ERROR : Invalid number of ants line")
		os.Exit(1)
	} else {
		numberAntInt2, err := strconv.Atoi(numberAnt)
		if err != nil {
			log.Fatal(err)
		}
		return numberAntInt2
	}
	return numberAntInt
}

func VerifyStartEnd(content string) (string, string) {
	countEnd := 0
	roomEnd := ""
	countStart := 0
	roomStart := ""
	contentSlice := strings.Split(content, "\n")
	isEnd := false
	isStart := false
	for _, lineStr := range contentSlice {
		if isEnd {
			roomEnd = lineStr
			isEnd = false
		} else if isStart {
			roomStart = lineStr
			isStart = false
		}
		if lineStr == "##start" {
			isStart = true
			countStart++
		} else if lineStr == "##end" {
			isEnd = true
			countEnd++
		}
	}
	if countStart != 1 {
		fmt.Println("ERROR : Invalid number of Start line")
		os.Exit(2)
	} else if countEnd != 1 {
		fmt.Println("ERROR : Invalid number of End line")
		os.Exit(2)
	}
	if !isRoom(roomEnd) || !isRoom(roomStart) {
		fmt.Println("ERROR : Invalid start/end room format")
		os.Exit(3)
	}
	return roomStart, roomEnd
}
func isRoom(str string) bool {
	contentStr := strings.Split(str, " ")
	if len(contentStr) != 3 {
		return false
	} else if contentStr[0][0] == 'L' || contentStr[0][0] == '#' {
		fmt.Println("ERROR : Invalid room name")
		os.Exit(4)
	}
	return true
}

func VerifyRoom(content string) []string {
	roomSlice := []string{}
	contentSlice := strings.Split(content, "\n")
	for _, lineStr := range contentSlice {
		contentLine := strings.Split(lineStr, " ")
		if len(contentLine) == 3 {
			if contentLine[0][0] == 'L' || contentLine[0][0] == '#' {
				fmt.Println("ERROR : Invalid room name")
				os.Exit(4)
			}
			coordX := contentLine[1]
			coordY := contentLine[2]
			if _, errX := strconv.Atoi(coordX); errX == nil {
				if _, errY := strconv.Atoi(coordY); errY == nil {
					roomSlice = append(roomSlice, lineStr)
				}
			}

		}
	}
	if len(roomSlice) < 2 {
		fmt.Println("ERROR : Invalid number of room")
		os.Exit(5)
	}
	nameMap := make(map[string]bool)
	coordMap := make(map[string]bool)
	for _, room := range roomSlice {
		roomSplit := strings.Split(room, " ")
		name := roomSplit[0]
		coord := roomSplit[1] + " " + roomSplit[2]

		if nameMap[name] {
			fmt.Println("ERROR : Duplicate room name:", name)
			os.Exit(6)
		}
		nameMap[name] = true

		if coordMap[coord] {
			fmt.Println("ERROR : Duplicate room coord:", coord)
			os.Exit(6)
		}
		coordMap[coord] = true
	}
	return roomSlice
}

func VerifyRelation(content string, allRooms []string) []string {
	relaSlice := []string{}
	contentSlice := strings.Split(content, "\n")
	for _, lineStr := range contentSlice {
		contentLine := strings.Split(lineStr, "-")
		if len(contentLine) == 2 {
			relaSlice = append(relaSlice, lineStr)
		}
	}
	if len(relaSlice) < 1 {
		fmt.Println("ERROR : Invalid number of relation")
		os.Exit(7)
	}

	relaMap := make(map[string]bool)
	roomNameMap := make(map[string]bool)
	for _, room := range allRooms {
		roomContent := strings.Split(room, " ")
		roomNameMap[roomContent[0]] = true
	}
	for _, rela := range relaSlice {
		roomRela := strings.Split(rela, "-")

		if !roomNameMap[roomRela[0]] || !roomNameMap[roomRela[1]] {
			fmt.Println("ERROR : Room related to unknown room :", rela)
			os.Exit(8)
		}

		relationKey := roomRela[0] + "-" + roomRela[1]
		reverseKey := roomRela[1] + "-" + roomRela[0]
		if relationKey == reverseKey {
			fmt.Println("ERROR : Room related to itself :", rela)
			os.Exit(9)
		}
		if relaMap[relationKey] || relaMap[reverseKey] {
			fmt.Println("ERROR : Duplicate relation :", rela)
			os.Exit(9)
		}
		relaMap[relationKey] = true
	}
	return relaSlice
}
