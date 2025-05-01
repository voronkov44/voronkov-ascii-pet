package pet

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PetHandler struct {
	filePath string
}

func NewPetHandler(router *http.ServeMux) *PetHandler {
	handler := &PetHandler{
		filePath: "pet.json",
	}
	router.HandleFunc("GET /v1/pet", handler.Get())
	router.HandleFunc("PUT /v1/pet", handler.Put())
	router.HandleFunc("DELETE /v1/pet", handler.Delete())
	return handler
}

func (handler *PetHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(handler.filePath)
		if err != nil {
			http.Error(w, "", http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Println("method GET success")
	}

}

func (handler *PetHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Читаем тело запроса
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Парсим JSON в структуру Pet
		var pet Pet
		if err := json.Unmarshal(body, &pet); err != nil {
			http.Error(w, "invalid JSON format", http.StatusBadRequest)
			return
		}

		// Валидация полей
		if pet.Ascii == "" || pet.Description == "" {
			http.Error(w, "ascii and description fields are required", http.StatusBadRequest)
			return
		}

		// Если файл существует - удаляем
		if _, err := os.Stat(handler.filePath); err == nil {
			if err := os.Remove(handler.filePath); err != nil {
				http.Error(w, "failed to remove pet file", http.StatusInternalServerError)
				return
			}
		}

		// Сохраняем новые данные в файл
		if err := ioutil.WriteFile(handler.filePath, body, 0644); err != nil {
			http.Error(w, "failed to write pet file", http.StatusInternalServerError)
			return
		}

		// Возвращаем успешный ответ
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pet updated successfully"))
		log.Println("method PUT success")
	}
}

func (handler *PetHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := os.Remove(handler.filePath); err != nil {
			http.Error(w, "failed to delete pet", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		log.Println("method DELETE success")
	}
}
