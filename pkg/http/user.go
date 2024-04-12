package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rodblg/ReservationRestAPIGolang/pkg/http/models"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/get-users", c.GetUsers)

	return r
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	u := models.User{
		FirstName: "Rodrigo",
		LastName:  "Blancas",
	}
	render.JSON(w, r, u)
}
