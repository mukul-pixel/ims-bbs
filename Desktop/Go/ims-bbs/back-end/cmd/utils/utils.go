package utils

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"

	"github.com/mukul-pixel/ims-bbs/cmd/types"

)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ReadFromFile(file multipart.File) ([]types.AdminPayload, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, err
	}

	var admins []types.AdminPayload

	for i, row := range rows {
		if i == 0 {
			continue //skipping heading row
		}
		if len(row) < 7 {
			continue
		}

		age, _ := strconv.Atoi(row[6])

		admin := types.AdminPayload{
			FirstName:   row[0],
			LastName:    row[1],
			Email:       row[2],
			Password:    row[3],
			Contact:     row[4],
			Address:     row[5],
			Age:         age,
			JoiningDate: row[7],
		}

		admins = append(admins, admin)
	}

	return admins,nil
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
