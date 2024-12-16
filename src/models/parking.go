package models

type Parking struct {
	ID          string `json:"id" param:"id"`
	Status      string `json:"status"`
	Pubname     string `json:"pubname"`
	Description string `json:"description"`
	MapURL      string `json:"mapurl"`
	Position    string `json:"position"`
}

func GetParkings() *[]Parking {
	p := &[]Parking{}
	db.Find(p)
	return p
}

func FindParking(p *Parking) (*Parking, error) {
	result := &Parking{}
	if err := db.Where("id = ?", p.ID).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func AddParking(p *Parking) error {
	result := db.Create(p)
	return result.Error
}

func UpdateParking(p *Parking) error {
	result := db.Model(&Parking{}).Where("id = ?", p.ID).Update("status", p.Status)
	return result.Error
}

func DeleteParking(p *Parking) error {
	result := db.Where("id = ?", p.ID).Delete(p)
	return result.Error
}

func CreatePark() {
	// Insert into parking
	db.Create(&Parking{
		ID:          "tyuushajou1",
		Pubname:     "駐車場1",
		Description: "参考例",
		Status:      "0",
		MapURL:      "https://sankou",
		Position:    "[1.000, 1.000]",
	})
	db.Create(&Parking{
		ID:          "tyuushajou2",
		Pubname:     "駐車場2",
		Description: "参考例",
		Status:      "0",
		MapURL:      "https://sankou2",
		Position:    "[2.000, 2.000]",
	})
}
