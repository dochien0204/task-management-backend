package presenter

type StatusPresenter struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CategoryPresenter struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
