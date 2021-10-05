package arkimmenintuisimemessolardinsheshelerdinamindarisiktimmen_test

import (
	"testing"

	"github.com/arkimmenintuisimemesolsin/arkimmenintuisimemessolardinsheshelerdinamindarisiktimmen"
)

func TestItIsTrue(t *testing.T) {
	for _, yes := range arkimmenintuisimemessolardinsheshelerdinamindarisiktimmen.ItIsTrue() {
		if yes == "" {
			t.Error("it's impossible")
		}
	}
}
