package main

import (
	"testing"
	"database/sql"
)

func TestKontoORM(t *testing.T){
	cases := []struct {
		in int; string; int
		want bool
	}{
		{1034968543, Jan, 700, false},
		{7492039553, Gaudus, 30, true},
		{1493920394, Kari, 1000,true},
	}
	for _, c := range cases {
		lagBruker(c.in)
		got := getKonto(c.in)

		if got != c.want {
			t.Errorf("select(%b, %s, %b) == %t, want %t", c.in, got, c.want)
		}
	}
}

func setUpMyDB(t *testing.T)(db *sql.DB, cleanup func() error){

}
