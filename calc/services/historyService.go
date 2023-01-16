package services

import "main.go/entities"

var memoryCalc []entities.CalcResult

func AppendCalcResult(resStruct entities.CalcResult) {
	memoryCalc = append(memoryCalc, resStruct)
}

func ShowHistory() (history []entities.CalcResult) {
	return memoryCalc
}
