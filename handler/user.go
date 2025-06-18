package handler

import (
	"fmt"
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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		// validasi error
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		// validasi gagal
		response := helper.APIResponse("C", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		// validasi gagal
		response := helper.APIResponse("Pengecekan email gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": false}
	response := helper.APIResponse("Foto Profil berhasil diupload", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
