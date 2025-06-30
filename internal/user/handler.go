package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/M-kos/crm-user/internal/config"
	"github.com/M-kos/crm-user/internal/logger"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	config  *config.Config
	logger  *logger.Logger
	service *UserService
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		h.logger.Error("get all user err: ", err)
		http.Error(w, "[GetAllUsers] something went wrong", http.StatusInternalServerError) // TODO Возможно изменить ответ
		return
	}

	w.Header().Set("Cotent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest CreateUserRequest
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&createUserRequest); err != nil {
		h.logger.Error("[CreateUser] decode new user err: ", err)
		http.Error(w, "something went wrong", http.StatusBadRequest) // TODO Возможно изменить ответ
		return
	}

	if err := createUserRequest.Validate(); err != nil {
		h.logger.Error("[CreateUser] validate new user err: ", err)
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}

	newUser := NewUser(&UserFromRequest{
		FirstName:   createUserRequest.FirstName,
		LastName:    createUserRequest.LastName,
		Email:       createUserRequest.Email,
		Password:    createUserRequest.Password,
		Permissions: createUserRequest.Permissions,
		UserType:    createUserRequest.Type,
		Phone:       createUserRequest.Phone,
		Tg:          createUserRequest.TgName,
	})

	if err := h.service.Create(r.Context(), newUser); err != nil {
		h.logger.Error("[CreateUser] create new user err: ", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError) // TODO Возможно изменить ответ
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	validate := validator.New()

	if err := validate.Var(userId, "required,uuid"); err != nil {
		h.logger.Error("[GetUserById] user id validation err: ", err)
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetById(r.Context(), userId)
	if err != nil {
		h.logger.Error("[GetUserById] get user by id err: ", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError) // TODO Возможно изменить ответ
		return
	}

	w.Header().Set("Cotent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetUserByIdResponse{
		UUID:        result.UUID,
		Permissions: result.Permissions,
		Type:        result.UserType,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Phone:       result.Phone,
		TgName:      result.Tg,
		Status:      string(result.Status),
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	})
}

func (h *UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	validate := validator.New()

	if err := validate.Var(userId, "required,uuid"); err != nil {
		h.logger.Error("[UpdateUserById] user id validation err: ", err)
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var updateUserRequest UpdateUserRequest

	d := json.NewDecoder(r.Body)
	if err := d.Decode(&updateUserRequest); err != nil {
		h.logger.Error("[UpdateUserById] decode user err: ", err)
		http.Error(w, "something went wrong", http.StatusBadRequest) // TODO Возможно изменить ответ
		return
	}

	if err := updateUserRequest.Validate(); err != nil {
		h.logger.Error("[UpdateUserById] validate user err: ", err)
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}

	updatableUser := NewUser(&UserFromRequest{
		FirstName:   updateUserRequest.FirstName,
		LastName:    updateUserRequest.LastName,
		Email:       updateUserRequest.Email,
		Password:    updateUserRequest.Password,
		Permissions: updateUserRequest.Permissions,
		UserType:    updateUserRequest.Type,
		Phone:       updateUserRequest.Phone,
		Tg:          updateUserRequest.TgName,
	})

	if err := h.service.Update(r.Context(), updatableUser); err != nil {
		h.logger.Error("[UpdateUserById] update user err: ", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	validate := validator.New()

	if err := validate.Var(userId, "required,uuid"); err != nil {
		h.logger.Error("[DeleteUserById] user id validation err: ", err)
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), userId); err != nil {
		h.logger.Error("[DeleteUserById] delete user err: ", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewUserHandler(router *http.ServeMux, config *config.Config, logger *logger.Logger, service *UserService) {
	place := "[UserHandler]"

	logger.Log = logger.Log.With(slog.String("place", place))

	handler := &UserHandler{
		config:  config,
		logger:  logger,
		service: service,
	}

	router.HandleFunc("GET /users", handler.GetAllUsers)
	router.HandleFunc("POST /users", handler.CreateUser)
	router.HandleFunc("GET /users/{userId}", handler.GetUserById)
	router.HandleFunc("PUT /users/{userId}", handler.UpdateUserById)
	router.HandleFunc("DELETE /users/{userId}", handler.DeleteUserById)
}
