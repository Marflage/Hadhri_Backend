package handlers

import (
	"context"
	db "hadhri/Db"
	dtos "hadhri/Dtos"
	querymodels "hadhri/QueryModels"
	responses "hadhri/Responses"
	utils "hadhri/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetCoursePlans(c *gin.Context) {
	res := &responses.ApiResponse{}

	// TODO: Move this logic to get dbConn in a reusable method.
	dbConn, err := db.InitDb()

	if err != nil {
		// c.Errors.JSON()
		// c.Abort()
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	result, err := getCoursePlans(dbConn)

	if err != nil {
		res.Error = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, res) // gin.H{"error": err.Error()})
		return
	}

	if len(result) == 0 {
		res.Error = "Course plans not found."
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	coursePlans := createCoursePlans(result)

	res.Data = responses.GetCoursePlans{CoursePlans: coursePlans}
	res.Message = "Successfully fetched course plans"

	c.JSON(http.StatusOK, res)
}

// TODO: Move this into a respository.
func getCoursePlans(dbConn *pgx.Conn) ([]querymodels.GetCoursePlans, error) {
	// Order of the columns is not necessary as long as they match the field names in the query model.
	query := `
		SELECT c.name       AS CourseName,
			c.id         AS CourseId,
			sch.name     AS ClassScheduleName,
			sch.id       AS ClassScheduleId,
			ses.name     AS ClassSessionName,
			ses.id       AS ClassSessionId,
			avs.semester AS AvailableSemester
		FROM course_plans cp
			JOIN courses c
				ON cp.course_id = c.id
			JOIN class_schedules sch
				ON cp.class_schedule_id = sch.id
			JOIN class_sessions ses
				ON cp.class_session_id = ses.id
			JOIN available_semesters avs
				ON avs.course_plan_id = cp.id
		WHERE cp.is_active
		GROUP BY c.name, c.id, sch.name, sch.id, ses.name, ses.id, avs.semester;
	`

	rows, err := dbConn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[querymodels.GetCoursePlans])

	return result, nil
}

// TODO: All the below methods in the service layer.
func createCoursePlans(queryResult []querymodels.GetCoursePlans) []dtos.CoursePlan {
	var coursePlans []dtos.CoursePlan

	courseGroups := utils.GroupBy(queryResult, func(e querymodels.GetCoursePlans) (int, string) {
		return e.CourseId, e.CourseName
	})

	for key, value := range courseGroups {
		// TODO: Test if an instance is created without assigning the address of type.
		classSchedules := createClassSchedules(value.Data)

		// TODO: Test if an instance is created without assigning the address of type.
		course := dtos.Course{
			Id:             key,
			Name:           value.Name,
			ClassSchedules: classSchedules,
		}

		coursePlan := dtos.CoursePlan{Course: course}

		coursePlans = append(coursePlans, coursePlan)
	}

	return coursePlans
}

func createClassSchedules(courseGroup []querymodels.GetCoursePlans) []dtos.ClassSchedule {
	var classSchedules []dtos.ClassSchedule

	classScheduleGroups := utils.GroupBy(courseGroup, func(e querymodels.GetCoursePlans) (int, string) {
		return e.ClassScheduleId, e.ClassScheduleName
	})

	// Extract class schedules.
	for key, value := range classScheduleGroups {
		classSessions := createClassSessions(value.Data)

		// TODO: Test if class sessions are populated even though assignment done before population.
		classSchedule := dtos.ClassSchedule{
			Id:            key,
			Name:          value.Name,
			ClassSessions: classSessions,
		}

		classSchedules = append(classSchedules, classSchedule)
	}

	return classSchedules
}

func createClassSessions(classScheduleGroup []querymodels.GetCoursePlans) []dtos.ClassSession {
	var classSessions []dtos.ClassSession

	classSessionGroups := utils.GroupBy(classScheduleGroup, func(e querymodels.GetCoursePlans) (int, string) {
		return e.ClassSessionId, e.ClassSessionName
		// return e.ClassSessionName
	})

	// Extract class sessions.
	for key, value := range classSessionGroups {
		availableSemesters := extractAvailableSemesters(value.Data)

		classSession := dtos.ClassSession{
			Id:                 key,
			Name:               value.Name,
			AvailableSemesters: availableSemesters,
		}

		classSessions = append(classSessions, classSession)
	}

	return classSessions
}

func extractAvailableSemesters(classSessionGroup []querymodels.GetCoursePlans) []int {
	// TODO: Create a method to get a list of available semesters.
	var availableSemesters []int

	// Extract available semesters.
	for _, item := range classSessionGroup {
		availableSemesters = append(availableSemesters, item.AvailableSemester)
	}

	return availableSemesters
}
