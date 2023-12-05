package day5

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
)

type syncInt struct {
	mu sync.RWMutex

	val int
}

func (s *syncInt) compareAndSave(newVal int) {
	s.mu.Lock()

	defer s.mu.Unlock()

	if newVal < s.val {
		s.val = newVal
	}
}

type destinationSourceMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLen              int
}

type seedRange struct {
	start, rangeLen int
}

func ResolvePartOne(r io.Reader) int {
	s := bufio.NewScanner(r)
	seeds := []int{}

	// parse seeds
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}

		for _, rawSeed := range strings.Split(line[7:], " ") {
			seed, _ := strconv.Atoi(rawSeed)
			seeds = append(seeds, seed)
		}
	}

	// shift = value - sourceRangeStart
	// dest = destRangeStart + shift
	return findMinLocation(seeds, parseDestinationSourceMaps(s))
}

// 7873085 is to high
func ResolvePartTwo(r io.Reader) int {
	s := bufio.NewScanner(r)
	seeds := []int{}

	// parse seeds
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}

		for _, rawSeed := range strings.Split(line[7:], " ") {
			seed, _ := strconv.Atoi(rawSeed)
			seeds = append(seeds, seed)
		}
	}

	seedRanges := make([]seedRange, 0, len(seeds)/2)
	for idx, seed := range seeds {
		if (idx+1)%2 == 0 {
			seedRanges = append(seedRanges, seedRange{
				rangeLen: seed,
				start:    seeds[idx-1],
			})
		}
	}

	destinationSourceMaps := parseDestinationSourceMaps(s)

	var (
		minLocation syncInt
		wg          sync.WaitGroup
	)

	minLocation.val = math.MaxInt64

	wg.Add(len(seedRanges))
	for _, sr := range seedRanges {
		go func(sr seedRange) {
			defer wg.Done()

			for seed := sr.start; seed < sr.start+sr.rangeLen; seed++ {
				minLocation.compareAndSave(findMinLocationForSeed(seed, destinationSourceMaps))
			}

		}(sr)
	}

	fmt.Println("spawned all goroutines")
	wg.Wait()

	return minLocation.val
}

func findMinLocation(seeds []int, destinationSourceMaps [][]destinationSourceMap) int {
	minLocation := math.MaxInt64
	for _, seed := range seeds {
		val := findMinLocationForSeed(seed, destinationSourceMaps)
		if val < minLocation {
			minLocation = val
		}
	}

	return minLocation
}

func findMinLocationForSeed(seed int, destinationSourceMaps [][]destinationSourceMap) int {
	val := seed
	for _, dsm := range destinationSourceMaps {
		for _, ds := range dsm {
			if val >= ds.sourceRangeStart && val <= ds.sourceRangeStart+ds.rangeLen {
				val = ds.destinationRangeStart + val - ds.sourceRangeStart
				break
			}
		}
	}

	return val
}

func parseDestinationSourceMaps(s *bufio.Scanner) [][]destinationSourceMap {
	result := [][]destinationSourceMap{}
	for s.Scan() {
		dsms := []destinationSourceMap{}
		for s.Scan() {
			if s.Text() == "" {
				break
			}

			rawMap := strings.Split(s.Text(), " ")
			drs, _ := strconv.Atoi(rawMap[0])
			srs, _ := strconv.Atoi(rawMap[1])
			rl, _ := strconv.Atoi(rawMap[2])

			dsms = append(dsms,
				destinationSourceMap{
					destinationRangeStart: drs,
					sourceRangeStart:      srs,
					rangeLen:              rl,
				})
		}

		result = append(result, dsms)
	}

	return result
}
