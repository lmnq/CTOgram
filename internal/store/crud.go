package store

import (
	"ctogram/internal/models"
)

// GetCityByID ..
func (s *Store) GetCityByID(id int) (models.City, error) {
	city := models.City{}
	row := s.DB.QueryRow(`SELECT * FROM cities WHERE ID=?`, id)
	err := row.Scan(&city.ID, &city.Name, &city.Code, &city.CountryCode)
	return city, err
}

// GetCities ..
func (s *Store) GetCities() ([]models.City, error) {
	cities := []models.City{}
	rows, err := s.DB.Query(`SELECT * FROM cities`)
	if err != nil {
		return cities, err
	}
	for rows.Next() {
		city := models.City{}
		err := rows.Scan(&city.ID, &city.Name, &city.Code, &city.CountryCode)
		if err != nil {
			return cities, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}

// AddCity ..
func (s *Store) AddCity(city models.City) error {
	_, err := s.DB.Exec(`
			INSERT INTO cities (name, code, country_code)
			VALUES (?, ?, ?);
	`, city.Name, city.Code, city.CountryCode)

	return err
}

// DeleteCityByID ..
func (s *Store) DeleteCityByID(id int) error {
	tx, err := s.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	row := s.DB.QueryRow(`SELECT ID FROM cities WHERE ID=?`, id)
	var existingID int
	if err := row.Scan(&existingID); err != nil {
		tx.Rollback()
		return err
	}
	_, err = s.DB.Exec(`DELETE FROM cities WHERE ID=?`, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// UpdateCity ..
func (s *Store) UpdateCity(id int, city models.City) error {
	_, err := s.DB.Exec(`
		UPDATE cities 
		SET name=?, code=?, country_code=? WHERE ID=?`,
		city.Name, city.Code, city.CountryCode, id)
	return err
}
