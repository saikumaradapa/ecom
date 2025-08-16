package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/saikumaradapa/ecom/types"
)

func TestRegisterHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should return 400 Bad Request for invalid user payload", func(t *testing.T) {
		invalidPayload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@", // invalid email
			Password:  "strongpass",
		}

		requestBody, _ := json.Marshal(invalidPayload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, recorder.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
