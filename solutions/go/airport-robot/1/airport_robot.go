package airportrobot

import "fmt"

// Write your code here.
type Greeter interface{
    LanguageName() string
    Greet(name string) string
}

func SayHello(name string, greeter Greeter) string{
    greeting := fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
    return greeting
}

type Italian struct {}

func (i Italian) LanguageName() string{
    return "Italian"
}

func (i Italian) Greet(name string) string {
    return "Ciao " + name + "!"
}

type Portuguese struct{}

func (i Portuguese) LanguageName() string{
    return "Portuguese"
}

func (i Portuguese) Greet(name string) string {
    return "Olá " + name + "!"
}

// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
