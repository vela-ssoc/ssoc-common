package model

type ExposeAddress struct {
	Name string `bson:"name" json:"name" validate:"required"`
	Addr string `bson:"addr" json:"addr" validate:"required"`
}

type ExposeAddresses []*ExposeAddress

func (eas ExposeAddresses) Addresses() []string {
	rets := make([]string, 0, 10)
	uniq := make(map[string]struct{}, 8)
	for _, ea := range eas {
		addr := ea.Addr
		if _, exists := uniq[addr]; !exists {
			uniq[addr] = struct{}{}
			rets = append(rets, addr)
		}
	}

	return rets
}
