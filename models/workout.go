package models

import "time"

type Workout struct {
	Exercise Exercise  `json:"exercise"`
	Reps     int       `json:"reps"`
	Weight   float64   `json:"weight,omitempty"`
	Date     time.Time `json:"date"`
}

var WorkoutList []Workout

func init() {
	today := time.Now().Truncate(24 * time.Hour) // Set to start of today

	WorkoutList = []Workout{
		{
			Exercise: Pushups,
			Reps:     20,
			Weight:   0, // bodyweight exercise
			Date:     today,
		},
		{
			Exercise: Squats,
			Reps:     15,
			Weight:   135.5, // in pounds or kg, depending on your preference
			Date:     today,
		},
		{
			Exercise: Pullups,
			Reps:     10,
			Weight:   0, // bodyweight exercise
			Date:     today,
		},
		{
			Exercise: BicepCurls,
			Reps:     12,
			Weight:   25.0,
			Date:     today,
		},
		{
			Exercise: Plank,
			Reps:     1,
			Weight:   0,                          // bodyweight exercise
			Date:     today.Add(-24 * time.Hour), // Yesterday's workout
		},
	}
}
