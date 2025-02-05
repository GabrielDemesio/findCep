package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func main() {
	fmt.Println("Digite um cep: ")
	var inputCep string
	fmt.Scanf("%s", &inputCep)

	resp, err := http.Get("https://viacep.com.br/ws/" + inputCep + "/json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		return
	}

	var data Cep
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing json: %v\n", err)
		return
	}
	file, err := os.Create("city.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLocalidade: %s\nUF: %s\n", data.Cep, data.Localidade, data.Uf))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
	}
	// testee de cmmit
}
