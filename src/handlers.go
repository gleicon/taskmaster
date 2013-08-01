// Copyright 2013 taskmaster authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
    "regexp"
)

var crontab_id = regexp.MustCompile(`^/([0-9]+)$`)

func SchedulerHandler(w http.ResponseWriter, r *http.Request) {
	bar, _ := Redis.Get("foo") // redis-cli set foo bar
	fmt.Fprintln(w, "Hello, world", bar)
    
    action := r.URL.Path[len("/api/scheduler/"):]
    fmt.Fprintln(w, "action: ", action)
    switch {
        case action == "new":
            fmt.Fprintln(w, "new");
        case action == "remove":
            fmt.Fprintln(w, "remove");
        case action == "list":
            fmt.Fprintln(w, "list");
        case crontab_id.MatchString(action):
            fmt.Fprintln(w, "id: %s", action);
        default:
            fmt.Fprintln(w, "index")
    }

}

