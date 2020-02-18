package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func splitTokens(expr string) ([]string,error) {
	var tokens [] string
	curNum := ""
	bracketsNum := 0
	for r := range expr {
		if expr[r] == ' ' {
			continue
		}
		if expr[r] == '(' || expr[r] == ')'{
			switch expr[r] {
			case '(' : bracketsNum++
			case ')' : bracketsNum--
			}
			if bracketsNum < 0 {
				return nil, errors.New("wrong brackets sequence")
			}
			if curNum != ""  {
				tokens = append(tokens, curNum)
				curNum = ""
			}

			if len(tokens) >= 2 && tokens[len(tokens)-1] == "-" && tokens[len(tokens)- 2] == "("{
				tokens = append(tokens[:len(tokens) - 1], append([]string{"0"}, tokens[len(tokens) - 1:]...)...)
			}
			if len(tokens) == 1 && tokens[0] == "-" {
				tokens = append( []string{"0"}, tokens...)
			}

			tokens = append(tokens, string(expr[r]))
			continue
		}
		if  _, err := strconv.Atoi(string(expr[r])); err == nil {
			okCond := len(tokens) < 2
			if len(tokens) >= 2 {
				_, err := strconv.ParseFloat(tokens[len(tokens) - 2], 64)
				if err == nil {
					okCond = false
				} else {
					if tokens[len(tokens) - 2] == ")" {
						okCond = false
					} else {
						okCond = true
					}
				}
			}
			if curNum == "" && okCond {
				if len(tokens) > 0 &&
					( tokens[len(tokens) - 1] == "-" || tokens[len(tokens) - 1] == "+" ){
					curNum += tokens[len(tokens) - 1]
					tokens = tokens[:len(tokens) - 1]
				}
			}
			curNum += string(expr[r])
			continue
		}
		if expr[r] == '.' {
			if curNum != "" && !strings.Contains(curNum, ".") {
				curNum += string(expr[r])
				continue
			}
			return nil, errors.New("wrong input: double dot in number")
		}
		if expr[r] == '+' || expr[r] == '-' || expr[r] == '/' || expr[r] == '*' {
			if r >= 1 &&
				(expr[r-1] == '+' || expr[r-1] == '-' ||
				expr[r-1] == '/' || expr[r-1] == '*') {
				return nil, errors.New("wrong input: double sign")
			}
			if curNum != "" {
				tokens = append(tokens, curNum)
				curNum = ""
			}
			tokens = append(tokens, string(expr[r]))
			continue
		}
		return nil, errors.New("wrong input")
	}
	if curNum != "" {
		tokens = append(tokens, curNum)
	}
	return tokens, nil
}

func calcSimpleExpr(tokens [] string) (string, error){

	if len(tokens) == 1 {
		return tokens[0], nil
	}
	if len(tokens) == 3 {
		left, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			return "", err
		}
		right, err := strconv.ParseFloat(tokens[2], 64)
		if err != nil {
			return "", err
		}
		switch tokens[1] {
		case "+":
			return fmt.Sprintf("%f", left + right), nil
		case "-":
			return fmt.Sprintf("%f", left - right), nil
		case "*":
			return fmt.Sprintf("%f", left * right), nil
		case "/":
			if right == 0.0 {
				return "", errors.New("division by zero")
			}
			return fmt.Sprintf("%f", left / right), nil
		default:
			return "", errors.New("could not calc simple expr because of strange operation")
		}
	}
	r := 0
	for len(tokens) > 2 && r < len(tokens) {
		if tokens[r] == "*" || tokens[r] == "/" {
			result, err := calcSimpleExpr(tokens[r-1 : r+2])
			if err != nil {
				return "", err
			}
			tokens = append(append(tokens[:r-1],  result), tokens[r+2:]...)
			r = r - 2
		}
		r++
	}
	r = 0
	for len(tokens) > 1 && r < len(tokens)  {
			if tokens[r] == "+" || tokens[r] == "-" {
				if r - 1 < 0 || r + 1 >= len(tokens) {
					return "", errors.New("invalid expression : + or (bin) - in wrong places")
				}
				result, err := calcSimpleExpr(tokens[r-1 : r+2])
				if err != nil {
					return "", err
				}
				tokens = append(append(tokens[:r-1], result), tokens[r+2:]...)
				r = r - 2
			}
			r++
	}

	if len(tokens) == 1 {
		return tokens[0], nil
	}
	return "", errors.New("could not calc simple expr")
}

func calcInsideBrackets(tokens [] string) (float64, error) {
	containsBrackets := true
	start, end, bracketsSum := 0, len(tokens) - 1, 0
	for containsBrackets {
		containsBrackets = false
		bracketsSum = 0
		for r:=0; r < len(tokens); r++ {
			if tokens[r] == "(" {
				if bracketsSum == 0 {
					start = r
				}
				bracketsSum++
				containsBrackets = true
				continue
			}
			if tokens[r] == ")" {
				bracketsSum--
				if bracketsSum == 0 {
					end = r
					res, err := calcInsideBrackets(tokens[start+1:end])
					if err != nil {
						return 0.0, err
					}
					tokens = append(append(tokens[:start], fmt.Sprintf("%f", res)), tokens[end+1:]...)
					r = start
					continue
				}
				continue
			}
		}
		if bracketsSum != 0 {
			return 0.0, errors.New("wrong brackets")
		}

	}


	calcResult, err := calcSimpleExpr(tokens)
	if err != nil {
		return 0.0, err
	}
	resInFloat, err := strconv.ParseFloat(calcResult, 64)
	if err != nil {
		return 0.0, err
	}
	return resInFloat, nil
	return 0.0, errors.New("can not calc inside brackets")
}

func Calculate(expr string) (float64, error){

	tokens, err := splitTokens(expr)
	if err != nil {
		return 0.0, err
	}

	res, err := calcInsideBrackets(tokens)
	if err != nil {
		return 0.0, err
	}

	return res, nil
}

