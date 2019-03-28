package main

import (
	"testing"
)

func TestKontoORM(t *testing.T){
	cases := []struct {
		kontonummer int; Kunde string; penger float64

	}{
		{1, "Jan", 700.00},
		{2, "Gaudus", 30.00},
		{3, "Kari", 1000.00},
	}

	var db, err = OpprettForbindelse()
	if err !=nil{
		t.Errorf("connecting issues")
	}
	defer db.Close()
	WipeDatabase(db) //vil ha tom og frisk database


	for _, c := range cases {
		LagBruker(db, c.Kunde, c.penger)

		got := GetKonto(db, c.Kunde)

		if got.Kunde != c.Kunde && got.Kontonummer != c.kontonummer && got.Penger != c.penger {
			t.Errorf("incorrect customer, got: %s, want: %s", got.Kunde, c.Kunde)
		}
	}

}

