package main

import (
  "time"
	"encoding/json"
	"fmt"
	"log"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
  "github.com/charmbracelet/huh/spinner"
	"os"
	"strconv"
) 


type Currencies struct {
  USRates map[string]float64 `json:"data"`
  EURates map[string]float64 `json:"EUdata"`
  JPRates map[string]float64 `json:"JPdata"`
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
				Options(huh.NewOptions("USD $", "EUR €", "JPY ¥")...).
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
    result = FAmmount * JSONdata.USRates["EUR"]
  case choosenCurrency == "USD $" && choosenCurrency2 == "JPY ¥":
    result = FAmmount * JSONdata.USRates["JPY"]
  case choosenCurrency == "EUR €" && choosenCurrency2 == "USD $":
    result = FAmmount * JSONdata.EURates["USD"]
  case choosenCurrency == "EUR €" && choosenCurrency2 == "JPY ¥":
    result = FAmmount * JSONdata.EURates["JPY"]
  case choosenCurrency == "JPY ¥" && choosenCurrency2 == "USD $":
    result = FAmmount * JSONdata.JPRates["USD"]
  case choosenCurrency == "JPY ¥" && choosenCurrency2 == "EUR €":
    result = FAmmount * JSONdata.JPRates["EUR"]
  }  

  if choosenCurrency == "USD $" && choosenCurrency2 == "USD $" {
    panic("You can't convert a currency with the exact same one!")
  }
  if choosenCurrency == "EUR €" && choosenCurrency2 == "EUR €" {
   panic("You can't convert a currency with the exact same one!") 
  }
  if choosenCurrency == "JPY ¥" && choosenCurrency2 == "JPY ¥" {
   panic("You can't convert a currency with the exact same one!") 
  }
  
  Sresult := strconv.FormatFloat(result, 'f', -1, 64)

  calculateCurrency := func ()  {
    time.Sleep(2 * time.Second)
  }

  _ = spinner.New().Title("Calculating Currency...").Action(calculateCurrency).Run()  

  var style = lipgloss.NewStyle().
  Width(40).
  BorderStyle(lipgloss.RoundedBorder()).
  BorderForeground(lipgloss.Color("63")).
  Padding(1, 2)

  fmt.Println(style.Render("result is: ", Sresult))

}
