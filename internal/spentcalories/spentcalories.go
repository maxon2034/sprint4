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
	temp := strings.Split(data, ",")
	if len(temp) != 3 {
		errData := fmt.Errorf("error wrong data")
		return 0, "", 0, errData
	}
	stepString := temp[0]
	step, errConvToInt := strconv.Atoi(stepString)
	if errConvToInt != nil || step <= 0 {
		errConvToInt = fmt.Errorf("error in step data")
		return 0, "", 0, errConvToInt
	}
	walkDurString := temp[2]
	walkDur, errConvToDur := time.ParseDuration(walkDurString)
	if errConvToDur != nil || walkDur <= 0 {
		errConvToDur = fmt.Errorf("error in duration data")
		return 0, "", 0, errConvToDur
	}
	activityType := temp[1]
	if activityType != "Ходьба" && activityType != "Бег" {
		errActivityType := fmt.Errorf("неизвестный тип тренировки")
		return 0, "", 0, errActivityType
	}
	return step, activityType, walkDur, nil
}

func distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	dist := (float64(steps) * stepLen) / mInKm
	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		fmt.Errorf("wrong distance")
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, walkDur, errParse := parseTraining(data)
	if errParse != nil {
		log.Println(errParse)
		return "", errParse
	}
	var spentCalories float64
	var errWalking, errRunning error
	switch activityType {
	case "Ходьба":
		spentCalories, errWalking = WalkingSpentCalories(steps, weight, height, walkDur)
	case "Бег":
		spentCalories, errRunning = RunningSpentCalories(steps, weight, height, walkDur)
	}
	if errWalking != nil {
		log.Println(errWalking)
	}
	if errRunning != nil {
		log.Println(errRunning)
	}
	activityInfo := "Тип тренировки: " + activityType + "\n"
	durationInfo := "Длительность: " + fmt.Sprintf("%.2f", walkDur.Hours()) + " ч.\n"
	distanceInfo := "Дистанция: " + fmt.Sprintf("%.2f", distance(steps, height)) + " км.\n"
	speedInfo := "Скорость: " + fmt.Sprintf("%.2f", meanSpeed(steps, height, walkDur)) + " км/ч\n"
	burntCaloriesInfo := "Сожгли калорий: " + fmt.Sprintf("%.2f", spentCalories) + "\n"
	return activityInfo + durationInfo + distanceInfo + speedInfo + burntCaloriesInfo, nil
}
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		errSteps := fmt.Errorf("wrong steps data")
		return 0, errSteps
	}
	if weight <= 0 {
		errWeight := fmt.Errorf("wrong weight data")
		return 0, errWeight
	}
	if height <= 0 {
		errHeight := fmt.Errorf("wrong height data")
		return 0, errHeight
	}
	if float64(duration) <= 0 {
		errDuration := fmt.Errorf("wrong duration data")
		return 0, errDuration
	}
	runningSpentCalories := (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH
	return runningSpentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		errSteps := fmt.Errorf("wrong steps data")
		return 0, errSteps
	}
	if weight <= 0 {
		errWeight := fmt.Errorf("wrong weight data")
		return 0, errWeight
	}
	if height <= 0 {
		errHeight := fmt.Errorf("wrong height data")
		return 0, errHeight
	}
	if float64(duration) <= 0 {
		errDuration := fmt.Errorf("wrong duration data")
		return 0, errDuration
	}
	walkingSpentCalories := ((weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH) * walkingCaloriesCoefficient
	return walkingSpentCalories, nil
}
