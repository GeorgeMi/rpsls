package rest

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/rs/zerolog/log"
	"net/http"
)

type PlayRequest struct {
	Player int `json:"player"`
}

type PlayResponse struct {
	Results  string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}

func (r *Service) Play(request *restful.Request, response *restful.Response) {
	requestJson := new(PlayRequest)

	// parse body
	err := request.ReadEntity(requestJson)
	if err != nil {
		log.Error().Err(err).Msgf("unable to parse request: %s", err)
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("unable to parse request: %s", err.Error()))
		return
	}

	// validate request
	err = validatePlayRequest(*requestJson)
	if err != nil {
		log.Error().Err(err).Msgf("unable to validate request: %s", err)
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("unable to validate request: %s", err))
		return
	}

	var responseJson PlayResponse
	responseJson.Player = requestJson.Player
	responseJson.Computer, err = getRandomNumber()
	if err != nil {
		log.Error().Err(err).Msgf("unable to get random number from codechallenge: %s", err)
		buildEndPointErrorResponse(response, http.StatusInternalServerError, fmt.Sprintf("unable to get random number from codechallenge: %s", err))
		return
	}

	responseJson.Results = calculateWinner(responseJson.Player, responseJson.Computer)

	go r.scoreBoard.Add(responseJson.Results)

	_ = response.WriteAsJson(responseJson)
}

func validatePlayRequest(requestJson PlayRequest) error {
	if requestJson.Player < 0 || requestJson.Player > len(choicesMap) {
		return fmt.Errorf("invalid choice id")
	}

	return nil
}

func calculateWinner(a, b int) string {
	result := -1

	if a == b {
		return "tie"
	} else if a < b {
		for i := a + 1; i <= b; i++ {
			result = result * -1
		}
	} else {
		result = 1
		for i := a - 1; i >= b; i-- {
			result = result * -1
		}
	}

	if result == 1 {
		return "win"
	} else {
		return "lose"
	}
}
