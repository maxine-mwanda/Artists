package models

type Artist struct {
	Id          int64  `json:"id"`
	StageName   string `json:"stage_name"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	Nationality string `json:"nationality"`
	CreatedAt   string `json:"created_at"`
}

type NewArtist struct {
	StageName   string `json:"stage_name"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	Nationality string `json:"nationality"`
}
