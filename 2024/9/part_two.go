package main

import (
	"io"
)

type block struct {
	id        int
	alocBytes int
	freeBytes int

	next *block
	prev *block
	tail *block
}

func solvePartTwo(r io.Reader) (int, error) {
	b, _ := io.ReadAll(r)
	input := string(b)

	var blocks *block = nil

	id := 0
	for i := 0; i < len(input); i += 2 {
		blockSize := int(input[i] - '0')
		freeSize := 0
		if i+1 < len(input) {
			freeSize = int(input[i+1] - '0')
		}

		current := &block{
			id:        id,
			alocBytes: blockSize,
			freeBytes: freeSize,
			next:      blocks,
		}

		if blocks != nil {
			blocks.prev = current
			current.tail = blocks.tail
		}

		blocks = current
		if blocks.tail == nil {
			blocks.tail = current
		}

		id++
	}

	h := blocks
	for h != nil {
		next := h.next
		for t := blocks.tail; t != nil && t != h; t = t.prev {
			if t.id == h.id {
				break
			}

			if t.freeBytes < h.alocBytes {
				continue
			}

			c := h
			if c.prev != nil {
				c.prev.next = c.next
			}

			if c.next != nil {
				c.next.prev = c.prev
				c.next.freeBytes += c.alocBytes + c.freeBytes
			}

			c.prev = t.prev
			if t.prev != nil {
				t.prev.next = c
			}

			c.next, t.prev = t, c
			c.freeBytes = t.freeBytes - c.alocBytes
			t.freeBytes = 0
		}
		h = next
	}

	result := 0
	offset := 0
	for t := blocks.tail; t != nil; t = t.prev {
		for i := range t.alocBytes {
			result += t.id * (i + offset)
		}

		offset += t.alocBytes + t.freeBytes
	}

	return result, nil
}
