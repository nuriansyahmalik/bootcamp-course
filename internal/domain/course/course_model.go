package course

import (
	"encoding/json"
	"github.com/gofrs/uuid"
)

type Course struct {
	ID      uuid.UUID `db:"id"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
}

type CourseRequestFormat struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CourseResponseFormat struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
}

func (c Course) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ToResponseFormat())
}

func (c *Course) CourseRequestFormat(req CourseRequestFormat, id uuid.UUID) (course Course, err error) {
	id, _ = uuid.NewV4()
	course = Course{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}
	courses := make([]Course, 0)
	courses = append(courses, course)
	return
}

func (c *Course) ToResponseFormat() CourseResponseFormat {
	return CourseResponseFormat{
		ID:      c.ID,
		Title:   c.Title,
		Content: c.Content,
	}
}

func NewCourseResponseFromCourse(course Course) CourseResponseFormat {
	return CourseResponseFormat{
		ID:      course.ID,
		Title:   course.Title,
		Content: course.Content,
	}
}

func NewCourseResponsesFromCourses(courses []Course) []CourseResponseFormat {
	var responses []CourseResponseFormat
	for _, course := range courses {
		responses = append(responses, NewCourseResponseFromCourse(course))
	}
	return responses
}
