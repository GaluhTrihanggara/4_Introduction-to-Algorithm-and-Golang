package dto

// UserCreateRequestDTO adalah DTO untuk data input pembuatan pengguna (register user)
type UserCreateRequest struct {
	Email    string `json:"email" validate:"required,email"`    // Contoh: "user@example.com"
	Password string `json:"password" validate:"required,min=8"` // Contoh: "password123"
}

// UserUpdateRequestDTO adalah DTO untuk data input pembaruan pengguna (update user)
type UserUpdateRequest struct {
	ID       uint   `json:"id" validate:"required"`              // Contoh: 1
	Email    string `json:"email" validate:"omitempty,email"`    // Contoh: "user@example.com"
	Password string `json:"password" validate:"omitempty,min=8"` // Contoh: "password123"
}

// UserResponseDTO adalah DTO untuk data output pengguna (response user)
type UserResponse struct {
	ID    uint   `json:"id"`    // Contoh: 1
	Email string `json:"email"` // Contoh: "user@example.com"
}
