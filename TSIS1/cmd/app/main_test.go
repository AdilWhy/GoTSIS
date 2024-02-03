package main

import (
	"TSIS1/pkg/app"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetCompanions(t *testing.T) {
	req, err := http.NewRequest("GET", "/companions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(companions)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []models.Companion
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(response, models.Companions) {
		t.Errorf("Unexpected body: got %v want %v", response, models.Companions)
	}
}

func TestGetCompanionByName(t *testing.T) {
	req, err := http.NewRequest("GET", "/{name}", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"name": "Gale"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getCompanion)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"name":"Gale","race":"Human","class":"Wizard","background":"Sage","attributes":{"Str":8,"Dex":13,"Con":15,"Int":17,"Wis":10,"Cha":12}}`
	if rr.Body.String() != expected {
		t.Errorf("Unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}
