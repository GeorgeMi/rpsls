package rest

import (
	"net/http"
)

type Service struct {
	container *restful.Container
	blackList []int
}

func NewService() (*Service, error) {
	r := &Service{
		container: restful.NewContainer(),
		blackList: make([]int, 0),
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
	return ws
}
