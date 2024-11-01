package models

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	MuscleGroup MuscleGroup `json:"muscle_group"`
}

var ExerciseList []Exercise

var Pushups Exercise = Exercise{
	Name:        "Pushups",
	Description: "Pushups are a great exercise for the chest, shoulders, and triceps.",
	MuscleGroup: Chest,
}

var Squats = Exercise{
	Name:        "Squats",
	Description: "Squats target the legs, particularly the quadriceps, hamstrings, and glutes.",
	MuscleGroup: Legs,
}

var Pullups = Exercise{
	Name:        "Pull-ups",
	Description: "Pull-ups are an excellent exercise for building upper body strength, focusing on the back and biceps.",
	MuscleGroup: Back,
}

var BicepCurls = Exercise{
	Name:        "Bicep Curls",
	Description: "Bicep curls isolate and strengthen the biceps muscles in the upper arms.",
	MuscleGroup: Arms,
}

var Plank = Exercise{
	Name:        "Plank",
	Description: "Planks are a core strengthening exercise that also engages the shoulders, arms, and glutes.",
	MuscleGroup: Abs,
}

func init() {
	ExerciseList = []Exercise{
		Pushups,
		Squats,
		Pullups,
		BicepCurls,
		Plank,
	}
}
