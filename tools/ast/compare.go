package ast

import (
	"strings"

	"github.com/dipdup-net/go-lib/tools/types"
)

func compareNotOptimizedTypes(x, y Default, optimizer func(string) (string, error)) (int, error) {
	if x.ValueKind != valueKindBytes {
		value, err := optimizer(x.Value.(string))
		if err != nil {
			return 0, err
		}
		x.ValueKind = valueKindBytes
		x.Value = value
	}
	if y.ValueKind != valueKindBytes {
		value, err := optimizer(y.Value.(string))
		if err != nil {
			return 0, err
		}
		y.ValueKind = valueKindBytes
		y.Value = value
	}

	return strings.Compare(x.Value.(string), y.Value.(string)), nil
}

func compareBigInt(x, y Default) int {
	xi, ok := x.Value.(*types.BigInt)
	if !ok {
		return -1
	}
	yi, ok := y.Value.(*types.BigInt)
	if !ok {
		return 1
	}
	return xi.Cmp(yi.Int)
}
