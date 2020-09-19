package main

import {
	"net/http"
	"net/http/httptest"
	"testing"
}

func TestCreatePAirDevice(t *testing.T){
	req := httptest.NewRequest(http.MethodPost, target: "/pair-device", body:nil)
	rec := httptest.NewRecorder()

	PairDeviceHandler(rec,req)

	if htto.StatusOK != rec.Code{
		t.Error(args...:"expect 200 OK but got ", rec.Code)
	}

	expect := `{"status":"active"}`
	if rec.Body.String() != expect{
		t.Errorf(format: "expected %q but got %q\n", expected, rec.Body.String())
	}
}