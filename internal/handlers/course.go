package handlers

import (
	"encoding/json"
	"github.com/evermos/boilerplate-go/internal/domain/course"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/transport/http/middleware"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
)

type CourseHandler struct {
	CourseService  course.CourseService
	AuthMiddleware *middleware.Authentication
}

func ProvideCourseHandler(courseService course.CourseService, authMiddleware *middleware.Authentication) CourseHandler {
	return CourseHandler{
		CourseService:  courseService,
		AuthMiddleware: authMiddleware,
	}
}

func (h *CourseHandler) Router(r chi.Router) {
	r.Route("/course", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.ValidateJWTMiddleware)
			r.Get("/", h.GetCourse)
			r.Post("/", h.CreateCourse)
		})
	})
}

func (h *CourseHandler) GetCourse(w http.ResponseWriter, r *http.Request) {
	courses, err := h.CourseService.GetAllCourse()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responses := course.NewCourseResponsesFromCourses(courses)
	jsonResponse, err := json.Marshal(responses)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.WithJSON(w, http.StatusOK, jsonResponse)
}
func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestFormat course.CourseRequestFormat
	err := decoder.Decode(&requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	err = shared.GetValidator().Struct(requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	userID, _ := uuid.NewV4() // TODO: read from context
	user, err := h.CourseService.Create(requestFormat, userID)
	if err != nil {
		log.Error().Msg("error disni")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, user)
}
