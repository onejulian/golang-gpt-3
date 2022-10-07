package main

import (
	"context"
	"gpt3/config"
	"gpt3/util"
	"log"
	"math/rand"
	"os"
	"strconv"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {

	cadenaDeInstrucciones := `Carlos desea adelantar estudios a nivel tecnológico, su mayor problema radica en que no
	encuentra una institución cercana que le ofrezca la oportunidad de estudiar lo que desea.
	Por tal razón acude a un amigo para que lo oriente donde tomar la formación.
	A partir de lo anterior usted debe proponer alternativas para que Carlos pueda acceder a
	formación gratuita, incluyente y flexible 350`
	var err error
	var acumuladoFinalDeRespuestas string
	longitudEntrada := len(cadenaDeInstrucciones)
	var respuestaDesdeGpt3 string
	listaDeConjunciones := []string{" sin embargo,", " no obstante,", " pero lo más importante,", " por otro lado,", " tomando en cuenta lo anterior", " y por eso se tiene que", " por eso es bueno que", " complementando lo anterior", " pero"}

	maxPalabrasSolicitadas, err := strconv.Atoi(string(cadenaDeInstrucciones[len(cadenaDeInstrucciones)-3:]))
	if err != nil {
		log.Println("no hay número al final de la instrución", err)
	}
	cadenaDeInstrucciones = cadenaDeInstrucciones[:longitudEntrada-3]

	respuestaDesdeGpt3, err = obtenerRespuesta(cadenaDeInstrucciones)
	if err != nil {
		log.Println(err)
	}
	acumuladoFinalDeRespuestas += respuestaDesdeGpt3

	ctdPalabrasAcumulado := util.CountWords(acumuladoFinalDeRespuestas)

	for ctdPalabrasAcumulado <= maxPalabrasSolicitadas {

		acumuladoFinalDeRespuestas += listaDeConjunciones[rand.Intn(len(listaDeConjunciones))]

		respuestaDesdeGpt3, err = obtenerRespuesta(acumuladoFinalDeRespuestas)
		if err != nil {
			log.Println(err)
		}
		acumuladoFinalDeRespuestas += respuestaDesdeGpt3

		ctdPalabrasAcumulado = util.CountWords(acumuladoFinalDeRespuestas)

	}

	f, err := os.Create("output.txt")

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	_, err = f.WriteString(acumuladoFinalDeRespuestas)

	if err != nil {
		log.Println(err)
	}

}

func obtenerRespuesta(promt string) (string, error) {

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
