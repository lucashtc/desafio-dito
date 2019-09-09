package timeline

import (
	"encoding/json"
	"log"
	"net/http"
)

// TimelineHand ...
func TimelineHand(w http.ResponseWriter, r *http.Request) {
	t, err := timelineEvents()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(t)
}
