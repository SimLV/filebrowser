package http

import (
	"net/http"

	"github.com/gtsteffaniak/filebrowser/search"
)

var searchHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	response := []map[string]interface{}{}
	query := r.URL.Query().Get("query")
	indexInfo, fileTypes := search.SearchAllIndexes(query, r.URL.Path)
	for _, path := range indexInfo {
		responseObj := map[string]interface{}{
			"path": path,
			"dir":  true,
		}
		if _, ok := fileTypes[path]; ok {
			responseObj["dir"] = false
			for filterType, value := range fileTypes[path] {
				if value {
					responseObj[filterType] = value
				}
			}
		}
		response = append(response, responseObj)
	}
	return renderJSON(w, r, response)
})
