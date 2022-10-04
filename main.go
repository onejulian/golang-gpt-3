package main

import (
	"context"
	"fmt"
	"gpt3/config"
	"log"
	"os"
	"strconv"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {

	entrada := "Escribir un artículo en español sobre los beneficios del agua de clorofila para adelgazar según nutricionistas 800"

	var err error
	var result string
	longitudEntrada := len(entrada)
	var response string
	var espacios int = 1
	palabras, err := strconv.Atoi(string(entrada[len(entrada)-3:]))
	if err != nil {
		log.Panic("no hay número al final", err)
	}
	entrada = entrada[:longitudEntrada-3]
	fmt.Println(palabras)

	response, err = ObtenerRespuesta(entrada)
	if err != nil {
		log.Println(err)
	}
	result += response

	for i := 0; i < len(result); i++ {
		if result[i] == ' ' {
			espacios += 1
		}
	}
	fmt.Println(espacios)
	
	for espacios < palabras {

		espaciosAntes := espacios
		
		response, err = ObtenerRespuesta(result)
		result += response
		if err != nil {
			log.Println(err)
		}
		// fmt.Println(response)
		
		espacios = 1
		
		for i := 0; i < len(result); i++ {
			if result[i] == ' ' {
				espacios += 1
			}
		}
		fmt.Println(espacios)

		if espaciosAntes == espacios {
			result += " pero"
		}
		// if espaciosAntes == espacios {
			// break
		// }

	}
	f, err := os.Create("output.txt")

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	_, err = f.WriteString(result)

	if err != nil {
		log.Println(err)
	}

}

func ObtenerRespuesta(promt string) (string, error) {

	app := config.InstanceApp()
	c := gogpt.NewClient(app.Gpt3Token())
	ctx := context.Background()
	var result string

	req := gogpt.CompletionRequest{
		Model:       "text-davinci-002",
		MaxTokens:   256,
		Prompt:      promt,
		Temperature: 0.9,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return result, err
	}

	return resp.Choices[0].Text, nil

}
