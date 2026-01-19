package main

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/bootdotdev/learn-cicd-starter/internal/database"
    "github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
    type parameters struct {
        Name string `json:"name"`
    }

    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
        return
    }

    // Generate API key for the new user
    apiKey, err := generateRandomSHA256Hash()
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't generate API key", err)
        return
    }

    // Prepare user data
    userID := uuid.New().String()
    createdAt := time.Now().UTC().Format(time.RFC3339)
    updatedAt := time.Now().UTC().Format(time.RFC3339)

    // Create user in database
    err = cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
    ID:        userID,
    CreatedAt: createdAt,
    UpdatedAt: updatedAt,
    Name:      params.Name,
    ApiKey:    apiKey,
})
if err != nil {
    log.Printf("CreateUser ERROR details: %v", err)
    log.Printf("CreateUser params: ID=%s, Name=%s, ApiKey=%s", userID, params.Name, apiKey)
    respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
    return
}

    // Return the created user data (no need to fetch from database)
    respondWithJSON(w, http.StatusCreated, map[string]interface{}{
        "id":         userID,
        "created_at": createdAt,
        "updated_at": updatedAt,
        "name":       params.Name,
        "api_key":    apiKey,
    })
}

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
    userResp, err := databaseUserToUser(user)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't convert user", err)
        return
    }

    respondWithJSON(w, http.StatusOK, userResp)
}

func generateRandomSHA256Hash() (string, error) {
    randomBytes := make([]byte, 32)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }
    hash := sha256.Sum256(randomBytes)
    hashString := hex.EncodeToString(hash[:])
    return hashString, nil
}
