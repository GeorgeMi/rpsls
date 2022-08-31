package rest

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/rs/zerolog/log"
	"net/http"
)

var choicesMap = map[int]string{
	0: "paper",
	1: "rock",
	2: "lizard",
	3: "spock",
	4: "scissors",
}

type GetChoicesRequest struct {
}

type GetChoicesResponse struct {
	Response []Choice `json:"response"`
}

type GetChoiceRequest struct {
}

type GetChoiceResponse struct {
	Response Choice `json:"response"`
}

type GetRandomNumberResponse struct {
	RandomNumber int `json:"random_number"`
}

type Choice struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r *Service) GetChoicesRequest(_ *restful.Request, response *restful.Response) {
	var responseJson GetChoicesResponse

	for k, v := range choicesMap {
		responseJson.Response = append(responseJson.Response, Choice{
			Id:   k,
			Name: v,
		})
	}

	response.WriteAsJson(responseJson)
}

func (r *Service) GetChoiceRequest(_ *restful.Request, response *restful.Response) {
	var responseJson GetChoiceResponse

	// get random number
	randomNuber, err := getRandomNumber()
	if err != nil {
		log.Error().Err(err).Msgf("unable to get random number from codechallenge: %s", err)
		buildEndPointErrorResponse(response, http.StatusInternalServerError, fmt.Sprintf("unable to get random number from codechallenge: %s", err))
		return
	}

	responseJson.Response = Choice{
		Id:   randomNuber,
		Name: choicesMap[randomNuber],
	}

	_ = response.WriteAsJson(responseJson)
}

func getRandomNumber() (int, error) {
	req, err := http.NewRequest("GET", "https://codechallenge.boohma.com/random", nil)
	if err != nil {
		return -1, fmt.Errorf("error creating the request for codechallenge api: %s", err.Error())
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, fmt.Errorf("error requesting codechallenge api: %s", err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var r GetRandomNumberResponse
	err = dec.Decode(&r)
	if err != nil {
		return -1, fmt.Errorf("error decoding codechallenge response: %s", err.Error())
	}

	if r.RandomNumber < 1 || r.RandomNumber > 100 {
		return -1, fmt.Errorf("got a number that is not in 1-100 interval")
	}

	// compute random mapKey
	response := r.RandomNumber % len(choicesMap)

	return response, nil
}
