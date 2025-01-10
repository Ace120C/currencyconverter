package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/charmbracelet/huh"
	// "github.com/charmbracelet/lipgloss"
	"os"
	"strconv"
) 


type Currencies struct {
  Rates map[string]float64 `json:"data"`
}

func main() {
	// accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

  var choosenCurrency, choosenCurrency2, ammount string
  var  result float64
  var JSONdata Currencies

  //open the json file
  JSONfile, err := os.ReadFile("./api/currencies.json")
  if err != nil {
    panic(err)
  }
  
  err = json.Unmarshal(JSONfile, &JSONdata)
  if err !=  nil {
    log.Fatalf("error: %v", err)
  }
  
	// Create the form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Currency Converter").
				Description("Welcome to Currency Converter.\n").
				Next(true).
				NextLabel("Next"),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("USD $", "EUR €")...).
				Title("Choose Your Currency").
				Value(&choosenCurrency),
	     huh.NewSelect[string]().
				Options(huh.NewOptions("USD $", "JPY ¥", "EUR €")...).
	       Title("Which Currency to convert to? ").
				Value(&choosenCurrency2),
	     ),
	   huh.NewGroup(
	       huh.NewInput().
         Title("Ammount: ").
	       Prompt(">").
	       Value(&ammount),
	  ),
	 )

	err = form.Run()

  if err != nil {
    panic("illegal character: it needs to be a number!")
  }

  FAmmount, err := strconv.ParseFloat(ammount, 64)
  if err != nil {
    fmt.Printf("err: %v", err)
  }

  switch  {
  case choosenCurrency == "USD $" && choosenCurrency2 == "EUR €":
    result = FAmmount * JSONdata.Rates["EUR"]
  case choosenCurrency == "EUR €" && choosenCurrency2 == "USD $":
    result = FAmmount * JSONdata.Rates["USD"]
  case choosenCurrency == "USD $" && choosenCurrency2 == "JPY ¥":
    result = FAmmount * JSONdata.Rates["JPY"]
  case choosenCurrency == "EUR €" && choosenCurrency2 == "JPY ¥":
    result = FAmmount * JSONdata.Rates["JPY"]

  }  



  fmt.Println(result)

}
