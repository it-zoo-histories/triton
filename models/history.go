package models

/*History - модель истории*/
type History struct {
	ID    string `json:"id"`    // идентификатор истории
	Type  int    `json:"type"`  // тип истории
	Value string `json:"value"` // значение истории
}
