package rest

import (
	"github.com/emicklei/go-restful/v3"
	"sync"
)

type ScoreBoard struct {
	mu         sync.Mutex
	scoreBoard []string
	history    int
}

type GetScoreboardResponse struct {
	Response []string `json:"response"`
}

type PutScoreboardResponse struct {
	Response []string `json:"response"`
}

// Add adds a score to score board
func (s *ScoreBoard) Add(score string) {
	s.mu.Lock()
	// Lock so only one goroutine at a time can access the scoreBoard
	s.scoreBoard = prependString(s.scoreBoard, score)
	if len(s.scoreBoard) >= s.history {
		s.scoreBoard = s.scoreBoard[:s.history]
	}
	s.mu.Unlock()
}

// Read returns the current score board
func (s *ScoreBoard) Read() []string {
	s.mu.Lock()
	// Lock so only one goroutine at a time can access the scoreBoard
	defer s.mu.Unlock()
	return s.scoreBoard
}

// Reset resets the current score board
func (s *ScoreBoard) Reset() {
	s.mu.Lock()
	// Lock so only one goroutine at a time can access the scoreBoard
	s.scoreBoard = []string{}

	s.mu.Unlock()
}

func (r *Service) GetScoreboardRequest(_ *restful.Request, response *restful.Response) {
	var responseJson GetScoreboardResponse

	responseJson.Response = r.scoreBoard.Read()

	_ = response.WriteAsJson(responseJson)
}

func (r *Service) PutScoreboardRequest(_ *restful.Request, response *restful.Response) {
	var responseJson GetScoreboardResponse

	r.scoreBoard.Reset()
	responseJson.Response = r.scoreBoard.Read()

	_ = response.WriteAsJson(responseJson)
}

func prependString(stringArray []string, stringToAdd string) []string {
	stringArray = append(stringArray, "")
	copy(stringArray[1:], stringArray)
	stringArray[0] = stringToAdd
	return stringArray
}
