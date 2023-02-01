package services

import (
	"log"
	"strconv"
	"strings"
)

type Service struct {
	Dct     map[string]int
	Data    [][]string
	inWork  bool
	workers int
}

func (sr *Service) InitService(dct map[string]int, data [][]string) *Service {
	sr.Dct = dct
	sr.Data = data
	sr.inWork = false
	sr.workers = 0
	return sr
}

func (sr *Service) solve(equation string) int {
	if sr.inWork {
		sr.workers += 1
	}
	if sr.workers > 100 {
		log.Panicln("Infinite Recursion found, stop!")
	}
	equation = strings.TrimPrefix(equation, "=")
	equation = strings.Trim(equation, " ")
	if equation == "" {
		return 0
	}
	if strings.Contains(equation, "/") {
		parts := strings.Split(equation, "/")
		left, errL := strconv.Atoi(parts[0])
		right, errR := strconv.Atoi(parts[1])
		if errL == nil && errR == nil {
			if right == 0 {
				log.Panicln("Division by zero found at: ", equation)
			}
			return left / right
		} else if errL != nil && errR == nil {
			return sr.solve(parts[0]) / right
		} else if errL == nil && errR != nil {
			return sr.solve(parts[1]) / left
		} else {
			rightSide := sr.solve(parts[1])
			if rightSide == 0 {
				log.Panicln("Division by zero found at: ", equation)
			}
			return sr.solve(parts[0]) / rightSide
		}
	} else if strings.Contains(equation, "+") {
		parts := strings.Split(equation, "+")
		left, errL := strconv.Atoi(parts[0])
		right, errR := strconv.Atoi(parts[1])
		if errL == nil && errR == nil {
			return left + right
		} else if errL != nil && errR == nil {
			return sr.solve(parts[0]) + right
		} else if errL == nil && errR != nil {
			return sr.solve(parts[1]) + left
		} else {
			return sr.solve(parts[0]) + sr.solve(parts[1])
		}
	} else if strings.Contains(equation, "-") {
		parts := strings.Split(equation, "-")
		left, errL := strconv.Atoi(parts[0])
		right, errR := strconv.Atoi(parts[1])
		if errL == nil && errR == nil {
			return left - right
		} else if errL != nil && errR == nil {
			return sr.solve(parts[0]) - right
		} else if errL == nil && errR != nil {
			return sr.solve(parts[1]) - left
		} else {
			return sr.solve(parts[0]) - sr.solve(parts[1])
		}
	} else if strings.Contains(equation, "*") {
		parts := strings.Split(equation, "*")
		left, errL := strconv.Atoi(parts[0])
		right, errR := strconv.Atoi(parts[1])
		if errL == nil && errR == nil {
			return left * right
		} else if errL != nil && errR == nil {
			return sr.solve(parts[0]) * right
		} else if errL == nil && errR != nil {
			return sr.solve(parts[1]) * left
		} else {
			return sr.solve(parts[0]) * sr.solve(parts[1])
		}

	}
	// Here we look at the cell value
	parseCell, errParsing := strconv.Atoi(equation)
	if errParsing != nil {
		myString := ""
		// checking if cell is valid
		for _, v := range equation {
			if strings.ToLower(string(v)) >= "a" && strings.ToLower(string(v)) <= "z" {
				myString += string(v)
			}
		}

		_, ok := sr.Dct[myString]
		if ok {
			number := ""
			for _, v := range equation {
				if string(v) >= "0" && string(v) <= "9" {
					number += string(v)
				}
			}
			parseNumber, errNumber := strconv.Atoi(number)
			if errNumber != nil {
				log.Panicln("Invalid cell found: ", equation)
			}
			if parseNumber <= 0 || parseNumber >= len(sr.Data) {
				log.Panicln("Invalid index found at: ", equation)
			}
			return sr.solve(sr.Data[parseNumber][sr.Dct[myString]])
		} else {
			log.Panicln("Invalid cell found: ", equation)
		}
	}
	return parseCell
}

func (sr *Service) Iterator() *Service {
	for i, v := range sr.Data {
		for j, el := range v {
			if strings.Contains(el, "=") {
				sr.inWork = true
				sr.Data[i][j] = strconv.Itoa(sr.solve(sr.Data[i][j]))
				sr.inWork = false
				sr.workers = 0
			}
		}
	}
	return sr
}
