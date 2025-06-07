package spentcalories

import (
	"errors"
	"fmt"
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

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("неверный формат данных")
	}
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || steps <= 0 {
		return 0, "", 0, errors.New("неверный формат количества шагов")
	}
	activity := strings.TrimSpace(parts[1])
	if activity != "Ходьба" && activity != "Бег" {
		return 0, "", 0, errors.New("неизвестный тип тренировки")
	}
	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return 0, "", 0, errors.New("некорректная продолжительность")
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("длительность тренировки должна быть больше нуля")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	speed := distanceKm / (duration.Seconds() / 3600)
	return speed
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("все параметры должны быть больше нуля")
	}

	RunningMeanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * RunningMeanSpeed * durationInMinutes) / minInH
	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("все параметры должны быть больше нуля")
	}

	WalkingMeanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * WalkingMeanSpeed * durationInMinutes) / minInH
	adjustedCalories := spentCalories * walkingCaloriesCoefficient
	return adjustedCalories, nil
}
func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	var calories float64
	var errCalories error

	switch activity {
	case "Бег":
		calories, errCalories = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, errCalories = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("неизвестный вид активности")
	}

	if errCalories != nil {
		return "", errCalories
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	return fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", activity, duration.Hours(), dist, speed, calories), nil
}
