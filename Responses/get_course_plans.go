package responses

import dtos "hadhri/Dtos"

type GetCoursePlans struct {
	CoursePlans []dtos.CoursePlan `json:"coursePlans"`
}
