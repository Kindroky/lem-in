echo "Try to run 'go run . example00.txt'" $'\n'
go run . TestFiles/example00.txt
echo "Is the program able to read the ant farm in this standard input ?" $'\n'
echo "Does the program accept only the commands ##start and ##end?" $'\n'
echo "Are the commands and the ants movements printed with the right format?" $'\n'$'\n'

echo "example00.txt, 6 moves or less"$'\n'
go run . TestFiles/example00.txt

echo $'\n'"example01.txt, 8 moves or less"$'\n'
go run . TestFiles/example01.txt

echo $'\n'"example02.txt, 11 moves or less"$'\n'
go run . TestFiles/example02.txt

echo $'\n'"example03.txt, 6 moves or less"$'\n'
go run . TestFiles/example03.txt

echo $'\n'"example04.txt, 6 moves or less"$'\n'
go run . TestFiles/example04.txt

echo $'\n'"example05.txt, 8 moves or less"$'\n'
go run . TestFiles/example05.txt

echo $'\n'"badexample00.txt"$'\n'
go run . TestFiles/badexample00.txt

echo $'\n'"badexample01.txt"$'\n'
go run . TestFiles/badexample01.txt

echo $'\n'"example06.txt, less than 1.5 minutes"$'\n'
go run . TestFiles/example06.txt

echo $'\n'"example07.txt, less than 2.5 minutes"$'\n'
go run . TestFiles/example07.txt