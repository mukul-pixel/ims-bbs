package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/mukul-pixel/ims-bbs/cmd/types"

)

type mockUserHandler struct{}

// CreateUser implements types.AdminStore.
func (m *mockUserHandler) CreateUser(user types.Admin) error {
	return nil
}

// GetUserByEmail implements types.AdminStore.
func (m *mockUserHandler) GetUserByEmail(email string) (*types.Admin, error) {
	return nil,fmt.Errorf("user not found with email:%s",email)
}

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserHandler{}
	userHandler := NewHandler(userStore)

	t.Run("should fail the api there and pass here", func(t *testing.T) {
		payload := types.AdminPayload{
			FirstName:   "Mukul",
			LastName:    "Khatri",
			Email:       "mukul@gmail.com",
			Password:    "hello123",
			Contact:     "9414269086",
			Address:     "laxmipura",
			Age:         21,
			JoiningDate: "24-07-2024",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/addAdmin", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/addAdmin", userHandler.handleAdmin)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest{
			t.Errorf("expected status code %d, got %d",http.StatusBadRequest,rr.Code)
		}
	})

}
