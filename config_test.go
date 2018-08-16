package main

import (
	"os"
	"testing"
)

func TestConf(t *testing.T) {

	t.Run("serverConf.port should be filled from environment var MYAC_PORT", func(t *testing.T){
		os.Setenv("MYAC_PORT", "8877")
		c := serverConf{}
//		configure()
		got := c.Server.Port
		want := os.Getenv("MYAC_PORT")
		if got != want {
			t.Errorf("\ngot:\n%s \nwant:\n%s", got, want)
		}
	})
}


//use cmp.Equals and cmp.Diff for tests