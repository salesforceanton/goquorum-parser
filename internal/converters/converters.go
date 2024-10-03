package converters

import (
	"errors"
	"math/big"
)

func StringToBigInt(str string) (*big.Int, error) {
	bigval := new(big.Int)
	bigval, ok := bigval.SetString(str, 10)
	if !ok {
		return nil, errors.New("invalid conversion")
	}
	return bigval, nil
}
