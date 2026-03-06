package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат строки: %s", datastring)
	}
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	t.Steps = steps
	t.TrainingType = strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), dist, speed, calories), nil
}
