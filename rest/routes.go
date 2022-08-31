package rest

import (
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

type Service struct {
	container         *restful.Container
	scoreBoard        ScoreBoard
	scoreBoardHistory int
}

func NewService() (*Service, error) {
	r := &Service{
		container:  restful.NewContainer(),
		scoreBoard: ScoreBoard{scoreBoard: make([]string, 0), history: 10},
	}

	r.container.Add(r.buildRoutes())

	if err := bootstrapSwagger(r.container); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Service) Container() *restful.Container {
	return r.container
}

func (r *Service) buildRoutes() *restful.WebService {
	ws := new(restful.WebService)

	ws.Path("/rpsls-api").
		Route(ws.GET("/health").
			Operation("health check").
			To(func(request *restful.Request, response *restful.Response) {
				_ = response.WriteErrorString(http.StatusOK, "rpsls-api is up and running!")
			}))

	ws.Route(ws.GET("/choices").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetChoicesResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.GetChoicesRequest).
		Writes(map[string]string{}))

	ws.Route(ws.GET("/choice").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetChoiceResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.GetChoiceRequest).
		Writes(map[string]string{}))

	ws.Route(ws.POST("/play").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), PlayResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.Play).
		Writes(map[string]string{}))

	ws.Route(ws.GET("/scoreboard").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), GetScoreboardResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.GetScoreboardRequest).
		Writes(map[string]string{}))

	ws.Route(ws.PUT("/scoreboard").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), PutScoreboardResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), EndpointErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), EndpointErrorResponse{}).
		To(r.PutScoreboardRequest).
		Writes(map[string]string{}))

	return ws
}
