package meander

import (
	"testing"

	"gitlab.com/Orenge/blueprints/meander"

	"github.com/cheekybits/is"
)

func TestParseCostRange(t *testing.T) {
	is := is.New(t)
	var l *meander.CostRange
	l = meander.ParseCostRange("$$...$$$")
	is.Equal(l.From, meander.Cost2)
	is.Equal(l.To, meander.Cost3)
	l = meander.ParseCostRange("$...$$$$$")
	is.Equal(l.From, meander.Cost1)
	is.Equal(l.To, meander.Cost5)
}
func TestCostRangeString(t *testing.T) {
	is := is.New(t)
	is.Equal("$$...$$$$", (&meander.CostRange{
		From: meander.Cost2,
		To:   meander.Cost4,
	}).String())
}
