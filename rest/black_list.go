package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful/v3"
	"github.com/rs/zerolog/log"
)

type GetBlackListResponse struct {
	BlackList []int `json:"blackList"`
}

type BlackListRequest struct {
	BlackListValue *int `json:"BlackListValue"`
}

func (r *Service) GetBlackList(request *restful.Request, response *restful.Response) {
	var responseJSON GetBlackListResponse

	// retrieve global blacklist array
	responseJSON.BlackList = r.blackList

	// write the response
	_ = response.WriteAsJson(responseJSON)
}

func (r *Service) AddBlackListValue(request *restful.Request, response *restful.Response) {
	// parse query string
	requestQuery, err := readBlackListRequest(request)
	if err != nil {
		log.Error().Err(err).Msgf("unable to parse request: %s", err)
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("unable to parse request: %s", err.Error()))
		return
	}

	// validate request
	err = validateBlackListRequest(requestQuery)
	if err != nil {
		log.Error().Err(err).Msgf("unable to validate request: %s", err)
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("unable to validate request: %s", err))
		return
	}

	// add value to global blacklist array if it not already exists
	if !checkIfExists(r.blackList, *requestQuery.BlackListValue) {
		r.blackList = append(r.blackList, *requestQuery.BlackListValue)
	}
}

func (r *Service) DeleteBlackListValue(request *restful.Request, response *restful.Response) {
	// parse query string
	requestQuery, err := readBlackListRequest(request)
	if err != nil {
		log.Error().Err(err).Msgf("unable to parse request: %s", err)
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}

	// validate request
	err = validateBlackListRequest(requestQuery)
	if err != nil {
		log.Error().Err(err).Msgf("unable to validate request: %s", err)
		buildEndPointErrorResponse(response, http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}

	// delete value from global blacklist array if it exists
	r.blackList = removeValueFromArray(r.blackList, *requestQuery.BlackListValue)

}

func readBlackListRequest(request *restful.Request) (requestQuery BlackListRequest, err error) {
	requestQuery.BlackListValue = new(int)

	if len(request.PathParameter(BlackListParameter)) > 0 {
		*requestQuery.BlackListValue, err = strconv.Atoi(request.PathParameter(BlackListParameter))
	}

	return
}

func validateBlackListRequest(requestQuery BlackListRequest) error {
	if requestQuery.BlackListValue == nil {
		return fmt.Errorf("no blackListValue provided in request")
	}

	if *requestQuery.BlackListValue < 0 {
		return fmt.Errorf("invalid blackListValue")
	}

	return nil
}

// search value within array
func checkIfExists(array []int, val int) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

// search value within array and delete
func removeValueFromArray(array []int, val int) []int {
	for i, v := range array {
		if v == val {
			array[i] = array[len(array)-1]
			return array[:len(array)-1]
		}
	}

	return array
}
