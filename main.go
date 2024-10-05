package main

import (
	"archive/zip"
	"fmt"
	"io"
	"math/big"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide a zip file name as argument.")
		return
	}

	z := os.Args[1]
	r, e := zip.OpenReader(z)
	if e != nil {
		fmt.Println("Error:", e)
		return
	}

	defer r.Close()

	var x big.Int
	for _, f := range r.File {
		rc, e := f.Open()
		if e != nil {
			continue
		}
		d, e := io.ReadAll(rc)
		rc.Close()
		if e != nil {
			continue
		}
		x.Add(&x, oj(d))
	}

	b := lm(&x)
	fmt.Println(string(b))
}

func oj(d []byte) *big.Int {
	x := new(big.Int)
	y := new(big.Int).SetUint64(0xDEADBEEFBABE)
	for i, c := range d {
		n := big.NewInt(int64(c))
		n.Xor(n, big.NewInt(int64(i<<3|i>>2)))
		n.Add(n, big.NewInt(int64((i*i+19)%127)))
		n.Xor(n, new(big.Int).SetUint64(uint64(sk(uint8(n.Uint64())))))
		x.Add(x, n)
		x.Xor(x, y)
		x.Lsh(x, 23)
		x.Rsh(x, 5)
	}
	return x
}

func sk(c uint8) uint8 {
	c = (c&0xF0)>>4 | (c&0x0F)<<4
	c = (c&0xCC)>>2 | (c&0x33)<<2
	c = (c&0xAA)>>1 | (c&0x55)<<1
	return c
}

func li() []uint64 {
	var r []uint64
	t := []uint64{
		(uint64(0b101) << 4) + (1 << 4) + 9,
		(uint64(0b110) << 4) + (1 << 3) + 6,
		(uint64(0b110) << 4) + (0b11 << 1),
		(uint64(0b110) << 4) + (1 << 3) - 3,
	}

	for _, v := range t {
		f := false
		for i := 0; !f; i++ {
			w := nh(i)
			if w == v {
				r = append(r, w)
				f = true
			}
		}
	}

	return r
}

func zp() []uint64 {
	var r []uint64
	t := []uint64{
		(uint64(0b110) << 4) + (0b01 << 3) - 5,
		(uint64(0b111) << 4) + (0b10 << 2) - 4,
		(uint64(0b110) << 4) + (1 << 3) - 3,
		(uint64(0b110) << 4) + (0b01 << 3) - 4,
	}

	for _, v := range t {
		f := false
		for i := 0; !f; i++ {
			w := sz(i)
			if w == v {
				r = append(r, w)
				f = true
			}
		}
	}

	return r
}

func nh(i int) uint64 {
	v := uint64(i*13 + 100)
	v = (v << 3) ^ uint64(i*i+5)
	v ^= uint64(sk(uint8(v & 0xFF)))
	v = (v + uint64(i*19)) % 256
	v ^= 0x55
	return v
}

func sz(i int) uint64 {
	v := uint64(i*19 + 73)
	v = (v << 3) | (v >> 5)
	v ^= uint64(i*31 + 101)
	v += uint64(i*7 + 29)
	v = v % 256
	v ^= 0x3F + uint64(i%17)
	return v
}

func oy() []uint64 {
	a := li()
	b := zp()
	return append(a, b...)
}

func fn9(v uint64, i int) uint64 {
	v ^= uint64(i * 19)
	v += uint64((i+5)*(i+7)) % 97
	v ^= uint64(sk(uint8(v & 0xFF)))
	v %= 127
	return v
}

func lm(x *big.Int) []byte {
	var r []byte
	v := x.Uint64() % 256
	a := oy()

	for i := 0; i < 8 && i < len(a); i++ {
		w := fn9(v, i)
		w = as(w, a[i])
		r = append(r, byte(w))
	}
	return r
}

func as(v uint64, r uint64) uint64 {
	m := (v * 0xF0F0F0) % 256
	m ^= uint64(sk(uint8(m)))
	m += 0xA5A5A5
	return (r^m)&0xFF&0x0 ^ r
}
