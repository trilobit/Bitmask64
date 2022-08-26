package bitmask

import "fmt"

// Bitmask64 is a struct which contains 64 bit flags from 0 to 63
type Bitmask64 struct {
	mask uint64
}

var (
	OutOfRangeErr     = fmt.Errorf("bit is out of range")
	NotInitializedErr = fmt.Errorf("bitmask is not initialized")
)

// NewBitmask64 is a constructor, argument is an initial value for bitmask
func NewBitmask64(mask uint64) Bitmask64 {
	return Bitmask64{mask: mask}
}

// SetBit is trying to set one bit (set as one)
// returns error if bitmask pointer is nil or n is out of range
func (b *Bitmask64) SetBit(n int) error {
	if err := b.checkBit(n); err != nil {
		return err
	}
	mask := uint64(1) << n
	return b.AddMask(mask)
}

// IsSetBit check the status of bit in selected position
// This method simplified and not return error if n is out of range or b is nil
func (b *Bitmask64) IsSetBit(n int) bool {
	if err := b.checkBit(n); err != nil {
		return false
	}
	mask := uint64(1) << n
	return b.mask&mask != 0
}

// UnsetBit is trying to remove one bit (set as zero)
// returns error if bitmask pointer is nil or n is out of range
func (b *Bitmask64) UnsetBit(n int) error {
	if err := b.checkBit(n); err != nil {
		return err
	}
	mask := uint64(1) << n
	b.mask &= ^mask
	return nil
}

// AddMask is trying to set multiple bits by OR operation
// returns error if bitmask pointer is nil
func (b *Bitmask64) AddMask(mask uint64) error {
	if b == nil {
		return NotInitializedErr
	}
	b.mask |= mask
	return nil
}

// SetMask replaces current mask to new
// returns error if bitmask pointer is nil
func (b *Bitmask64) SetMask(mask uint64) error {
	if b == nil {
		return NotInitializedErr
	}
	b.mask = mask
	return nil
}

// GetMask returns current mask as a uint64 number
// returns 0 if bitmask pointer is nil
func (b *Bitmask64) GetMask() uint64 {
	if b == nil {
		return 0
	}
	return b.mask
}

// checkBit is an internal method for validation
func (b *Bitmask64) checkBit(n int) error {
	if b == nil {
		return NotInitializedErr
	}
	if n < 0 || n > 63 {
		return OutOfRangeErr
	}
	return nil
}
