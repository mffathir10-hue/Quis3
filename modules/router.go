package modules

import (
	"database/sql"
	"fmt"
	"gintugas/modules/buku"
	"gintugas/modules/kategori"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api")
	{
		api.POST("/kategori", CreateKategoriRouter(db))
		api.GET("/kategori", GetAllKategoriRouter(db))
		api.POST("/buku", CreateBukuRouter(db))
		api.GET("/buku", GetAllBukuRouter(db))
		// api.GET("/buku/:id", GetBukuRouter(db))
		// api.PUT("/buku/:id", UpdateBukuRouter(db))
		// api.DELETE("/buku/:id", DeleteBukuRouter(db))
	}
}

func CreateKategoriRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			kategoriRepo = kategori.NewRepository(db)
			kategoriSrv  = kategori.NewService(kategoriRepo)
		)

		kategoris, err := kategoriSrv.CreateKategoriService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message":  fmt.Sprintf("Data kategori berhasil ditambahkan"),
			"Kategori": kategoris,
		})
	}
}

func GetAllKategoriRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			kategoriRepo = kategori.NewRepository(db)
			kategoriSrv  = kategori.NewService(kategoriRepo)
		)

		kategoris, err := kategoriSrv.GetAllKategoriService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  fmt.Sprintf("successfully get all kategori data"),
			"Kategori": kategoris,
		})
	}
}

func CreateBukuRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			bukuRepo = buku.NewRepository(db)
			bukuSrv  = buku.NewService(bukuRepo)
		)

		bukus, err := bukuSrv.CreateBukuService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": fmt.Sprintf("Data buku berhasil ditambahkan"),
			"Buku":    bukus,
		})
	}
}

func GetAllBukuRouter(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			bukuRepo = buku.NewRepository(db)
			bukuSrv  = buku.NewService(bukuRepo)
		)

		bukus, err := bukuSrv.GetAllBukuService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully get all buku data"),
			"Buku":    bukus,
		})
	}
}

// func GetBukuRouter(db *sql.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var (
// 			bioskopRepo = NewRepository(db)
// 			bioskopSrv  = NewService(bioskopRepo)
// 		)

// 		bioskop, err := bioskopSrv.GetBioskopService(ctx)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": fmt.Sprintf("successfully get bioskop data"),
// 			"Bioskop": bioskop,
// 		})
// 	}
// }

// func UpdateBukuRouter(db *sql.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var (
// 			bioskopRepo = NewRepository(db)
// 			bioskopSrv  = NewService(bioskopRepo)
// 		)

// 		bioskop, err := bioskopSrv.UpdateBioskopService(ctx)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": fmt.Sprintf("Data Bioskop Berhasil di Update"),
// 			"Bioskop": bioskop,
// 		})
// 	}
// }

// func DeleteBukuRouter(db *sql.DB) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var (
// 			bioskopRepo = NewRepository(db)
// 			bioskopSrv  = NewService(bioskopRepo)
// 		)

// 		err := bioskopSrv.DeleteBioskopService(ctx)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": "Bioskop berhasil dihapus",
// 		})
// 	}
// }
