package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type  Cow struct{food, locomotion, sound string}
type Snake struct{food, locomotion, sound  string}
type Bird struct{food, locomotion, sound  string}

func (cow Cow) Eat(){fmt.Println(cow.food)}
func (cow Cow) Move(){fmt.Println(cow.locomotion)}
func (cow Cow) Speak(){fmt.Println(cow.sound)}

func (snake Snake) Eat(){fmt.Println(snake.food)}
func (snake Snake) Move(){fmt.Println(snake.locomotion)}
func (snake Snake) Speak(){fmt.Println(snake.sound)}

func (bird Bird) Eat(){fmt.Println(bird.food)}
func (bird Bird) Move(){fmt.Println(bird.locomotion)}
func (bird Bird) Speak(){fmt.Println(bird.sound)}

//These functions  simply delegate the call to the methods of the similar name
func  Eat(animal Animal){
	animal.Eat()
}
func Move(animal Animal){
	animal.Move()
}
func Speak(animal Animal){
	animal.Speak()
}


type AnimalNameIndex struct {
	animalKindMAP     map[string] string //It will be ideal to create this attrribute as set with the valid animal type i.e. cow, snake, bird
	nameAnimalMap     map[string] Animal
	animalActivityMap map[string] func(Animal)
}

func (animalNameIndex *AnimalNameIndex) InitAnimalNameIndex(){
	animalNameIndex.nameAnimalMap = make(map[string]Animal)
	animalNameIndex.animalKindMAP = map[string]string { // It would be ideal to use set here
		"cow":"cow",
		"snake":"snake",
		"bird":"bird",
	}
	animalNameIndex.animalActivityMap = map[string] func(Animal) {
		"eat": Eat,
		"move": Move,
		"speak": Speak,
	}
}

func Query(animalNameIndex AnimalNameIndex, name, activity string) error{
	name = strings.ToLower(strings.TrimSpace(name))
	activity = strings.ToLower(strings.TrimSpace(activity))
	var ok bool
	var activityFUNC func(Animal)
	var animal Animal
	animal, ok = animalNameIndex.nameAnimalMap[name]
	if !ok {
		return errors.New("Invalid animal name. The name has never been assigned to any animal so far. ")
	}
	activityFUNC, ok = animalNameIndex.animalActivityMap[activity]
	if !ok {
		return errors.New("Invalid animal activity")
	}
	activityFUNC(animal)
	return nil
}

func NewAnimal(animalNameIndex AnimalNameIndex, name, kind string) error{
	name = strings.ToLower(strings.TrimSpace(name))
	kind = strings.ToLower(strings.TrimSpace(kind))
	var ok bool

	_, ok = animalNameIndex.animalKindMAP[kind]
	if !ok {
		return errors.New("Invalid animal type")
	}
	_, ok = animalNameIndex.nameAnimalMap[name]
	if ok {
		return errors.New("Invalid animal name. The name is already assigned to other aanimal")
	}
	if kind == "cow"{
		animalNameIndex.nameAnimalMap[name] = Cow{food: "gass", locomotion: "walk", sound: "moo"}
	} else if kind == "snake" {
		animalNameIndex.nameAnimalMap[name] = Snake{food: "mice", locomotion: "slither", sound: "hss"}
	} else {
		animalNameIndex.nameAnimalMap[name] = Bird{food: "insects", locomotion: "fly", sound: "peep"}
	}
	return nil
}


func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	var animalNameIndex  AnimalNameIndex
	animalNameIndex.InitAnimalNameIndex()
	var err error
	var line string

	for {
		fmt.Println("> Enter the input. [Enter 'X' to terminate and 'Y' to see sample input commands] : ")
		fmt.Print("> ")
		line, err = consoleReader.ReadString('\n')
		if err != nil {
			continue
		}
		tokens := strings.Split(strings.TrimSpace(line), " ")
		if len(tokens) != 3 && len(tokens) != 1 {
			fmt.Println("Invalid input. continue ...")
			continue
		}
		if len(tokens) == 1  {
			if strings.HasPrefix(tokens[0], "X") {
				fmt.Println("Exiting ...")
				return
			} else if strings.HasPrefix(tokens[0], "Y") {
				fmt.Println("newanimal <name> <cow|snake|bird> \n" +
					"query <name> <move|eat|speak>"+
					"\n Example Input : \n newanimal bunny cow \n newanimal baby bird \n query bunny walk \n query baby walk\n")
				continue
			} else {
				fmt.Println("Invalid Input. Continue ...")
				continue
			}
		}

		command := strings.ToLower(strings.TrimSpace(tokens[0]))

		if strings.Compare(command, "newanimal") == 0 {
			err = NewAnimal(animalNameIndex, tokens[1], tokens[2])
			if err != nil {
				fmt.Println( err.Error())
			} else {
				fmt.Println(" New animal created.")
			}
		} else if strings.Compare(command, "query")  == 0 {
			err = Query(animalNameIndex, tokens[1], tokens[2])
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Invalid command. Continnue ...")
		}
	}

}
