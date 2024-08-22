package main

import (
	"fmt"
)

func main() {
	msg := seyHello("go")

	msg = seyGoodbye("go")

	msg = seyThankYou("go")

	msg = seyWelcome("go")

	msg = seyGoodMorning("go")

	msg = seyGoodNight("go")

	msg = seyGoodLuck("go")

	msg = seyCongratulations("go")

	msg = seySorry("go")

	msg = seySeeYou("go")
	
	
	println(msg)
	

}

func seyHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func seyGoodbye(name string) string {
	return fmt.Sprintf("Goodbye %s", name)
}

func seyThankYou(name string) string {
	return fmt.Sprintf("Thank you %s", name)
}

func seyWelcome(name string) string {
	return fmt.Sprintf("Welcome %s", name)
}

func seyGoodMorning(name string) string {
	return fmt.Sprintf("Good morning %s", name)
}

func seyGoodNight(name string) string {
	return fmt.Sprintf("Good night %s", name)
}

func seyGoodLuck(name string) string {
	return fmt.Sprintf("Good luck %s", name)
}

func seyCongratulations(name string) string {
	return fmt.Sprintf("Congratulations %s", name)
}

func seySorry(name string) string {
	return fmt.Sprintf("Sorry %s", name)
}

func seySeeYou(name string) string {
	return fmt.Sprintf("See you %s", name)
}