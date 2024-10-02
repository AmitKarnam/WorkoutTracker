package models

type MuscleGroup string

const (
	Chest MuscleGroup = "chest"
	Back  MuscleGroup = "back"
	Legs  MuscleGroup = "legs"
	Arms  MuscleGroup = "arms"
	Abs   MuscleGroup = "abs"
)

var MuscleGroupList []MuscleGroup

func init() {
	MuscleGroupList = []MuscleGroup{Chest, Back, Legs, Arms, Abs}
}
