package test

import (
	"testing"
	"github.com/2103561941/douyin/repository"
)

func TestDB(t *testing.T) {
	if err := repository.Init(); err != nil {
		t.Log(err)
		return
	}
	
	t.Log("success!")
}
