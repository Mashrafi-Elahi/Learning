package main

import "fmt"

type Team struct {
    Name    string
    Players []string
}

func (t Team) AddPlayerValue(name string) {
    t.Players = append(t.Players, name)
}

func (t *Team) AddPlayerPointer(name string) {
    t.Players = append(t.Players, name)
}

func main() {
    myTeam := Team{Name: "Gophers", Players: []string{"Alice"}}

    myTeam.AddPlayerValue("Bob")
    myTeam.AddPlayerPointer("Charlie")

    fmt.Println(myTeam.Players)
}