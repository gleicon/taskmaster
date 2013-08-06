// Copyright 2013 taskmaster authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/* TaskCreateHandler creates new tasks. It accepts the following arguments:
- cron:		Cron string to schedule the task
- src[url]:	URL that's hit when cron executes [required]
- src[method]:	HTTP method used for accessing the source URL [default: GET]
- src[body]:	Optional POST payload for the source URL [default: empty]
- src[policy]:	Policy for accessing the source URL [default: once]
- dst[url]:	URL that'll receive the task response [required]
- dst[policy]:	Policy for accessing the destination URL [default: once]

Supported policies are:
- once:		Try once and give up if it fails [default]
- persist:	Retry for ever until it works
- retry,N:	Use N as the number of times to retry before giving up
*/
func TaskCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, http.StatusText(405), 405)
		return
	}
	var errmsg string
	cron := r.FormValue("cron")
	tsrc := TaskSrc{
		URL:    r.FormValue("src[url]"),
		Method: FormValueDefault(r, "src[method]", "GET"),
		Body:   r.FormValue("src[body]"),
		Policy: FormValueDefault(r, "src[policy]", "once"),
	}
	tdst := TaskDst{
		URL:    r.FormValue("dst[url]"),
		Policy: FormValueDefault(r, "dst[policy]", "once"),
	}
	switch {
	case cron == "":
		errmsg = "Missing crontabstring"
	case tsrc.URL == "":
		errmsg = "Missing src[url]"
	case !ValidMethod(tsrc.Method):
		errmsg = "Invalid src[method]"
	case !ValidPolicy(tsrc.Policy):
		errmsg = "Invalid src[policy]"
	case tdst.URL == "":
		errmsg = "Missing dst[url]"
	case !ValidPolicy(tdst.Policy):
		errmsg = "Invalid dst[policy]"
	}
	if errmsg != "" {
		http.Error(w, errmsg, 400)
		return
	}
	t, _ := NewTask(cron, tsrc, tdst)
	fmt.Println(t)
}

func ValidMethod(s string) bool {
	switch s {
	case "DELETE":
	case "GET":
	case "POST":
	case "PUT":
	default:
		return false
	}
	return true
}

func ValidPolicy(s string) bool {
	tmp := strings.SplitN(s, ",", 1)
	switch tmp[0] {
	case "once":
	case "persist":
	case "retry":
		if n, err := strconv.Atoi(tmp[1]); err != nil || n < 0 {
			return false
		}
	default:
		return false
	}
	return true
}

func FormValueDefault(r *http.Request, k, d string) string {
	if s := r.FormValue(k); s != "" {
		return s
	} else {
		return d
	}
}

func TaskListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, http.StatusText(405), 405)
		return
	}
}

func TaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, http.StatusText(405), 405)
		return
	}
	taskID := r.URL.Path[len("/api/v1/status/"):]
	fmt.Fprintln(w, "task id:", taskID)
}

func TaskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", "DELETE")
		http.Error(w, http.StatusText(405), 405)
		return
	}
}
