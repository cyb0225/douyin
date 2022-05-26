package test

import (
	"testing"
	initdao "github.com/2103561941/douyin/repository/init"
)

func TestDB(t *testing.T) {
	if err := initdao.Init(); err != nil {
		t.Log(err)
		return
	}
	
	t.Log("success!")
}
