package grade

import "fmt"

func validateGrades(grades []float32) error {
	if len(grades) <= 0 {
		return fmt.Errorf("no grades were recieved")
	}

	for i, grade := range grades {
		if grade < 0.0 || grade > 10.0 {
			return fmt.Errorf("grade %.2f at index %d is invalid", grade, i)
		}
	}

	return nil
}

func InsertGrades(output *[]float32, grades []float32) error {
	if err := validateGrades(grades); err != nil {
		return err
	}

	*output = grades
	return nil
}

func CalculateGradeAvg(grades []float32) *float32 {
	if err := validateGrades(grades); err != nil {
		return nil
	}

	var sum float32

	for _, grade := range grades {
		sum += grade
	}

	avg := sum / float32(len(grades))
	return &avg
}
