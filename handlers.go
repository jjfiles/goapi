package main

import (
	"encoding/json"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}