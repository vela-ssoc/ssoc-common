package shipx

import "github.com/xgfone/ship/v5"

type RouteRegister interface {
	RegisterRoute(rgb *ship.RouteGroupBuilder) error
}

func RegisterRoutes(base *ship.RouteGroupBuilder, rrs []RouteRegister) error {
	for _, rr := range rrs {
		if rr == nil {
			continue
		}

		if err := rr.RegisterRoute(base); err != nil {
			return err
		}
	}

	return nil
}
