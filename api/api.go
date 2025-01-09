package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Currencies struct {
  Data map[string]float64 `json:"data"`
}

func main()  {

  ApiUrl := "API_URL"
  response, err := http.Get(ApiUrl) 
  if err != nil {
    log.Fatalf("Failed to get the database: %v", err)
  }
  
  databaseData := response.Body

  buf := new(bytes.Buffer)
  buf.ReadFrom(databaseData)
  responseBytes := buf.String()

  var currency Currencies
  
  err = json.Unmarshal([]byte(responseBytes), &currency)
  if err != nil {
    panic(err)
  }
  
  dataJSON, err := json.MarshalIndent(responseBytes, "", " ")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(dataJSON))


  dbFileJSON := "currencies.json"
  FileJSON, err := os.Create(dbFileJSON)
  if err != nil {
    log.Fatal(err)
  }
  defer FileJSON.Close()

  _, err = FileJSON.WriteString(string(dataJSON))
  if err != nil {
    log.Fatalf("Error writing to file: %v", err)
  }
  fmt.Println("File Created Successfully")
}
