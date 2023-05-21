package category

import (
	"fmt"
	"learn-web_crud/config"
	"learn-web_crud/entitas"
	"time"
)

func GetAll() []entitas.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entitas.Category

	for rows.Next() {
		var category entitas.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		// Menghitung selisih waktu dengan waktu saat ini
		duration := time.Since(category.CreatedAt)
		// Mengubah selisih waktu menjadi format "2 jam yang lalu"
		formattedCreatedAt := formatDuration(duration)
		category.CreatedAtFormatted = formattedCreatedAt

		categories = append(categories, category)
	}

	return categories
}

// function Create
func Create(category entitas.Category) bool {
	result, err := config.DB.Exec(
		`INSERT INTO categories (name, created_at, updated_at) VALUE (?, ?, ?)`,
		category.Name, category.CreatedAt, category.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastId > 0
}

// func edit
func Detail(id int) entitas.Category {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)

	var category entitas.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		panic(err.Error())
	}

	return category
}

// func update
func Update(id int, category entitas.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

// mengubah format waktu
func formatDuration(durasi time.Duration) string {
	if durasi.Hours() >= 24 {
		hari := int(durasi.Hours() / 24)
		return fmt.Sprintf("%d hari yang lalu", hari)
	}

	if durasi.Hours() >= 1 {
		jam := int(durasi.Hours())
		return fmt.Sprintf("%d jam yang lalu", jam)
	}

	if durasi.Minutes() >= 1 {
		menit := int(durasi.Minutes())
		return fmt.Sprintf("%d menit yang lalu", menit)
	}

	return "Baru saja"
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}
