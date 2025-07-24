package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	parsedData := strings.Split(datastring, ",")
	if len(parsedData) != 2 {
		return errors.New("incorrect input format")
	}

	ds.Steps, err = strconv.Atoi(parsedData[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return errors.New("number of steps must be greater than 0")
	}

	ds.Duration, err = time.ParseDuration(parsedData[1])
	if err != nil {
		return err
	}
	if ds.Duration <= 0 {
		return errors.New("training duration must be greater than 0")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	walkingSpentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	outputFormat := "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"
	output := fmt.Sprintf(outputFormat, ds.Steps, distance, walkingSpentCalories)

	return output, nil
}
