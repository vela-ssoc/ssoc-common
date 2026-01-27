package dal_test

import (
	"testing"

	"github.com/vela-ssoc/ssoc-common/datalayer/model"
	"gorm.io/gen"
)

func TestGen(t *testing.T) {
	cfg := gen.Config{
		OutPath: "query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	}
	g := gen.NewGenerator(cfg)
	g.ApplyBasic(model.All()...)
	g.Execute()
}
