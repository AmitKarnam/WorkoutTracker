package controllers

type MuscleGroupController struct{}

// func (mc *MuscleGroupController) Get(c *gin.Context) {
// 	DbConn.First(&models.MuscleGroup{})
// }

// func (mc *MuscleGroupController) Post(c *gin.Context) {
// 	var newMuscleGroup models.MuscleGroup
// 	if err := c.ShouldBindBodyWithJSON(&newMuscleGroup); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	models.MuscleGroupList = append(models.MuscleGroupList, newMuscleGroup)
// 	c.JSON(http.StatusCreated, gin.H{"data": newMuscleGroup})
// }

// func (mc *MuscleGroupController) Delete(c *gin.Context) {
// 	musclegroup := c.Param("id")
// 	for _, mg := range models.MuscleGroupList {
// 		if mg == models.MuscleGroup(musclegroup) {

// 		}
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("deleted : %s", musclegroup)})
// }
