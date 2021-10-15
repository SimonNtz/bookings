package render

import (
	"net/http"
	"testing"

	"github.com/simonntz/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Fatal(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplate = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&ww, r, "non-existant.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Errorf("rendered template that not exists")
	}
}

func getSession() (*http.Request, error) {
	// create an empty http GET request
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	// extract request context
	ctx := r.Context()

	// add new session data to context
	ctx, err = session.Load(ctx, "") //r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplate(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplate = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
