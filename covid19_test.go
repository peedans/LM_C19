package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"Data":[{"ConfirmDate":"2021-05-04","No":5,"Age":51,"Gender":"Female","GenderEn":"Female","Nation":"China","NationEn":"China","Province":"Phrae","ProvinceId":46,"District":"hatyai","ProvinceEn":"Phrae","StatQuarantine":5}]}`))
	if err != nil {
		return
	}
}

func TestHTTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))

	defer server.Close()

	want := StructData{
		Data: []Response{{ConfirmDate: "2021-05-04", No: float64(5), Age: float64(51), Gender: "Female", GenderEn: "Female", Nation: "China", NationEn: "China", Province: "Phrae", ProvinceId: 46, District: "hatyai", ProvinceEn: "Phrae", StatQuarantine: 5}},
	}

	t.Run("Happy server response", func(t *testing.T) {

		resp, err := FetchData(server.URL)

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%+v), got (%+v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})
}

func TestCountAge(t *testing.T) {
	wantLess30 := 3
	want31to60 := 3
	wantMore61 := 1
	wantNil := 2

	covidData, _ := mockData()

	got := CountAge(covidData)

	if got["0-30"] != wantLess30 {
		t.Errorf("expected (%+v), got (%+v)", wantLess30, got["0-30"])
	} else if got["31-60"] != want31to60 {
		t.Errorf("expected (%+v), got (%+v)", want31to60, got["31-60"])
	} else if got["61+"] != wantMore61 {
		t.Errorf("expected (%+v), got (%+v)", wantMore61, got["61+"])
	} else if got["N/A"] != wantNil {
		t.Errorf("expected (%+v), got (%+v)", wantNil, got["N/A"])
	}
}

func TestCountProvince(t *testing.T) {
	t.Run("CountProvince", func(t *testing.T) {
		wantSongkhla := 4
		wantYala := 3
		wantNil := 2

		covidData, _ := mockData()

		got := CountProvince(covidData)

		if got["Yala"] != wantYala {
			t.Errorf("expected (%+v), got (%+v)", wantYala, got["Yala"])
		} else if got["Songkhla"] != wantSongkhla {
			t.Errorf("expected (%+v), got (%+v)", wantSongkhla, got["Songkhla"])
		} else if got["N/A"] != wantNil {
			t.Errorf("expected (%+v), got (%+v)", wantNil, got["N/A"])
		}
	})
}

func mockData() (StructData, error) {
	//Boundary value
	covidData := StructData{
		Data: []Response{{ConfirmDate: "2022-12-04", No: float64(7), Age: float64(0), Gender: "Female", GenderEn: "Male", Nation: "Thailand", NationEn: "Thailand", Province: "Yala", ProvinceId: 77, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 60},
			{ConfirmDate: "2022-12-04", No: float64(51), Age: float64(1), Gender: "Male", GenderEn: "Male", Nation: "Thailand", NationEn: "Thailand", Province: "Yala", ProvinceId: 99, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 50},
			{ConfirmDate: "2021-01-25", No: float64(60), Age: float64(30), Gender: "Female", GenderEn: "Male", Nation: "China", NationEn: "China", Province: "Songkhla", ProvinceId: 107, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 52},
			{ConfirmDate: "1997-01-12", No: float64(21), Age: float64(31), Gender: "Male", GenderEn: "Female", Nation: "France", NationEn: "France", Province: "Songkhla", ProvinceId: 60, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 63},
			{ConfirmDate: "1900-12-01", No: float64(36), Age: float64(45), Gender: "Female", GenderEn: "Female", Nation: "Ukraine", NationEn: "Ukraine", Province: nil, ProvinceId: 8, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 54},
			{ConfirmDate: "1990-09-26", No: float64(17), Age: float64(60), Gender: "Female", GenderEn: "Male", Nation: "SouthKorea", NationEn: "SouthKorea", Province: "Yala", ProvinceId: 16, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 23},
			{ConfirmDate: "2010-10-18", No: float64(99), Age: float64(61), Gender: "Male", GenderEn: "Male", Nation: "UnitedStates", NationEn: "UnitedStates", Province: "Songkhla", ProvinceId: 85, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 95},
			{ConfirmDate: "2005-12-07", No: float64(80), Age: nil, Gender: "Female", GenderEn: "Female", Nation: "Singapore", NationEn: "Singapore", Province: "Songkhla", ProvinceId: 64, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 64},
			{ConfirmDate: "2016-02-29", No: float64(70), Age: nil, Gender: "Male", GenderEn: "Male", Nation: "Japan", NationEn: "Japan", Province: nil, ProvinceId: 52, District: "Sateng", ProvinceEn: "Yala", StatQuarantine: 35}},
	}
	return covidData, nil
}
