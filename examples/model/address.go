package model

type Address struct {
	ID         string    `jsonld:"@id"`
	Type       string    `jsonld:"@type" default:"https://onerecord.iata.org/cargo#Address"`
	Street     string    `jsonld:"https://onerecord.iata.org/cargo#Address#street"`
	PostalCode string    `jsonld:"https://onerecord.iata.org/cargo#Address#postalCode"`
	CityName   string    `jsonld:"https://onerecord.iata.org/cargo#Address#cityName"`
	Country    string    `jsonld:"https://onerecord.iata.org/cargo#Address#country"`
	Something  Something `jsonld:"https://onerecord.iata.org/cargo#Address#Something"`
}

type Something struct {
	Test string   `jsonld:"https://onerecord.iata.org/cargo#Address#Something#test"`
	Type string   `jsonld:"@type" default:"https://onerecord.iata.org/cargo#Something"`
	List []Nested `jsonld:"https://onerecord.iata.org/cargo#Address#Something#list"`
}

type Nested struct {
	Type string `jsonld:"@type" default:"https://onerecord.iata.org/cargo#Nested"`
	Data string `jsonld:"https://onerecord.iata.org/cargo#Address#Something#list#Data"`
}
