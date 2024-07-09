package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com.br/Leodf/bookings/internal/model"
)

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d, got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}

	}
}

func TestRepositoryreservation(t *testing.T) {
	reservation := model.Reservation{
		ID: 1,
		Room: model.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session(reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test with non-existent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepositoryPostReservation(t *testing.T) {
	reqBody := "start_date=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code: %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for missing body
	req, _ = http.NewRequest("POST", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for posting missing body: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid start date
	reqBody = "start_date=invalid"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid start date: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
	// test for invalid end date
	reqBody = "start_date=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=invalid")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid end date: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
	// test for invalid id
	reqBody = "start_date=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=invalid")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid id: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
	// test for invalid data
	reqBody = "start_date=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=J")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code for invalid data: %d, wanted %d", rr.Code, http.StatusSeeOther)
	}
	// test for failure to insert reservation into database
	reqBody = "start_date=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=2")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when trying fail inserting reservation: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
	// test for failure to insert restriction into database
	reqBody = "start_date=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=123456789")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1000")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when trying fail inserting reservation: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepositoryAvailabilityJSON(t *testing.T) {
	// test invalid parseform
	req, _ := http.NewRequest("POST", "/search-availability-json", nil)
	// get context with session
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// make handler handlerfunc
	handler := http.HandlerFunc(Repo.AvailabilityJSON)
	// get response recorder
	rr := httptest.NewRecorder()
	// make request to our handler
	handler.ServeHTTP(rr, req)

	var j jsonResponse
	err := json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	// test for invalid start date
	reqBody := "start=invalid"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))
	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// make handler handlerfunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)
	// get response recorder
	rr = httptest.NewRecorder()
	// make request to our handler
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("AvailabilityJSON handler returned wrong response code for invalid start date: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid end date
	reqBody = "start=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=invalid")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))
	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// make handler handlerfunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)
	// get response recorder
	rr = httptest.NewRecorder()
	// make request to our handler
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("AvailabilityJSON handler returned wrong response code for invalid end date: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid end date
	reqBody = "start=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=invalid")
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))
	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// make handler handlerfunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)
	// get response recorder
	rr = httptest.NewRecorder()
	// make request to our handler
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("AvailabilityJSON handler returned wrong response code for invalid room id: %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for return a database connect from rooms
	reqBody = "start=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1000")
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))
	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// make handler handlerfunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)
	// get response recorder
	rr = httptest.NewRecorder()
	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	// test rooms are not available
	reqBody = "start=01/01/2050"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=02/01/2050")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
	// create request
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))
	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// make handler handlerfunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)
	// get response recorder
	rr = httptest.NewRecorder()
	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Sesssion"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
