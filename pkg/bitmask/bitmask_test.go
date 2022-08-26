package bitmask

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitmask(t *testing.T) {
	t.Run("empty bitmask", func(t *testing.T) {
		mask := NewBitmask64(0)
		require.Zero(t, mask.GetMask())
		require.False(t, mask.IsSetBit(1))
	})
	t.Run("not initialized bitmask", func(t *testing.T) {
		var mask *Bitmask64
		require.Zero(t, mask.GetMask())
		require.ErrorIs(t, mask.SetMask(255), NotInitializedErr)
		require.ErrorIs(t, mask.AddMask(255), NotInitializedErr)
		require.ErrorIs(t, mask.SetBit(1), NotInitializedErr)
		require.False(t, mask.IsSetBit(1))
		require.ErrorIs(t, mask.UnsetBit(1), NotInitializedErr)
	})
	t.Run("bit is too big", func(t *testing.T) {
		mask := NewBitmask64(0)
		err := mask.SetBit(64)
		require.ErrorIs(t, err, OutOfRangeErr)

		err = mask.UnsetBit(64)
		require.ErrorIs(t, err, OutOfRangeErr)

		require.False(t, mask.IsSetBit(64))
	})
	t.Run("bit can not be negative", func(t *testing.T) {
		mask := NewBitmask64(0)
		err := mask.SetBit(-1)
		require.ErrorIs(t, err, OutOfRangeErr)
	})
	t.Run("check set mask", func(t *testing.T) {
		mask := NewBitmask64(0)
		err := mask.SetMask(9)
		require.Nil(t, err)
		require.True(t, mask.IsSetBit(0))
		require.False(t, mask.IsSetBit(1))
		require.False(t, mask.IsSetBit(2))
		require.True(t, mask.IsSetBit(3))
		require.Equal(t, mask.GetMask(), uint64(9))
	})
	t.Run("check add mask", func(t *testing.T) {
		mask := NewBitmask64(0)
		err := mask.SetMask(9)
		require.Nil(t, err)
		err = mask.AddMask(10)
		require.Nil(t, err)
		require.Equal(t, mask.GetMask(), uint64(11))
		err = mask.AddMask(4)
		require.Nil(t, err)
		require.Equal(t, mask.GetMask(), uint64(15))
	})
	t.Run("test bits", func(t *testing.T) {
		mask := NewBitmask64(0)

		err := mask.SetBit(0)
		require.Nil(t, err)
		require.True(t, mask.IsSetBit(0))

		err = mask.SetBit(8)
		require.Nil(t, err)
		require.True(t, mask.IsSetBit(8))

		err = mask.SetBit(63)
		require.Nil(t, err)
		require.True(t, mask.IsSetBit(63))

		err = mask.UnsetBit(63)
		require.Nil(t, err)
		require.False(t, mask.IsSetBit(63))
	})
}
