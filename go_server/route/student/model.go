package student

//Student Model for the module
type Student struct {
	Name   string `form:"name" binding:"required" json:"name"`
	Class  string `form:"class" binding:"required" json:"class"`
	RollNo int    `form:"rollNo" binding:"required" json:"rollNo"`
}
