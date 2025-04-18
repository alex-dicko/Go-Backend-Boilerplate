package dto

// Used to return User Data in JSON response
// Makes sure everything is consistent
// Create more of these for models that will be returned in json
type UserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
