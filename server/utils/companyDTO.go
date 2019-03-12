package utils

type CompanyDTO struct {
	Status         string `json:"status"`
	Name           string `json:"name"`
	PhotoPath      string `json:"photo_path"`
	Phone          string `json:"phone"`
	Site           string `json:"site"`
	Email          string `json:"email"`
	Info           string `json:"info"`
	FullInfoPath   string `json:"full_info_path"`
	RequisitesPath string `json:"requisites_path"`
}
