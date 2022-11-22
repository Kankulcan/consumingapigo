package main

import (
	"testing"
)

func Test_create_struct(t *testing.T) {
	result := []Advice{
		{Message{Id: 8, Advice: "Sei lá"}},
		{Message{Id: 8, Advice: "Sei lá"}},
		{Message{Id: 8, Advice: "Sei lá"}},
		{Message{Id: 8, Advice: "Sei lá"}},
	}

	for _, e := range result {
		createArray := Message{e.Advices.Id, e.Advices.Advice}
		if createArray.Id != e.Advices.Id && createArray.Advice != e.Advices.Advice {
			t.Errorf("Aqui houve um erro de criação de Struct")
		}

	}

}
