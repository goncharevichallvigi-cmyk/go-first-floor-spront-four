package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slcData := strings.Split(data, ",")
	if len(slcData) != 2 {
		return 0, 0, errors.New("некорректное содержимое data")
	}
	numberOfStep, err := strconv.Atoi(slcData[0])
	if err != nil {
		return 0, 0, err
	}
	if numberOfStep <= 0 {
		return 0, 0, errors.New("некорректное количество шагов")
	}
	walkTime, err := time.ParseDuration(slcData[1])
	if err != nil {
		return 0, 0, err
	}
	if walkTime <= 0 {
		return 0, 0, errors.New("неккоректное время")
	}
	return numberOfStep, walkTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	numberOfStep, walkTime, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if numberOfStep <= 0 {
		return ""
	}
	walkingDistance := float64(numberOfStep) * stepLength
	distanceKm := walkingDistance / float64(mInKm)
	numberOfCalories, err := spentcalories.WalkingSpentCalories(numberOfStep, weight, height, walkTime)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", numberOfStep, distanceKm, numberOfCalories)
}
