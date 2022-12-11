package essayApi

import (
	"mark/entity"
	"mark/models"
	"encoding/json"
	"net/http"
)

func FindEssay(response http.ResponseWriter, request *http.Request){
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1{
		responseWithError(response, http.StatusBadRequest, "missing param id")
		return
	} else {
		essay, err := models.FindEssay(ids[0])
		if err != nil {
			responseWithError(response, http.StatusBadRequest, "essay not found")
			return
		}
		responseWithJson(response, http.StatusOK, essay)
	}
}

func GetAll(response http.ResponseWriter, request *http.Request){
	listEssay := models.GetAllEssay()
	responseWithJson(response, http.StatusOK, listEssay)
}

func CreateEssay(response http.ResponseWriter, request *http.Request){
	var essay entity.Essay
	err := json.NewDecoder(request.Body).Decode(&essay)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		result := models.CreateEssay(&essay)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not create essay")
			return
		}
		responseWithJson(response, http.StatusOK, essay)
		return
	}
}

func UpdateEssay(response http.ResponseWriter, request *http.Request){
	var essay entity.Essay
	err := json.NewDecoder(request.Body).Decode(&essay)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		result := models.UpdateEssay(&essay)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not update essay")
			return
		}
		responseWithJson(response, http.StatusOK, "Update essay successful")
		return
	}
}

func DeleteEssay(response http.ResponseWriter, request *http.Request){
	ids, ok:= request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "URL param is missing")
		return
	}
	result := models.DeleteEssay(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Could not delete essay")
		return
	}
	responseWithJson(response, http.StatusOK, "Delete essay successfully")
}


func responseWithError(response http.ResponseWriter, statusCode int, msg string){
	responseWithJson(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJson(response http.ResponseWriter, statusCode int, data any){
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}