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
		fPtr: false,
		uPtr: false,
		rPtr: false,
		oPtr: "test.o",
		nPtr: false,
		kPtr: 1,
	}, {
		fPtr: false,
		uPtr: false,
		rPtr: false,
		oPtr: "test.o",
		nPtr: false,
		kPtr: 1,
	}, {
		fPtr: true,
		uPtr: false,
		rPtr: false,
		oPtr: "test.o",
		nPtr: false,
		kPtr: 1,
	}, {
		fPtr: false,
		uPtr: false,
		rPtr: true,
		oPtr: "test.o",
		nPtr: false,
		kPtr: 1,
	},  {
		fPtr: false,
		uPtr: false,
		rPtr: false,
		oPtr: "test.o",
		nPtr: true,
		kPtr: 1,
	},{
		fPtr: false,
		uPtr: false,
		rPtr: true,
		oPtr: "test.o",
		nPtr: true,
		kPtr: 1,
	},{
		fPtr: false,
		uPtr: false,
		rPtr: false,
		oPtr: "test.o",
		nPtr: false,
		kPtr: 2,
	},{
		fPtr: false,
		uPtr: true,
		rPtr: false,
		oPtr: "test.o",
		nPtr: false,
		kPtr: 1,
	},{
		fPtr: false,
		uPtr: false,
		rPtr: false,
		oPtr: "test.o",
		nPtr: true,
		kPtr: 1,
	},

	}

	for i:=1; i < TestNum; i++ {
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
