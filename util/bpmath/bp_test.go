package bpmath

import (
	"fmt"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestBPToDec(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name string
		a    FixedBP
		exp  math.LegacyDec
	}{
		{"t1", 99999, math.LegacyMustNewDecFromStr("9.9999")},
		{"t2", One * 10, math.LegacyMustNewDecFromStr("10.0")},
	}
	require := require.New(t)
	for _, tc := range tcs {
		bp := BP(tc.a).ToDec()
		require.Equal(tc.exp.String(), bp.String(), fmt.Sprint("test-bp ", tc.name))
	}
}

// Tests if it works both with sdk.Int and math.Int
func TestInt(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	bp := BP(100)
	si := sdk.NewInt(1234)

	var sresult sdk.Int = Mul(si, bp)
	require.Equal(sresult, sdk.NewInt(12))
	var mresult math.Int = Mul(si, bp)
	require.Equal(mresult, math.NewInt(12))

	// now let's check math.Int
	mi := math.NewInt(1234)
	sresult = Mul(mi, bp)
	require.Equal(sresult, sdk.NewInt(12))
	mresult = Mul(mi, bp)
	require.Equal(mresult, math.NewInt(12))

	// test rounding
	si = sdk.NewInt(1299)
	require.Equal(bp.Mul(si), sdk.NewInt(12))

	si = sdk.NewInt(-1299)
	require.Equal(bp.Mul(si), sdk.NewInt(-12))

	si = sdk.NewInt(-1201)
	require.Equal(bp.Mul(si), sdk.NewInt(-12))
}

func TestBPMulDec(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	bp := BP(1000)
	bp2 := BP(1)
	bp3 := BP(5000)
	bp4 := BP(20000)
	d := sdk.MustNewDecFromStr("12.5002")
	d2 := sdk.NewDec(10000)
	d3 := sdk.NewDec(1000)

	require.Equal(d, MulDec(d, One))
	require.Equal(sdk.ZeroDec(), MulDec(d, Zero))
	require.Equal(sdk.OneDec(), bp2.MulDec(d2))
	require.Equal(sdk.MustNewDecFromStr("0.1"), bp2.MulDec(d3))

	require.Equal(sdk.MustNewDecFromStr("1.25002"), bp.MulDec(d))
	require.Equal(sdk.MustNewDecFromStr("0.00125002"), bp2.MulDec(d))
	require.Equal(sdk.MustNewDecFromStr("6.2501"), bp3.MulDec(d))
	require.Equal(sdk.MustNewDecFromStr("25.0004"), bp4.MulDec(d))
}

func TestBPEqual(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	var b1 BP = 1
	var b2 BP = 1
	var b3 BP = 10
	require.True(b1.Equal(b2))
	require.True(b2.Equal(b2))
	require.False(b1.Equal(b3))
}
