package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {

	parsedData := strings.Split(datastring, ",")
	if len(parsedData) != 3 {
		return errors.New("incorrect input format")
	}

	t.Steps, err = strconv.Atoi(parsedData[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return errors.New("number of steps must be greater than 0")
	}

	t.TrainingType = parsedData[1]

	t.Duration, err = time.ParseDuration(parsedData[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return errors.New("training duration must be greater than 0")
	}

	return err
}

func (t Training) ActionInfo() (string, error) {

	var spentCalories float64
	var err error

	distance := spentenergy.Distance(t.Steps, t.Height)

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("unknown training type")
	}

	outputFormat := "Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n"
	output := fmt.Sprintf(outputFormat, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories)

	return output, nil
}
