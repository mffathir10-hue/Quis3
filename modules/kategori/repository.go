package kategori

import (
	"database/sql"
)

type Repository interface {
	CreateKategoriRepository(kategori Kategori) (result []Kategori, err error)
	GetAllKategoriRepository() (result []Kategori, err error)
	// GetBioskopRepository(id int) (kategori, error)
	// DeleteBioskopRepository(id int) (err error)
	// UpdateBioskopRepository(kategori kategori) (kategori, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateKategoriRepository(kategori Kategori) (result []Kategori, err error) {
	query := `
        INSERT INTO kategori (nama_kategori, description_kategori) 
        VALUES ($1, $2) 
        RETURNING id_kategori
    `

	var id int
	err = r.db.QueryRow(query, kategori.NAMA_KATEGORI, kategori.DESCRIPTION_KATEGORI).Scan(&id)
	if err != nil {
		return nil, err
	}

	kategori.ID_KATEGORI = id

	allkategori, err := r.GetAllKategoriRepository()
	if err != nil {
		return nil, err
	}

	return allkategori, nil
}

func (r *repository) GetAllKategoriRepository() ([]Kategori, error) {
	query := "SELECT id_kategori, nama_kategori, description_kategori FROM kategori ORDER BY id_kategori"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kategoris []Kategori
	for rows.Next() {
		var kategori Kategori
		err := rows.Scan(&kategori.ID_KATEGORI, &kategori.NAMA_KATEGORI, &kategori.DESCRIPTION_KATEGORI)
		if err != nil {
			return nil, err
		}
		kategoris = append(kategoris, kategori)
	}

	return kategoris, nil
}

// func (r *repository) GetBioskopRepository(id int) (kategori, error) {
// 	sql := "SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1"

// 	var kategori kategori
// 	err := r.db.QueryRow(sql, id).Scan(&kategori.ID, &kategori.Nama, &kategori.Lokasi, &kategori.Rating)

// 	if err != nil {
// 		return kategori{}, errors.New("gagal mengambil data bioskop: " + err.Error())
// 	}

// 	return kategori, nil
// }

// func (r *repository) UpdateBioskopRepository(kategori kategori) (kategori, error) {
// 	sql := "UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4 RETURNING id, nama, lokasi, rating"

// 	var updatedkategori kategori
// 	err := r.db.QueryRow(sql, kategori.Nama, kategori.Lokasi, kategori.Rating, kategori.ID).
// 		Scan(&updatedkategori.ID, &updatedkategori.Nama, &updatedkategori.Lokasi, &updatedkategori.Rating)

// 	if err != nil {
// 		return kategori{}, errors.New("gagal mengupdate bioskop: " + err.Error())
// 	}

// 	return updatedkategori, nil
// }

// func (r *repository) DeleteBioskopRepository(id int) (err error) {

// 	sql := "DELETE FROM bioskop WHERE id = $1"

// 	result, err := r.db.Exec(sql, id)
// 	if err != nil {
// 		return errors.New("gagal menghapus bioskop: " + err.Error())
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return errors.New("gagal memeriksa rows affected: " + err.Error())
// 	}

// 	if rowsAffected == 0 {
// 		return errors.New("bioskop tidak ditemukan")
// 	}

// 	return nil
// }
