// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// Binary encoding

package usbip

import "github.com/OpenPrinting/go-mfp/internal/assert"

// Encoder helps to encode USB binary data
type encoder struct {
	buf []byte
}

// NewEncoder creates a new encoder.
//
// The supplied parameter is the size hint. It defines capacity
// of the preallocated buffer.
func newEncoder(sz int) *encoder {
	return &encoder{make([]byte, 0, sz)}
}

// PutU8 writes a single byte to the encoder.
func (enc *encoder) PutU8(v byte) {
	enc.buf = append(enc.buf, v)
}

// PutBytes writes multiple bytes to the encoder.
func (enc *encoder) PutBytes(v ...byte) {
	enc.buf = append(enc.buf, v...)
}

// PutLE16 writes a 16-bit little endian word to the encoder.
func (enc *encoder) PutLE16(v uint16) {
	enc.PutBytes(byte(v), byte(v>>8))
}

// PutBE16 writes a 16-bit big endian word to the encoder.
func (enc *encoder) PutBE16(v uint16) {
	enc.PutBytes(byte(v>>8), byte(v))
}

// PutLE32 writes a 32-bit little endian word to the encoder.
func (enc *encoder) PutLE32(v uint32) {
	enc.PutBytes(byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

// PutBE32 writes a 32-bit big endian word to the encoder.
func (enc *encoder) PutBE32(v uint32) {
	enc.PutBytes(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// Bytes returns the encoder's internal buffer with the encoded data.
func (enc *encoder) Bytes() []byte {
	return enc.buf
}

// Decoder helps to decode USB binary data.
type decoder struct {
	buf []byte
}

// newDecoder creates a new decoder.
func newDecoder(data []byte) *decoder {
	return &decoder{data}
}

// Len returns count of bytes available in the decoder.
func (dec *decoder) Len() int {
	return len(dec.buf)
}

// GetU8 reads a single byte from the decoder.
func (dec *decoder) GetU8() byte {
	v := dec.buf[0]
	dec.buf = dec.buf[1:]
	return v
}

// GetLE16 reads a 16-bit little endian word from the decoder.
func (dec *decoder) GetLE16() uint16 {
	v0 := dec.buf[1]
	v1 := dec.buf[0]
	dec.buf = dec.buf[2:]
	return (uint16(v0) << 8) | uint16(v1)
}

// GetBE16 reads a 16-bit big endian word from the decoder.
func (dec *decoder) GetBE16() uint16 {
	v0 := dec.buf[0]
	v1 := dec.buf[1]
	dec.buf = dec.buf[2:]
	return (uint16(v0) << 8) | uint16(v1)
}

// GetLE32 reads a 32-bit little endian word from the decoder.
func (dec *decoder) GetLE32() uint32 {
	v0 := dec.buf[3]
	v1 := dec.buf[2]
	v2 := dec.buf[1]
	v3 := dec.buf[0]
	dec.buf = dec.buf[4:]
	return (uint32(v0) << 24) |
		(uint32(v1) << 16) |
		(uint32(v2) << 8) |
		uint32(v3)
}

// GetBE32 reads a 32-bit big endian word from the decoder.
func (dec *decoder) GetBE32() uint32 {
	v0 := dec.buf[0]
	v1 := dec.buf[1]
	v2 := dec.buf[2]
	v3 := dec.buf[3]
	dec.buf = dec.buf[4:]
	return (uint32(v0) << 24) |
		(uint32(v1) << 16) |
		(uint32(v2) << 8) |
		uint32(v3)
}

// GetData reads fixed amoubt bytes from the decoder.
func (dec *decoder) GetData(data []byte) {
	assert.Must(len(dec.buf) >= len(data))
	copy(data, dec.buf)
	dec.buf = dec.buf[len(data):]
}
