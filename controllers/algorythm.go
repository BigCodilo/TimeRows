package controllers

import (
	"fmt"
	"math"
)

//MiddleSlice - метод скользящей средней
func MiddleSlice(inputRow []float64, smoothK float64) float64 {
	outputRow := []float64{}
	resultRow := []float64{}
	var result float64
	Range := int(smoothK) / 2 //Диапазон +- значений которые мы будем брать для нахождения скользящей средней
	isSmoothKOdd := 0
	if int(smoothK)%2 != 0 {
		isSmoothKOdd = 1
	}

	//Бегаем циклом по всему массиву
	for i := 0; i < len(inputRow); i++ {
		var number float64 = 0
		//Проверка на то что мы не выйдем из диапазона ряда при сглаживании
		if i >= Range && i <= len(inputRow)-Range-isSmoothKOdd {
			//Берем несколько чисел с диапазоном сглаживания (который рпавен коэффициенту сглаживания):
			//Ряд[индекс - (диапазон / 2)] ... ряд[индекс + (диапазон / 2)]
			//Затем сумируем и делим на коэффициент сглаждивания
			for j := i - Range; j < i-Range+int(smoothK); j++ {
				number += inputRow[j]
			}
			outputRow = append(outputRow, math.Round(number/smoothK/0.05)*0.05)
			resultRow = append(resultRow, math.Pow(inputRow[i]-number/smoothK, 2))
		}
	}

	for _, value := range resultRow {
		result += value
	}

	fmt.Println("Input row: ", inputRow)
	fmt.Println("Output row: ", outputRow)
	fmt.Println("Result row: ", resultRow)
	fmt.Println("Result: ", result)
	return result
}

//ExpSmooth - метод экспоненциального сглаживания
func ExpSmooth(inputRow []float64, alpha float64) float64 {
	outputRow := []float64{}
	resultRow := []float64{}
	var result float64
	var s0 float64 //Сумма первых нескольких элементов ряда, деленное на их количество

	if len(inputRow) >= 3 {
		s0 = (inputRow[0] + inputRow[1] + inputRow[2]) / 3
	} else {
		s0 = inputRow[0]
	}

	fmt.Println("\nInput row: ", inputRow, "\n")

	for i := 0; i < len(inputRow); i++ {
		var t float64
		if i == 0 {
			t = (1-alpha)*inputRow[i] + alpha*s0
			fmt.Println("Iteration", i, ": (1 -", alpha, ") *", inputRow[i], "+", alpha, "*", s0, "=", t)
		} else {
			t = (1-alpha)*inputRow[i] + alpha*outputRow[i-1]
			fmt.Println("Iteration", i, ": (1 -", alpha, ") *", inputRow[i], "+", alpha, "*", outputRow[i-1], "=", t)
		}
		outputRow = append(outputRow, math.Round(t/0.05)*0.05)
	}

	fmt.Println("\nOutput row: ", outputRow)
	fmt.Println("\n\n")

	for i := 0; i < len(inputRow); i++ {
		resultRow = append(resultRow, math.Pow(inputRow[i]-outputRow[i], 2))
		fmt.Println(inputRow[i], " => ", outputRow[i])
		result += resultRow[i]
	}

	fmt.Println("Result: ", result)
	return result
}
