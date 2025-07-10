package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func testHitungVolume(t *testing.T) {
	var hasil = 2 * 3
	assert.Equal(t, hasil, 6, "Perhitungan volume salah")
}
// func testHitungLuas(t *testing.T) {
// 	assert.Equal(t, kubus.Luas(), luasSeharusnya, "Perhitungan luas salah")
// }
// func testHitungKeliling(t *testing.T) {
// 	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "Perhitungan keliling salah")
// }

