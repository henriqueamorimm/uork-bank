package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"uorkbank/controllers"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, fmt.Sprintf("Erro ao enviar o formulário: %v", err), http.StatusInternalServerError)
			return
		}

		name := r.PostForm.Get("name")
		email := r.PostForm.Get("email")
		senha := r.PostForm.Get("senha")
		saldoStr := r.PostForm.Get("saldo")
		saldo, err := strconv.Atoi(saldoStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao converter saldo para inteiro: %v", err), http.StatusInternalServerError)
			return
		}

		err = controllers.CreateUserSubscription(name, email, senha, saldo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao salvar os dados: %v", err), http.StatusInternalServerError)
			return
		}

		err = controllers.StoreAccountInLocalStorage(w, r, name, email, saldo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao armazenar a conta no Local Storage: %v", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/dashboard/index.html", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "handlers/home/index.html")
}
