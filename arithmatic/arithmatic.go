package arithmatic

import "errors"

func div(x float64, y float64) (float64, error)  {
	if y == 0 {
		return 0, errors.New("divide by 0 not allowed")
	}

	return x/y, nil
}
