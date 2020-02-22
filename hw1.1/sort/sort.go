package sort

import (
	"errors"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	IsFold     bool
	IsUnique   bool
	IsReversed bool
	OutputFile string
	IsNumbers  bool
	ColumnNum  int
}

func calcMaskWithFlags(st string, flags Flags) (string, error) {
	if flags.ColumnNum > 1 {
		for strings.Contains(st, "  ") {
			st = strings.ReplaceAll(st, "  ", " ")
		}
		temp := strings.Split(st, " ")
		if len(temp) < flags.ColumnNum {
			return "", errors.New("количество столбцов меньше чем значение аргумента K") // отправляем вниз левое, если у него столбцов меньше, чем надо
		} else {
			st = temp[flags.ColumnNum-1] // иначе рассматриваем интересующий столбец
		}
	}
	if flags.IsFold {
		st = strings.ToLower(st)
	}
	return st, nil
}

func unify(lines []string, mask map[string]string) []string {
	result := make([]string, len(lines))
	count := 0
	if len(lines) > 0 {
		result[0] = lines[0]
		count++
	}
	for i := 1; i < len(lines); i++ {
		if mask[lines[i]] != mask[lines[i-1]] {

			result[count] = lines[i]
			count++

		}
	}

	return result[:count]
}

func SortWithFlags(lines []string, flags Flags) (string, error) {
	mask := make(map[string]string)

	sort.Slice(lines, func(i, j int) bool {

		if mask[lines[i]] == "" {
			temp, err := calcMaskWithFlags(lines[i], flags)
			if err != nil {
				log.Println(err)
				return true
			}
			mask[lines[i]] = temp
		}

		if mask[lines[j]] == "" {
			temp, err := calcMaskWithFlags(lines[j], flags)
			if err != nil {
				log.Println(err)
				return true
			}
			mask[lines[j]] = temp
		}
		compRes := mask[lines[i]] < mask[lines[j]]

		if flags.IsNumbers {
			left, err1 := strconv.Atoi(mask[lines[i]])
			if err1 != nil {
				compRes = false
			}
			right, err2 := strconv.Atoi(mask[lines[j]])
			if err2 != nil {
				compRes = true
			}
			if err1 == nil && err2 == nil {
				compRes = left < right
			}
		}

		if flags.IsReversed {
			return !compRes
		} else {
			return compRes
		}

	})

	if flags.IsUnique {
		lines = unify(lines, mask)
	}
	return strings.Join(lines, "\n"), nil
}
