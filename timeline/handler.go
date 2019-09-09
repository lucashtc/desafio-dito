package timeline

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// TimelineHand ...
func TimelineHand(w http.ResponseWriter, r *http.Request) {
	t, err := timelineEvents()
	if err != nil {
		errString := fmt.Sprint(errors.Wrap(err, "Falha"))
		http.Error(w, errString, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(t)
}
