package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/maddevsio/instagram-agent/models"
	"github.com/stretchr/testify/assert"
)

type mockDB struct{}

func (mdb *mockDB) CountersCreate(*models.Counter) error {
	return nil
}

func (mdb *mockDB) CountersFindLast() (*models.Counter, error) {
	counter := &models.Counter{time.Date(2016, time.October, 21, 0, 0, 0, 0, time.UTC), "testuser", 10, 15, 20}
	return counter, nil
}

func (mdb *mockDB) CountersLastMonth() ([]*models.AverageCounter, error) {
	avgCounters := make([]*models.AverageCounter, 0)
	avgCounters = append(avgCounters, &models.AverageCounter{"2016-10-04", 10, 15, 20})
	avgCounters = append(avgCounters, &models.AverageCounter{"2016-10-05", 12, 16, 22})
	return avgCounters, nil
}

func TestCountersLast(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := new(http.Request)
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/counters")
	env := Env{db: &mockDB{}}
	countersJSON := `{"created":"2016-10-21T00:00:00Z","username":"testuser","media":10,"follows":15,"followed_by":20}`

	// Assertions
	if assert.NoError(t, env.countersLast(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, countersJSON, rec.Body.String())
	}
}

func TestCountersLastMonth(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := new(http.Request)
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/counters-last-month")
	env := Env{db: &mockDB{}}
	expectedJSON := `{"media":[{"date":"2016-10-04","value":10},{"date":"2016-10-05","value":12}],"follows":[{"date":"2016-10-04","value":15},{"date":"2016-10-05","value":16}],"followed_by":[{"date":"2016-10-04","value":20},{"date":"2016-10-05","value":22}]}`

	if assert.NoError(t, env.countersLastMonth(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedJSON, rec.Body.String())
	}
}
