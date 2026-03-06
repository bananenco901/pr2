package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат строки: %s", datastring)
	}
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, dist, calories), nil
}
