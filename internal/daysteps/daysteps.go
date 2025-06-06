package daysteps

import (
	"errors"
	"fmt"
	"github.com/vladkohan/go-sprint4/internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, " ")
	if len(parts) != 2 {
		return 0, 0, errors.New("Неверное количество элементов")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("Количество шагов должно быть больше нуля")
	}

	walkTime, err := time.ParseDuration(parts[1])

	if err != nil {
		return 0, 0, err
	}
	if walkTime <= 0 {
		return 0, 0, errors.New("Время прогулки должно быть больше нуля")
	}

	return steps, walkTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkTime, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * stepLength
	distanceKm := distance / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkTime)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("Пройдено %.2f км, потрачено %.2f калорий за %s.", distanceKm, calories, walkTime)

}
