// Copyright 2013 taskmaster authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// +build db

// These are just examples and are not currently used by the server.

package main

import "database/sql"

type User struct {
	Id       int
	Email    string
	Passwd   string
	FullName sql.NullString
	IsActive bool
}

func NewUser(email, passwd, fullname string, active bool) (*User, error) {
	if _, err := MySQL.Exec(`
		insert into User (Email, Passwd, FullName, IsActive)
		values (?, SHA1(?), ?, ?)`,
		email,
		passwd,
		fullname,
		active,
	); err != nil {
		return nil, err
	}
	return GetUser(email)
}

func UserExists(email string) (bool, error) {
	var count int
	if err := MySQL.QueryRow(
		"select count(*) from User where Email=?", email,
	).Scan(
		&count,
	); err != nil {
		return false, err
	}
	return count >= 1, nil

}

// TODO: cache
func GetUser(email string) (*User, error) {
	var u User
	if err := MySQL.QueryRow(
		"select * from User where Email=?", email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.Passwd,
		&u.FullName,
		&u.IsActive,
	); err != nil {
		return nil, err
	}
	return &u, nil
}

// TODO: cache
func GetUserById(id int) (*User, error) {
	var u User
	if err := MySQL.QueryRow(
		"select * from User where Id=?", id,
	).Scan(
		&u.Id,
		&u.Email,
		&u.Passwd,
		&u.FullName,
		&u.IsActive,
	); err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserWithPasswd(email, passwd string) (*User, error) {
	var u User
	if err := MySQL.QueryRow(
		"select * from User where Email=? and Passwd=SHA1(?)",
		email,
		passwd,
	).Scan(
		&u.Id,
		&u.Email,
		&u.Passwd,
		&u.FullName,
		&u.IsActive,
	); err != nil {
		return nil, err
	}
	return &u, nil
}

func DelUser(u *User) error {
	if _, err := MySQL.Exec(
		"delete from Users where Id=?", u.Id,
	); err != nil {
		return err
	}
	return nil
}

func UpdateUser(u *User) error {
	if _, err := MySQL.Exec(
		"update User set Passwd=?, FullName=?, IsActive=? where Id=?",
		u.Passwd,
		u.FullName.String,
		u.IsActive,
		u.Id,
	); err != nil {
		return err
	}
	return nil
}

// Tasks

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
	TaskId            int
	CrontabString     string
	DateCreation      sql.Timestamp
	DateLastModified  sql.Timestamp
	DateLastExecution sql.Timestamp
	Errors            int
	Success           int
	IsActive          bool
	Src               TaskSrc
	Dst               TaskDst
}

// Model
func NewTask(CrontabString string, taskSrc TaskSrc, taskDst TaskDst) (id, error) {}
func FindTaskById(Id int) (*Task, error)                                         {}
func (t Task) ActivateTask(Id string) (bool, error)                              {}
func (t Task) DeactivateTask(Id string) (bool, error)                            {}
func (t Task) GetErrors(Id string) (int, error)                                  {}
func (t Task) GetSuccess(Id string) (int, error)                                 {}
func (t Task) GetDateLastExecution(Id string) (sql.Timestamp, error)             {}
