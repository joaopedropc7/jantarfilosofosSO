package main

import (
	"fmt"
	"time"
)

type filosofo struct {
	nomeFilosofo      string
	processoTotal     int64
	processoRealizado int64
	concluido         bool
	status            status
}

type status int

const (
	Executando  status = 0
	Concluido   status = 1
	Descansando status = 2
)

func (s status) String() string {
	switch s {
	case Executando:
		return "Executando"
	case Concluido:
		return "Concluído"
	case Descansando:
		return "Descansando"
	default:
		return "Desconhecido"
	}
}

func main() {
	var filosofos [5]filosofo
	filosofos[0] = filosofo{"Filoso Tadeu", 100, 0, false, 2}
	filosofos[1] = filosofo{"Filoso Jonatan", 100, 0, false, 2}
	filosofos[2] = filosofo{"Filoso Icaro", 100, 0, false, 2}
	filosofos[3] = filosofo{"Filoso João", 100, 0, false, 2}
	filosofos[4] = filosofo{"Filoso Manel", 100, 0, false, 2}

	var processosConcluidos int = 0

	for processosConcluidos < 5 {
		for i := 0; i < len(filosofos); i++ {
			filosofos[i].status = status(Executando)
			time.Sleep(2 * time.Second)
			filosofos[i].processoRealizado += 20
			fmt.Println(filosofos[i].nomeFilosofo, filosofos[i].status)
			filosofos[i].status = status(Descansando)
			fmt.Println(filosofos[i].nomeFilosofo, filosofos[i].status)
			time.Sleep(2 * time.Second)
			if filosofos[i].processoRealizado == filosofos[i].processoTotal {
				filosofos[i].concluido = true
				filosofos[i].status = status(Concluido)
				fmt.Println(filosofos[i].nomeFilosofo, "Processo concluido")
				time.Sleep(2 * time.Second)
				processosConcluidos++
			}

		}
		for i := 0; i < len(filosofos); i++ {
			fmt.Println(filosofos[i].nomeFilosofo, filosofos[i].processoRealizado, "% ", filosofos[i].status.String())
		}
	}
}
