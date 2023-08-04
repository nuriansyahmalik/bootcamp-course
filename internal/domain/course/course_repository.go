package course

import (
	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/gofrs/uuid"
)

var (
	courseQueries = struct {
		selectCourse         string
		insertCourse         string
		selectCourseByUserID string
	}{
		selectCourse: `
            SELECT 
				c.id,
				c.title,
				c.content,
            	c.user_id
			FROM course c	
`,
		insertCourse: `
			INSERT INTO course (
			 	id,
			    title,
			    content,
			    user_id
			) VALUES (
				:id,
			    :title,
			    :content,
			    :user_id)`,
		selectCourseByUserID: `
			SELECT 
				c.id,
				c.title,
				c.content,
				c.user_id
			FROM course c
			WHERE c.user_id = ?`,
	}
)

type CourseRepository interface {
	Create(course Course) (err error)
	ExistsByID(id uuid.UUID) (exists bool, err error)
	GetAllCoursesByUserID(userID uuid.UUID) ([]Course, error)
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
func (c *CourseRepositoryMySQL) GetAllCoursesByUserID(userID uuid.UUID) ([]Course, error) {
	rows, err := c.DB.Read.Query(courseQueries.selectCourseByUserID, userID)
	if err != nil {
		logger.ErrorWithStack(err)
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		var userID uuid.NullUUID
		err := rows.Scan(&course.ID, &course.Title, &course.Content, &userID)
		if err != nil {
			logger.ErrorWithStack(err)
			return nil, err
		}
		if userID.Valid {
			course.UserID = userID.UUID
		}
		courses = append(courses, course)
	}

	return courses, nil
}
