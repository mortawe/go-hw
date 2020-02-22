package sort

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestNum = 9

func TestOK(t *testing.T) {

	flags := []Flags{{
		IsFold:     false,
		IsUnique:   false,
		IsReversed: false,
		OutputFile: "test.o",
		IsNumbers:  false,
		isColomn:   1,
	}, {
		IsFold:     false,
		IsUnique:   false,
		IsReversed: false,
		OutputFile: "test.o",
		IsNumbers:  false,
		isColomn:   1,
	}, {
		IsFold:     true,
		IsUnique:   false,
		IsReversed: false,
		OutputFile: "test.o",
		IsNumbers:  false,
		isColomn:   1,
	}, {
		IsFold:     false,
		IsUnique:   false,
		IsReversed: true,
		OutputFile: "test.o",
		IsNumbers:  false,
		isColomn:   1,
	}, {
		IsFold:     false,
		IsUnique:   false,
		IsReversed: false,
		OutputFile: "test.o",
		IsNumbers:  true,
		isColomn:   1,
	}, {
		IsFold:     false,
		IsUnique:   false,
		IsReversed: true,
		OutputFile: "test.o",
		IsNumbers:  true,
		isColomn:   1,
	}, {
		IsFold:     false,
		IsUnique:   false,
		IsReversed: false,
		OutputFile: "test.o",
		IsNumbers:  false,
		isColomn:   2,
	}, {
		IsFold:     false,
		IsUnique:   true,
		IsReversed: false,
		OutputFile: "test.o",
		IsNumbers:  false,
		isColomn:   1,
	}, {
		IsFold:     false,
		IsUnique:   false,
		IsReversed: false,
		OutputFile: "test.o",
		IsNumbers:  true,
		isColomn:   1,
	},
	}

	for i := 1; i < TestNum; i++ {
		lines, err := ReadLines("tests/" + strconv.Itoa(i) + ".in")
		if err != nil {
			log.Println(err)
		}

		lines, err = SortWithFlags(lines, flags[i])
		if err != nil {
			log.Println(err)
		}
		expectLines, err := ReadLines("tests/" + strconv.Itoa(i) + ".a")
		assert.Equal(t, expectLines, lines)
	}
}
