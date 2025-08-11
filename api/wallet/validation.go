package wallet

import "github.com/google/uuid"

func createValidation(req createWalletReq) map[string][]string {
	errValidation := make(map[string][]string, 4)

	_, err := uuid.Parse(req.UserID)
	if err != nil {
		errValidation["userID"] = append(errValidation["userID"], "must be uuid")
	}

	if req.Name == "" {
		errValidation["name"] = append(errValidation["name"], "cannot empty")
	}

	if req.Currency == "" {
		errValidation["currency"] = append(errValidation["currency"], "cannot empty")
	}

	return errValidation
}
