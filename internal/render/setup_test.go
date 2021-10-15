package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/simonntz/bookings/internal/config"
	"github.com/simonntz/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// tell what we are storing to session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 22 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func (tw *myWriter) WriteHeader(statusCode int) {}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}
