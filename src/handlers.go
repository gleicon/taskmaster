// Copyright 2013 taskmaster authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"regexp"
)

/* TaskCreateHandler creates new tasks. It accepts the following arguments:
- cron:		Cron string to schedule the task
- src[url]:	URL that's hit when cron executes
- src[method]:	HTTP method used for accessing the source URL
- src[body]:	Optional POST payload for the source URL
- src[policy]:	Policy for accessing the source URL
- dst[url]:	URL that'll receive the task response
- dst[policy]:	Policy for accessing the destination URL

Supported policies are:
- once:		Try once and give up if it fails
- persist:	Retry for ever until it works
- retry,N:	Use N as the number of times to retry before giving up
*/
func TaskCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, http.StatusText(405), 405)
		return
	}
	t := &Task{
		Cron: r.FormValue("cron"),
		Src: TaskSrc{
			URL:    r.FormValue("src[url]"),
			Method: r.FormValue("src[method]"),
			Body:   r.FormValue("src[body]"),
			Policy: r.FormValue("src[policy]"),
		},
		Dst: TaskDst{
			URL:    r.FormValue("dst[url]"),
			Policy: r.FormValue("dst[policy]"),
		},
	}
	fmt.Println(t)
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
