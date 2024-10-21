package main

import (
	"os"
	"testing"

	"github.com.Shashwat1977.todo/internal/model"
)

func Test_Add(t *testing.T) {
	l := &model.List{}

	task := "New Task"

	l.Add(task)
	if len(*l) == 0 {
		t.Errorf("Length should be 1")
	}
	if (*l)[0].Task != task {
		t.Errorf("Not matching")
	}
}

func Test_Complete(t *testing.T) {
	l := &model.List{}

	task := "New Task"
	l.Add(task)

	if (*l)[0].Done {
		t.Errorf("Should be marked as false")
	}

	l.Complete(1)
	if !(*l)[0].Done {
		t.Errorf("Should be marked as false")
	}
}

func Test_Delete(t *testing.T) {
	l := &model.List{}

	l.Add("1st task")
	l.Add("2nd task")
	l.Add("3rd task")

	if len(*l) != 3 {
		t.Errorf("Length should be 3")
	}

	l.Delete(2)
	if len(*l) != 2 {
		t.Errorf("Length should be 2")
	}
}

func TestSaveAndGet(t *testing.T) {
	l := &model.List{}

	l.Add("1st task")
	l.Add("2nd task")
	l.Add("3rd task")

	l2 := model.List{}

	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Errorf("Error occurred while creating the test file")
	}

	defer os.Remove(file.Name())

	if err := l.Save(file.Name()); err != nil {
		t.Errorf("Error occurred while saving the list to a file")
	}

	if err := l2.Get(file.Name()); err != nil {
		t.Errorf("Error occurred while trying to read from file")
	}

	if len(l2) != 3 {
		t.Errorf("There are errors in Save and Get operations")
	}
}
