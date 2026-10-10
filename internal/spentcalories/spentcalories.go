package spentcalories

import (
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

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("ошибка")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("шаги должны быть больше 0")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("время тртенировки должно быть больше 0")
	}

	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	return distanceKm

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	distanceKm := distance(steps, height)
	durationHours := duration.Hours()

	return distanceKm / durationHours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	tmplMsg := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"

	switch trainingType {
	case "Бег":
		distanceKm := distance(steps, height)
		speed := meanSpeed(steps, height, duration)

		calories, err := RunningSpentCalories(
			steps,
			weight,
			height,
			duration,
		)
		if err != nil {
			log.Println(err)
			return "", err
		}

		return fmt.Sprintf(
			tmplMsg,
			trainingType,
			duration.Hours(),
			distanceKm,
			speed,
			calories,
		), nil

	case "Ходьба":
		distanceKm := distance(steps, height)
		speed := meanSpeed(steps, height, duration)

		calories, err := WalkingSpentCalories(
			steps,
			weight,
			height,
			duration,
		)
		if err != nil {
			log.Println(err)
			return "", err
		}

		return fmt.Sprintf(
			tmplMsg,
			trainingType,
			duration.Hours(),
			distanceKm,
			speed,
			calories,
		), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("ошибка")
	}

	speed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("ошибка")
	}

	speed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}
