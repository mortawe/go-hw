package sort

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type Flags struct {
	fPtr bool
	uPtr bool
	rPtr bool
	oPtr string
	nPtr bool
	kPtr int
}

func writeFile(lines [] string, flags Flags) error {

	file, err := os.Create(flags.oPtr)
	defer file.Close()
	if err != nil {
		return err
	}
	for i := 0; i < len(lines)-1; i++ {
		_, err := file.Write([]byte(lines[i] + "\n"))
		if err != nil {
			return err
		}
	}
	if len(lines) > 0 {
		_, err := file.Write([]byte(lines[len(lines)-1]))
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteResult(lines [] string, flags Flags) error {
	if flags.oPtr != "" {
		return writeFile(lines, flags)
	}
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

	return nil
}

func calcMaskWithFlags(st string, flags Flags) (string, error) {
	if flags.kPtr > 1 {
		for strings.Contains(st, "  ") {
			st = strings.ReplaceAll(st, "  ", " ")
		}
		temp := strings.Split(st, " ")
		if len(temp) < flags.kPtr {
			return "", errors.New("количество столбцов меньше чем значение аргумента K") // отправляем вниз левое, если у него столбцов меньше, чем надо
		} else {
			st = temp[flags.kPtr-1] // иначе рассматриваем интересующий столбец
		}
	}
	if flags.fPtr {
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

func SortWithFlags(lines[] string, flags Flags) ([]string, error) {
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

		if flags.nPtr {
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

		if flags.rPtr {
			return !compRes
		} else {
			return  compRes
		}

	})

	if flags.uPtr {
		lines = unify(lines, mask)
	}
	return lines, nil
}



func InitFlags(flags Flags)  Flags{

	flag.BoolVar(&flags.fPtr, "f", false, "ignore register")
	flag.BoolVar(&flags.uPtr,"u", false, "shows first of equals")
	flag.BoolVar(&flags.rPtr,"r", false, "reverse")
	flag.StringVar(&flags.oPtr, "o", "", "output file")
	flag.BoolVar(&flags.nPtr, "n", false, "number sort")
	flag.IntVar(&flags.kPtr, "k", 1, "column number")

	flag.Parse()
	return flags
}