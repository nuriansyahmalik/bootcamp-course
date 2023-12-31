package course

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type CourseService interface {
	Create(requestFormat CourseRequestFormat, userID uuid.UUID) (course Course, err error)
	GetAllCoursesByUserID(userID uuid.UUID) ([]CourseResponseFormat, error)
}
type CourseServiceImpl struct {
	CourseRepository CourseRepository
	Config           *configs.Config
}

func ProvideCourseServiceImpl(courseRepository CourseRepository, config *configs.Config) *CourseServiceImpl {
	return &CourseServiceImpl{CourseRepository: courseRepository, Config: config}
}

func (c *CourseServiceImpl) Create(requestFormat CourseRequestFormat, userID uuid.UUID) (course Course, err error) {
	course, err = course.CourseRequestFormat(requestFormat, userID)
	if err != nil {
		return
	}
	if err != nil {
		return course, failure.BadRequest(err)
	}
	err = c.CourseRepository.Create(course)
	if err != nil {
		log.Info().Err(err)
		return
	}
	return
}

func (c *CourseServiceImpl) GetAllCoursesByUserID(userID uuid.UUID) ([]CourseResponseFormat, error) {
	courses, err := c.CourseRepository.GetAllCoursesByUserID(userID)
	if err != nil {
		return nil, err
	}

	var courseResponses []CourseResponseFormat
	for _, course := range courses {
		courseResponses = append(courseResponses, course.ToResponseFormat())
	}

	return courseResponses, nil
}
