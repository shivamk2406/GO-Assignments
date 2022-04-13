package courses

type Course struct {
	Name CourseType
}

func (name Course) validateCourse() bool {
	return name.Name.IsACourseType()
}
