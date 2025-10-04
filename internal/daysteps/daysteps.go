package daysteps

import (
	"fmt"
	"log"
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
	temp := strings.Split(data, ",")
	if len(temp) != 2 {
		errData := fmt.Errorf("error wrong data")
		return 0, 0, errData
	}
	stepString := temp[0]
	step, errStep := strconv.Atoi(stepString)
	if errStep != nil || step <= 0 {
		errStep = fmt.Errorf("error in step")
		return 0, 0, errStep
	}
	walkDurString := temp[1]
	walkDur, errWalkDur := time.ParseDuration(walkDurString)
	if errWalkDur != nil || float64(walkDur) <= 0 {
		errWalkDur = fmt.Errorf("error in walk duration")
		return 0, 0, errWalkDur
	}
	return step, walkDur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	step, walkDur, errParse := parsePackage(data)
	if errParse != nil {
		log.Println(errParse)
		return ""
	}
	distance := (float64(step) * stepLength) / mInKm
	spentCal, errCal := spentcalories.WalkingSpentCalories(step, weight, height, walkDur)
	if errCal != nil {
		log.Println(errCal)
	}
	actionInfo := "Количество шагов: " + fmt.Sprint(step) + ".\nДистанция составила " + fmt.Sprintf("%.2f", distance) + " км.\nВы сожгли " + fmt.Sprintf("%.2f", spentCal) + " ккал.\n"
	return actionInfo
}
