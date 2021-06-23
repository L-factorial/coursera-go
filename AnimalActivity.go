package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//type Animal struct {
//	food, locomotion, noise string
//}
//func (animal *Animal) Eat(){
//	fmt.Println("> "+animal.food)
//}
//func (animal *Animal) Move(){
//	fmt.Println("> "+animal.locomotion)
//}
//func (animal *Animal) Speak(){
//	fmt.Println("> "+animal.noise)
//}
//func (animal *Animal) InitMe(food, locomotion, noise string){
//	animal.food = food
//	animal.locomotion = locomotion
//	animal.noise = noise
//}
//
////These functions are simply delegate the call to the methods of the similar name
//func  Eat(animal Animal){
//	animal.Eat()
//}
//func Move(animal Animal){
//	animal.Move()
//}
//func Speak(animal Animal){
//	animal.Speak()
//}
//
//func InitAnimal() map[string]Animal{
//	var cow Animal
//	cow.InitMe("grass", "walk", "moo")
//	var bird Animal
//	bird.InitMe("worms", "fly", "peep")
//	var snake Animal
//	snake.InitMe("mice", "slither", "hss")
//
//	var animalMAP = map[string]Animal {
//		"cow":cow,
//		"bird": bird,
//		"snake":snake,
//	}
//	return animalMAP
//}
//
//func InitActivity() map[string] func(Animal) {
//	var activityMAP = map[string] func(Animal) {
//		"eat": Eat,
//		"move": Move,
//		"speak": Speak,
//	}
//	return activityMAP
//}
//
//
//func main() {
//	var animalMAP = InitAnimal()
//	var activityMAP = InitActivity()
//
//	var animal, activity string
//	consoleReader := bufio.NewReader(os.Stdin)
//
//	for {
//		fmt.Println("> Enter the input <animalNAME anivamACTIVITY>.[ Example input : cow move] [Enter 'X' to terminate] : ")
//		fmt.Print("> ")
//		line, err := consoleReader.ReadString('\n')
//		if err != nil {
//			continue
//		}
//		tokens := strings.Split(line, " ")
//		if len(tokens) > 2 || len(tokens) < 0 {
//			fmt.Println("> Invalid input. continue ...")
//			continue
//		}
//		if len(tokens) == 1  {
//			if strings.HasPrefix(tokens[0], "X") {
//				fmt.Println("> Exiting ...")
//				return
//			}
//			fmt.Println("> Invalid Input. Continue ...")
//			continue
//		}
//		animal = strings.ToLower(strings.TrimSpace(tokens[0]))
//		activity = strings.ToLower(strings.TrimSpace(tokens[1]))
//
//		animalOBJ, flag1 := animalMAP[animal]
//		if !flag1 {
//			fmt.Println("> Invalid Animal name. Try again ...")
//			continue
//		}
//		activityFUNC, flag2 := activityMAP[activity]
//		if !flag2{
//			fmt.Println("> Invalid animal activity. Try Again ...")
//			continue
//		}
//
//		activityFUNC(animalOBJ)
//	}
//}
