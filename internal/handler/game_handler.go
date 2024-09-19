package handler

import (
	"context"
	"log"
	"net/http"

	"be/internal/models"
	"be/internal/utils"

	"github.com/gorilla/mux"
	"encoding/json"
)

func (h *MongoDBHandler) GetAllGames(w http.ResponseWriter, r *http.Request) {

	cursor, err := h.service.GetAllGames(r.Context())

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

	var games []models.Game

	for cursor.Next(context.TODO()) {
		var game models.Game

		err := cursor.Decode(&game)
		if err != nil {
			http.Error(w, "Error decoding games", http.StatusInternalServerError)
			return
		}

		games = append(games, game)

	}

	if err := cursor.Err(); err != nil {
		utils.WriteJsonResponse(w, http.StatusInternalServerError, nil, "Fail")
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, games, "Success")
}

func (h *MongoDBHandler) GetGameByGameId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	GameId := vars["game_id"]

	result := h.service.GetGameByGameId(r.Context(), GameId)

	if result == nil {
		utils.WriteJsonResponse(w, http.StatusNotFound, nil, "Fail")
		return
	}

	var game models.Game

	err := result.Decode(&game)

	if err != nil {
		utils.WriteJsonResponse(w, http.StatusInternalServerError, nil, "Fail")
		return
	}

	utils.WriteJsonResponse(w, http.StatusOK, game, "Success")
}

func (h *MongoDBHandler) UpdateGameByGameId(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    game_id := vars["game_id"]
    if game_id == "" {
        http.Error(w, "Missing game_id in URL", http.StatusBadRequest)
        return
    }

    contentType := r.Header.Get("Content-Type")
    if contentType != "application/json" {
        http.Error(w, "Invalid Content-Type, expected application/json", http.StatusUnsupportedMediaType)
        return
    }

    var updates map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    result, err := h.service.UpdateGameByGameId(r.Context(), game_id, updates)
    if err != nil {
        log.Printf("Error updating game: %v", err)
        http.Error(w, "Failed to update game", http.StatusInternalServerError)
        return
    }

    log.Printf("Matched %d documents and modified %d documents", result.MatchedCount, result.ModifiedCount)

    if result.UpsertedCount > 0 {
        log.Printf("Upserted document with ID: %v", result.UpsertedID)
    }

    utils.WriteJsonResponse(w, http.StatusOK, nil, "Success")
}
