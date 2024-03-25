package main

// func indexSearchSpaces(writer http.ResponseWriter, _ *http.Request) {
// 	json.NewEncoder(writer).Encode(searchSpaces)
// }

// func createSearchSpace(writer http.ResponseWriter, request *http.Request) {
// 	var newSpace SearchSpace
// 	_ = json.NewDecoder(request.Body).Decode(&newSpace)
// 	searchSpaces = append(searchSpaces, newSpace)
// 	json.NewEncoder(writer).Encode(newSpace)
// }

// func showSearchSpace(writer http.ResponseWriter, request *http.Request) {
// 	params := mux.Vars(request)
// 	for _, item := range searchSpaces {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(writer).Encode(item)
// 			return
// 		}
// 	}
// 	writer.WriteHeader(http.StatusNotFound)
// }

// func deleteSearchSpace(writer http.ResponseWriter, request *http.Request) {
// 	params := mux.Vars(request)
// 	for index, item := range searchSpaces {
// 		if item.ID == params["id"] {
// 			searchSpaces = append(searchSpaces[:index], searchSpaces[index+1:]...)
// 			break
// 		}
// 	}
// 	writer.WriteHeader(http.StatusOK)
// }
