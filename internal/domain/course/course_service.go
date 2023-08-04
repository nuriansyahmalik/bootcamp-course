package course

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type CourseService interface {
	Create(requestFormat CourseRequestFormat, id uuid.UUID) (course Course, err error)
	GetAllCourse() ([]Course, error)
}
type CourseServiceImpl struct {
	CourseRepository CourseRepository
	Config           *configs.Config
}

func ProvideCourseServiceImpl(courseRepository CourseRepository, config *configs.Config) *CourseServiceImpl {
	return &CourseServiceImpl{CourseRepository: courseRepository, Config: config}
}

func (c *CourseServiceImpl) Create(requestFormat CourseRequestFormat, id uuid.UUID) (course Course, err error) {
	course, err = course.CourseRequestFormat(requestFormat, id)
	if err != nil {
		return
	}

	if err != nil {
		return course, failure.BadRequest(err)
	}
	err = c.CourseRepository.Create(course)
	if err != nil {
		log.Info().Err(err)
	}
	return
}

func (c *CourseServiceImpl) GetAllCourse() ([]CourseResponseFormat, error) {
	courses, err := c.CourseRepository.GetAllCourses()
	if err != nil {
		return nil, err
	}

	var courseResponses []CourseResponseFormat
	for _, course := range courses {
		courseResponses = append(courseResponses, course.ToResponseFormat())
	}

	return courseResponses, nil
}
