package day12

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type cacheKey struct {
	rP, cP, defCnt int
}

type record struct {
	row        string
	conditions []int
}

func ResolvePartOne(r io.Reader) int {
	records := parseRecords(r)
	result := 0
	cache := make(map[cacheKey]int)
	for _, rec := range records {
		clear(cache)
		result += calculateRecordArrangements(rec, 0, 0, 0, cache)
	}

	return result
}

func ResolvePartTwo(r io.Reader) int {
	records := parseRecords(r)
	cache := make(map[cacheKey]int)
	result := 0

	for _, rec := range records {
		clear(cache)
		result += calculateRecordArrangements(
			unfoldRecord(rec), 0, 0, 0, make(map[cacheKey]int),
		)
	}

	return result
}

func unfoldRecord(rec record) record {
	newRow := ""
	cLen := len(rec.conditions)
	newConditions := make([]int, cLen*5)
	suffix := "?"
	for i := 0; i < 5; i++ {
		if i == 4 {
			suffix = ""
		}
		newRow += rec.row + suffix
		copy(newConditions[i*cLen:i*cLen+cLen], rec.conditions)
	}

	return record{row: newRow, conditions: newConditions}
}

func calculateRecordArrangements(
	rec record, rP, cP, defCnt int, cache map[cacheKey]int,
) int {
	ck := cacheKey{
		rP, cP, defCnt,
	}
	if val, ok := cache[ck]; ok {
		return val
	}

	if rP == len(rec.row) {
		if (cP == len(rec.conditions) && defCnt == 0) ||
			(cP == len(rec.conditions)-1 && rec.conditions[cP] == defCnt) {
			return 1
		}

		return 0
	}

	res := 0
	for _, ch := range []byte{'.', '#'} {
		if rec.row[rP] == ch || rec.row[rP] == '?' {
			if ch == '.' && defCnt == 0 {
				res += calculateRecordArrangements(
					rec, rP+1, cP, 0, cache,
				)
			} else if ch == '.' && defCnt > 0 && cP < len(rec.conditions) && rec.conditions[cP] == defCnt {
				res += calculateRecordArrangements(
					rec, rP+1, cP+1, 0, cache,
				)
			} else if ch == '#' {
				res += calculateRecordArrangements(
					rec, rP+1, cP, defCnt+1, cache,
				)
			}
		}
	}

	cache[ck] = res
	return res
}

func parseRecords(r io.Reader) []record {
	var records []record

	s := bufio.NewScanner(r)
	for s.Scan() {
		ts := strings.Split(s.Text(), " ")
		row := ts[0]
		rawConditions := strings.Split(ts[1], ",")
		conditions := make([]int, len(rawConditions))
		for idx, rawCondition := range rawConditions {
			conditions[idx], _ = strconv.Atoi(rawCondition)
		}

		records = append(records, record{
			row:        row,
			conditions: conditions,
		})
	}

	return records
}
