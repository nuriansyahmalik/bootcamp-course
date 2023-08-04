package course

import (
	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/gofrs/uuid"
)

var (
	courseQueries = struct {
		selectCourse string
		insertCourse string
	}{
		selectCourse: `
            SELECT 
				c.id,
				c.title,
				c.content
			FROM course c	
`,
		insertCourse: `
			INSERT INTO course (
			 	id,
			    title,
			    content
			) VALUES (
				:id,
			    :title,
			    :content)`,
	}
)

type CourseRepository interface {
	Create(course Course) (err error)
	ExistsByID(id uuid.UUID) (exists bool, err error)
	GetAllCourses() ([]Course, error)
}

type CourseRepositoryMySQL struct {
	DB *infras.MySQLConn
}

func ProvideCourseRepositoryMySQL(db *infras.MySQLConn) *CourseRepositoryMySQL {
	return &CourseRepositoryMySQL{DB: db}
}

func (c *CourseRepositoryMySQL) Create(course Course) (err error) {
	stmt, err := c.DB.Write.PrepareNamed(courseQueries.insertCourse)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(course)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}
func (c *CourseRepositoryMySQL) ExistsByID(id uuid.UUID) (exists bool, err error) {
	err = c.DB.Read.Get(
		&exists,
		"SELECT COUNT(id) FROM course WHERE course.id = ?",
		id.String())
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}
func (c *CourseRepositoryMySQL) GetAllCourses() ([]Course, error) {
	rows, err := c.DB.Read.Query(courseQueries.selectCourse)
	if err != nil {
		logger.ErrorWithStack(err)
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Title, &course.Content)
		if err != nil {
			logger.ErrorWithStack(err)
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
