package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// parseTraning парсинг строки
func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slcData := strings.Split(data, ",")
	if len(slcData) != 3 {
		return 0, "", 0, errors.New("некорректное содержимое data")
	}
	numberOfStep, err := strconv.Atoi(slcData[0])
	if err != nil {
		return 0, "", 0, err
	}
	if numberOfStep <= 0 {
		return 0, "", 0, errors.New("неккоректное количество шагов")
	}
	traningTime, err := time.ParseDuration(slcData[2])
	if err != nil {
		return 0, "", 0, err
	}
	if traningTime <= 0 {
		return 0, "", 0, errors.New("неккоректное время")
	}
	return numberOfStep, slcData[1], traningTime, nil
}

// distance вычисление дистанции в километрах
func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	strideLength := height * stepLengthCoefficient
	walkDistance := strideLength * float64(steps)
	return walkDistance / float64(mInKm)
}

// meanSpeed вычисление средней скорости в км/ч
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	return distanceKm / duration.Hours()
}

// TrainingInfo возврат информации о тренировке
func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	numberOfStep, typeOfActivity, traningTime, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	var distanceKm, avarageSpeed, numberOfCalories float64
	switch typeOfActivity {
	case "Ходьба":
		numberOfCalories, err = WalkingSpentCalories(numberOfStep, weight, height, traningTime)
		if err != nil {
			return "", err
		}
	case "Бег":
		numberOfCalories, err = RunningSpentCalories(numberOfStep, weight, height, traningTime)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	distanceKm = distance(numberOfStep, height)
	avarageSpeed = meanSpeed(numberOfStep, height, traningTime)

	stringResult := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n", typeOfActivity, traningTime.Hours(), distanceKm)
	stringResult += fmt.Sprintf("Скорость: %.2f км/ч\nСожгли калорий: %.2f\n", avarageSpeed, numberOfCalories)

	return stringResult, nil

}

// RunningSpentCalories вычисление потраченых каллорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные данные")
	}
	avarageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * avarageSpeed * durationInMinutes) / float64(minInH), nil
}

// WalkingSpentCalories вычисление потраченых каллорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные данные")
	}
	avarageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	numberOfCalories := (weight * avarageSpeed * durationInMinutes) / float64(minInH)
	return numberOfCalories * walkingCaloriesCoefficient, nil
}
