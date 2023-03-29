package endpoints

import (
	"context"

	"github.com/aszeta/ig/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	Add endpoint.Endpoint
}

// MathReq struct holds the endpoint request definition
type MathReq struct {
	NumA float32
	NumB float32
}

// MathResp struct holds the endpoint response definition
type MathResp struct {
	Result float32
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Add(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}
