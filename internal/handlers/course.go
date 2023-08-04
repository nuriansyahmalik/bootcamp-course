package handlers

import (
	"encoding/json"
	"github.com/evermos/boilerplate-go/internal/domain/course"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/jwt"
	"github.com/evermos/boilerplate-go/transport/http/middleware"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
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
			r.Use(middleware.CheckRole)
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

	response.WithJSON(w, http.StatusOK, responses)
}
func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(jwt.Claims)
	if !ok {
		http.Error(w, "Error Claims", http.StatusUnauthorized)
		return
	}
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
	//id, _ := uuid.NewV4() // TODO: read from context
	user, err := h.CourseService.Create(requestFormat, claims.ID)
	if err != nil {
		log.Error().Msg("error disni")
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, user)
}
