package controller

import (
	"net/http"

	"github.com/omeg21/udemy-throwaway/pkg/config"
	"github.com/omeg21/udemy-throwaway/pkg/models"
	"github.com/omeg21/udemy-throwaway/pkg/render"
)

// Repo the repository used by the controller
var Repo *Repository

// Repository is the repositry type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewController sets the repository for the handlers
func NewController(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there."

	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// func New(w http.ResponseWriter, r *http.Request) {
// 	render.RenderTemplate(w, "new.page.html")
// }
