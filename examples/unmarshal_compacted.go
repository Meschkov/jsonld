package main

import (
	"fmt"
	"github.com/Meschkov/jsonld"
)

func main() {
	data := `{
	  "@id": "https://onerecordserver.de/address/MaxMustermann",
	  "@type": "https://onerecord.iata.org/cargo#Address",
	  "https://onerecord.iata.org/cargo#Address#Something": {
		"@type": "https://onerecord.iata.org/cargo#Something",
		"https://onerecord.iata.org/cargo#Address#Something#list": [
		  {
			"@type": "https://onerecord.iata.org/cargo#Nested",
			"https://onerecord.iata.org/cargo#Address#Something#list#Data": "testdata 1"
		  },
		  {
			"@type": "https://onerecord.iata.org/cargo#Nested",
			"https://onerecord.iata.org/cargo#Address#Something#list#Data": "testdata 2"
		  }
		],
		"https://onerecord.iata.org/cargo#Address#Something#test": "Testdata"
	  },
	  "https://onerecord.iata.org/cargo#Address#cityName": "Irgendwo",
	  "https://onerecord.iata.org/cargo#Address#country": "Germany",
	  "https://onerecord.iata.org/cargo#Address#postalCode": "12345",
	  "https://onerecord.iata.org/cargo#Address#street": "Musterstraße 1"
	}`
	var x interface{}
	err := jsonld.UnmarshalCompacted([]byte(data), x)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(x)
}
