package gaussian

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestCDF(t *testing.T) {
	points := []float64{
		-4.0, -3.5, -3.0, -2.5, -2.0, -1.5, -1.0, -0.5,
		0.0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0,
	}

	values := []float64{
		6.209665325776139e-03,
		1.222447265504470e-02,
		2.275013194817922e-02,
		4.005915686381709e-02,
		6.680720126885809e-02,
		1.056497736668553e-01,
		1.586552539314571e-01,
		2.266273523768682e-01,
		3.085375387259869e-01,
		4.012936743170763e-01,
		5.000000000000000e-01,
		5.987063256829237e-01,
		6.914624612740131e-01,
		7.733726476231317e-01,
		8.413447460685429e-01,
		8.943502263331446e-01,
		9.331927987311419e-01,
	}

	assert.AlmostEqual(New(1, 2*2).CDF(points), values, t)
}
