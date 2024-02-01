package dto

type DriverResponseDTO struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateDriverDTO struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	LicenseNumber string `json:"license_number"`
}

type UpdateDriverDTO struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	LicenseNumber string `json:"license_number"`
}
