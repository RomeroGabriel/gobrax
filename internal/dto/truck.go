package dto

type TruckResponseDTO struct {
	Id           string `json:"id"`
	ModelType    string `json:"model_type"`
	Year         uint16 `json:"year"`
	Manufacturer string `json:"manufacturer"`
	LicensePlate string `json:"license_plate"`
	FuelType     string `json:"fuel_type"`
}

type CreateTruckDTO struct {
	ModelType    string `json:"model_type"`
	Year         uint16 `json:"year"`
	Manufacturer string `json:"manufacturer"`
	LicensePlate string `json:"license_plate"`
	FuelType     string `json:"fuel_type"`
}

type UpdateTruckDTO struct {
	Id           string `json:"id"`
	ModelType    string `json:"model_type"`
	LicensePlate string `json:"license_plate"`
	// TODO: Add Status
}
