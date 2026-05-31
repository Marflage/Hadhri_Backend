package webapi

type PathParam struct {
	Id uint `uri:"id" binding:"required,gt=0"`
}
