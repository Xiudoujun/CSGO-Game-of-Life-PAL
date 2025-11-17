package tests

import (
	"fmt"
	"os"
	"testing"

	"uk.ac.bris.cs/gameoflife/gol"
)

type CCA struct {
	name  string
	w, h  int
	turns int
}

// Small / Medium / Large 三种规模
var ST = []CCA{
	{"64", 64, 64, 100},
	{"256", 256, 256, 500},
	{"512", 512, 512, 1000},
}

// const benchLength = 1000
//const optimalThreads = 8

func BenchmarkGol(b *testing.B) {
	threadConfigs := []int{1, 2, 4, 8, 16}
	for _, threads := range threadConfigs {
		for _, c := range ST {
			//for threads := 1; threads <= 16; threads++ {
			os.Stdout = nil // Disable all program output apart from benchmark results
			p := gol.Params{
				Turns:   c.turns,
				Threads: threads,
				//Threads:     optimalThreads,
				ImageWidth:  c.w,
				ImageHeight: c.h,
				//Turns:       benchLength,
				//Threads:     threads,
				//ImageWidth:  512,
				//ImageHeight: 512,
			}
			name := fmt.Sprintf("%dx%dx%d-%d", p.ImageWidth, p.ImageHeight, p.Turns, p.Threads)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					events := make(chan gol.Event)
					go gol.Run(p, events, nil)
					for range events {

					}
				}
			})
		}

	}
}
