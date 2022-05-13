package main

import (
	"fmt"
	"github.com/Meschkov/jsonld"
	"github.com/Meschkov/jsonld/examples/model"
)

func main() {
	data := model.Address{
		ID:         "https://onerecordserver.de/address/MaxMustermann",
		Street:     "Musterstraße 1",
		PostalCode: "12345",
		CityName:   "Irgendwo",
		Country:    "Germany",
		Something: model.Something{
			Test: "Testdata",
			List: []model.Nested{
				model.Nested{Data: "testdata 1"},
				model.Nested{Data: "testdata 2"},
			},
		},
	}

	j, err := jsonld.MarshalCompacted(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
}
