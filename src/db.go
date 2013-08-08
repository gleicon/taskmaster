// Copyright 2013 taskmaster authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	//	"database/sql"
	"encoding/json"
	"time"
)

type TaskSrc struct {
	URL    string
	Method string
	Body   string
	Policy string
}

type TaskDst struct {
	URL    string
	Policy string
}

// Crontab scheduling
type Task struct {
	Id                int
	UserId            int
	CrontabString     string
	DateCreation      time.Time
	DateLastModified  time.Time
	DateLastExecution time.Time
	Errors            int
	Success           int
	IsActive          bool
	Src               TaskSrc
	Dst               TaskDst
}

func (t Task) String() (s string) {
	b, err := json.Marshal(t)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

// Model
func NewTask(CrontabString string, taskSrc TaskSrc, taskDst TaskDst) (*Task, error) {
	return nil, nil
}

func FindTaskById(Id string) (*Task, error) {
	var t Task
	if err := MySQL.QueryRow(
		"select * from Task where Id=?", Id,
	).Scan(
		&t.Id,
		&t.UserId,
		&t.CrontabString,
		&t.DateCreation,
		&t.DateLastModified,
		&t.DateLastExecution,
		&t.Errors,
		&t.Success,
		&t.IsActive,
		&t.Src,
		&t.Dst,
	); err != nil {
		return nil, err
	}
	return &t, nil
}

func (t Task) ActivateTask(Id string) error {
	if _, err := MySQL.Exec(
		"update Crontab set IsActive=True where Id=?",
		t.Id,
	); err != nil {
		return err
	}
	return nil
}

func (t Task) DeactivateTask(Id string) error {
	if _, err := MySQL.Exec(
		"update Crontab set IsActive=False where Id=?",
		t.Id,
	); err != nil {
		return err
	}
	return nil
}

func (t Task) GetErrors(Id string) (int, error) {
	var error_count int
	if err := MySQL.QueryRow(
		"select Errors from Crontab where Id=?", Id,
	).Scan(
		&error_count,
	); err != nil {
		return -1, err
	}
	return error_count, nil
}

func (t Task) GetSuccess(Id string) (int, error) {
	return 0, nil
}

func (t Task) GetDateLastExecution(Id string) (time.Time, error) {
	return time.Time{}, nil
}
