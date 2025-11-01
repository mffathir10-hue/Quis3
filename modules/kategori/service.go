package kategori

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateKategoriService(ctx *gin.Context) (Kategori, error)
	GetAllKategoriService(ctx *gin.Context) (result []Kategori, err error)
	// GetBioskopService(ctx *gin.Context) (kategori, error)
	// UpdateBioskopService(ctx *gin.Context) (kategori, error)
	// DeleteBioskopService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (s *userService) CreateKategoriService(ctx *gin.Context) (Kategori, error) {
	var kategori Kategori

	if err := ctx.ShouldBindJSON(&kategori); err != nil {
		return kategori, err
	}

	if kategori.NAMA_KATEGORI == "" {
		return kategori, errors.New("nama kategori harus diisi")
	}

	if kategori.DESCRIPTION_KATEGORI == "" {
		return kategori, errors.New("deskripsi kategori harus diisi")
	}

	result, err := s.repository.CreateKategoriRepository(kategori)
	if err != nil {
		return kategori, errors.New("gagal menambahkan kategori: " + err.Error())
	}

	return result[kategori.ID_KATEGORI], nil
}

func (s *userService) GetAllKategoriService(ctx *gin.Context) ([]Kategori, error) {
	kategoris, err := s.repository.GetAllKategoriRepository()
	if err != nil {
		return nil, errors.New("gagal mengambil data kategori: " + err.Error())
	}

	return kategoris, nil
}

// func (s *userService) GetBioskopService(ctx *gin.Context) (result kategori, err error) {
// 	idStr := ctx.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return kategori{}, errors.New("ID bioskop tidak valid")
// 	}

// 	if id <= 0 {
// 		return kategori{}, errors.New("ID bioskop harus lebih dari 0")
// 	}

// 	result, err = s.repository.GetBioskopRepository(id)
// 	if err != nil {
// 		return kategori{}, errors.New("bioskop tidak ditemukan")
// 	}

// 	return result, nil
// }

// func (s *userService) UpdateBioskopService(ctx *gin.Context) (result kategori, err error) {

// 	idStr := ctx.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return kategori{}, errors.New("ID bioskop tidak valid")
// 	}

// 	if id <= 0 {
// 		return kategori{}, errors.New("ID bioskop harus lebih dari 0")
// 	}

// 	var kategori kategori
// 	if err := ctx.ShouldBindJSON(&kategori); err != nil {
// 		return kategori{}, errors.New("data request tidak valid")
// 	}

// 	kategori.ID = id

// 	if strings.TrimSpace(kategori.Nama) == "" {
// 		return kategori{}, errors.New("nama bioskop harus diisi")
// 	}

// 	if strings.TrimSpace(kategori.Lokasi) == "" {
// 		return kategori{}, errors.New("lokasi bioskop harus diisi")
// 	}

// 	result, err = s.repository.UpdateBioskopRepository(kategori)
// 	if err != nil {
// 		return kategori{}, err
// 	}

// 	return result, nil
// }

// func (s *userService) DeleteBioskopService(ctx *gin.Context) (err error) {
// 	idStr := ctx.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		return errors.New("ID bioskop tidak valid")
// 	}

// 	if id <= 0 {
// 		return errors.New("ID bioskop harus lebih dari 0")
// 	}

// 	err = s.repository.DeleteBioskopRepository(id)
// 	if err != nil {
// 		return err
// 	}

// 	return
// }
