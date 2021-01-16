package bitmap

type Bitmap struct {
	bits []int64
	cnt  int
}

func New() *Bitmap {
	return &Bitmap{bits: make([]int64, 1)}
}

func NewWithLength(len int) *Bitmap {
	if len <= 64 {
		return &Bitmap{bits: make([]int64, 1)}
	}
	return &Bitmap{bits: make([]int64, (len-1)>>6+1)}
}

func (bm *Bitmap) GetBit(idx int) int {
	i := idx >> 6
	if i >= len(bm.bits) {
		return 0
	}
	j := idx & 63
	return int(bm.bits[i] >> j & 1)
}

func (bm *Bitmap) SetBit(idx, bit int) {
	if bit != 0 && bit != 1 {
		return
	}
	i := idx >> 6
	if i >= len(bm.bits) {
		newBits := make([]int64, i+1)
		copy(newBits, bm.bits)
		bm.bits = newBits
	}
	j := idx & 63
	x := int64(bit) << j
	if bm.bits[i]>>j&1 == 0 {
		bm.bits[i] |= x
		if bit == 1 {
			bm.cnt++
		}
	} else {
		bm.bits[i] &= x
		if bit == 0 {
			bm.cnt--
		}
	}
}

func (bm *Bitmap) BitCount() int {
	return bm.cnt
}

func (bm *Bitmap) Clear() {
	bm.cnt = 0
	for i := range bm.bits {
		bm.bits[i] = 0
	}
}
