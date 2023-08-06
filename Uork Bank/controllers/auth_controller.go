package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Account struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Saldo int    `json:"saldo"`
}

func StoreAccountInLocalStorage(w http.ResponseWriter, r *http.Request, nome, email string, saldo int) error {
	account := Account{Nome: nome, Email: email, Saldo: saldo}

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	accountStr := string(accountJSON)

	jsCode := fmt.Sprintf(`<script>localStorage.setItem("accountData", %s);</script>`, strconv.Quote(accountStr))

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(jsCode))

	http.Redirect(w, r, "/dashboard/index.html", http.StatusSeeOther)

	return nil
}
