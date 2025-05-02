package pet

import (
	"encoding/json"
	"io"
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

// Get godoc
// @Summary      Получить pet
// @Description  Возвращает сохранённый pet в формате JSON
// @Tags         pet
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      204  {string}  string  "No Content"
// @Router       /v1/pet [get]
func (handler *PetHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile(handler.filePath)
		if err != nil {
			http.Error(w, "", http.StatusNoContent)
			log.Printf("GET error: failed to read pet file: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Println("method GET success")
	}

}

// Put godoc
// @Summary      Обновить pet
// @Description  Загружает новый pet в формате JSON. Валидируются поля ascii и description
// @Tags         pet
// @Accept       json
// @Produce      plain
// @Param        pet  body      Pet  true  "Pet данные"
// @Success      200  {string}  string  "pet updated successfully"
// @Failure      400  {string}  string  "Bad Request"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /v1/pet [put]
func (handler *PetHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Читаем тело запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			log.Printf("PUT error: failed to read request body: %s", err)
			return
		}
		defer r.Body.Close()

		// Проверка валидного JSON-объекта
		var raw map[string]interface{}
		if err := json.Unmarshal(body, &raw); err != nil {
			http.Error(w, "invalid JSON format", http.StatusBadRequest)
			log.Printf("PUT error: invalid JSON format: %s", err)
			return
		}

		// Проверка что в JSON ровно 2 ключа (по сути на фронте только 2 поля, но лучше валидировать поля)
		if len(raw) != 2 {
			http.Error(w, "only 'ascii' and 'description' fields are allowed", http.StatusBadRequest)
			log.Println("PUT error: only 'ascii' and 'description' fields are allowed")
			return
		}

		// Парсим JSON в структуру Pet
		var pet Pet
		if err := json.Unmarshal(body, &pet); err != nil {
			http.Error(w, "invalid JSON format", http.StatusBadRequest)
			log.Printf("PUT error: invalid JSON format: %s", err)
			return
		}

		// Валидация на пустые поля
		if pet.Ascii == "" || pet.Description == "" {
			http.Error(w, "ascii and description fields are required", http.StatusBadRequest)
			log.Println("PUT error: ascii and description fields are required")
			return
		}

		// Удаляем старый файл, если он есть
		if _, err := os.Stat(handler.filePath); err == nil {
			if err := os.Remove(handler.filePath); err != nil {
				http.Error(w, "failed to remove pet file", http.StatusInternalServerError)
				log.Printf("PUT error: failed to remove pet file: %s", err)
				return
			}
		}

		// Пишем и сохраняем новый файл
		if err := os.WriteFile(handler.filePath, body, 0644); err != nil {
			http.Error(w, "failed to write pet file", http.StatusInternalServerError)
			log.Println("PUT error: failed to write pet file", err)
			return
		}

		// Возвращаем успешный ответ
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pet updated successfully"))
		log.Println("method PUT success")
	}
}

// Delete godoc
// @Summary      Удалить pet
// @Description  Удаляет pet файл
// @Tags         pet
// @Produce      plain
// @Success      204  {string}  string  "No Content"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /v1/pet [delete]
func (handler *PetHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(handler.filePath); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNoContent)
			log.Println("method PUT success")
			return
		}

		if err := os.Remove(handler.filePath); err != nil {
			http.Error(w, "failed to delete pet", http.StatusInternalServerError)
			log.Println("PUT error: failed to delete pet", err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		log.Println("method DELETE success")
	}
}
