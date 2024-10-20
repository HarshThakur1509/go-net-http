package controllers

import (
	"encoding/json"
	"net/http"
	"test/initializers"
	"test/models"
)

func PostIdea(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string
	}
	json.NewDecoder(r.Body).Decode(&body)

	idea := models.Idea{Title: body.Title}
	result := initializers.DB.Create(&idea)

	if result.Error != nil {
		http.Error(w, "Something went wrong!!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(idea)

}

func GetIdeas(w http.ResponseWriter, r *http.Request) {
	var ideas []models.Idea
	initializers.DB.Find(&ideas)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ideas)
}

func GetIdeaIndex(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var idea models.Idea
	initializers.DB.First(&idea, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(idea)
}

func UpdateIdea(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var idea models.Idea
	initializers.DB.First(&idea, id)

	var body struct {
		Title string
	}
	json.NewDecoder(r.Body).Decode(&body)
	newIdea := models.Idea{Title: body.Title}

	result := initializers.DB.Model(&idea).Updates(&newIdea)

	if result.Error != nil {
		http.Error(w, "Something went wrong!!", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 201 Created
	json.NewEncoder(w).Encode(idea)

}

func DeleteIdea(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var idea models.Idea
	initializers.DB.Delete(&idea, id)
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
