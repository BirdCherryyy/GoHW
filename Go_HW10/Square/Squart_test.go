package square

//testify & Теабличные тесты

//Теабличные тесты
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type squareStr struct {
	A           int
	B           int
	Res         int
	description string
}

func TestSquare(t *testing.T) {
	tests := []squareStr{
		{A: 1, B: 1, Res: 1, description: "for A = 1 and B = 1 Res = 1"},
		{A: 2, B: 1, Res: 2, description: "for A = 2(1) and B = 1(2) Res = 2"},
		{A: 3, B: 2, Res: 6, description: "for A = 3(2) and B = 2(3) Res = 6"},
	}
	for _, i := range tests {
		result := square(i.A, i.B)
		if result != i.Res {
			t.Errorf(i.description, "but not =", result)
		}
	}
}

//testify
func TestSquare2(t *testing.T) {
	tests := []squareStr{
		{A: 1, B: 1, Res: 1, description: "for A = 1 and B = 1 Res = 1"},
		{A: 2, B: 1, Res: 2, description: "for A = 2(1) and B = 1(2) Res = 2"},
		{A: 3, B: 2, Res: 6, description: "for A = 3(2) and B = 2(3) Res = 6"},
	}
	for _, i := range tests {
		assert.Equal(t, square(i.A, i.B), i.Res, i.description)
	}
}
