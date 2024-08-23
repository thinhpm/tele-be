package handler

import (
	"context"
	"log"
	"net/http"

	"be/internal/models"
	"be/internal/utils"
)

func (h *MongoDBHandler) GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("active")

	cursor, err := h.service.GetAllSchedules(r.Context(), active)

	if err != nil {
		http.Error(w, "Failed to fetch accounts", http.StatusInternalServerError)
		return
	}

	if cursor == nil {
		log.Fatal("Cursor is nil")
	}

	if !cursor.Next(context.TODO()) {
		log.Println("No documents found or cursor is empty")
	}

	defer cursor.Close(context.TODO())

	var schedules []models.Schedule

	for cursor.Next(context.TODO()) {
		var schedule models.Schedule

		err := cursor.Decode(&schedule)
		if err != nil {
			http.Error(w, "Error decoding games", http.StatusInternalServerError)
			return
		}

		schedules = append(schedules, schedule)

	}

	if err := cursor.Err(); err != nil {
		utils.WriteJsonResponse(w, http.StatusInternalServerError, nil, "Fail")
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, schedules, "Success")
}
