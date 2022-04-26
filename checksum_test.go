package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	salt := "abcdefghijklmnopqrstuvwxyz123456"

	data := []byte(`{"v": 2,"tid":"G-T8JSQCLDWE","gtm":"2oe4p0","_p":1477146715,"_z":"ccd.NbB","_pp":1,"cid":947181994.1650784513,"ul":"ja","sr":"3413x1440","_s":1,"sid":1650976491,"sct":2,"seg":0,"dl":"https://qubena.com/","dr2":"https://www.google.com/","dt":"TOP - Qubena（株式会社COMPASS） - AI型教材","en":"page_view","_ss":"1"}`)

	if Checksum(data, salt) != "6656be9e7a80ab2d8b427732483bdf83" {
		t.Error("test")
	}
}

func BenchmarkChecksum(b *testing.B) {
	salt := "abcdefghijklmnopqrstuvwxyz123456"

	data := []byte(`{"v": 2,"tid":"G-T8JSQCLDWE","gtm":"2oe4p0","_p":1477146715,"_z":"ccd.NbB","_pp":1,"cid":947181994.1650784513,"ul":"ja","sr":"3413x1440","_s":1,"sid":1650976491,"sct":2,"seg":0,"dl":"https://qubena.com/","dr2":"https://www.google.com/","dt":"TOP - Qubena（株式会社COMPASS） - AI型教材","en":"page_view","_ss":"1"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if Checksum(data, salt) != "6656be9e7a80ab2d8b427732483bdf83" {
			b.Error("test")
		}
	}
}
