package handlers

import (
	"net/http"

	"github.com/eylemabz/go-course/bookings/pkg/config"
	"github.com/eylemabz/go-course/bookings/pkg/models"
	"github.com/eylemabz/go-course/bookings/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sest the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page andler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home-page.html", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello ,again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["repote_ip"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, "about-page.md", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.md", &models.TemplateData{})
}

//Generals render room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "general.page.md", &models.TemplateData{})
}

//Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.md", &models.TemplateData{})
}

//Availability renders the availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.md", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.md", &models.TemplateData{})
}
