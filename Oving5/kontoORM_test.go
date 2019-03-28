package main

import "testing"

func TestKontoORM(t *testing.T){
	cases := []struct {
		kontonummer int; Kunde string; penger int

	}{
		{1, "Jan", 700},
		{2, "Gaudus", 30},
		{3, "Kari", 1000},
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
			t.Errorf("select(%b, %s, %b) == %t, want %t", c.Kunde, got, c.Kunde)
		}
	}
}