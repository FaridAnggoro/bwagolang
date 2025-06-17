package handler

import (
	"net/http"
	"startup/helper"
	"startup/user"

	"github.com/gin-gonic/gin"
)

// struct
type userHandler struct {
	// handler mempunyai dependency terhadap service
	userService user.Service
}

// function membuat handler baru
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// buat handler
func (h *userHandler) RegisterUser(c *gin.Context) {

	// ambil input dari user
	var input user.RegisterUserInput

	// ubah json ke RegisterUserInput
	// mapping input dari user ke struct RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		// validasi error
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		// validasi gagal
		response := helper.APIResponse("Akun gagal terdaftar", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	// struct di atas kita panggil ke service
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Akun gagal terdaftar", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.APIResponse("Akun berhasil terdaftar", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

	// service akan memanggil repository
	// repository akan memanggil db
}

func (h *userHandler) Login(c *gin.Context) {

	// ambil input dari user
	var input user.LoginInput

	// ubah json ke LoginInput
	// mapping input dari user ke struct LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		// validasi error
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		// validasi gagal
		response := helper.APIResponse("Akun gagal masuk", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		// validasi gagal
		response := helper.APIResponse("Akun gagal masuk", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Akun berhasil masuk", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
