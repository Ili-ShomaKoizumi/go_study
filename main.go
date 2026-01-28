package main

import (
	"gin/service"
	"gin/utils"

	"gin/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := utils.NewDBConnection()

	reviewRepository := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepository)
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, reviewRepository)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	purHistoryRepository := repository.NewPurHistoryRepository(db)
	purHistoryService := service.NewPurHistoryService(purHistoryRepository)

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/products", productService.RegisterProduct)                         // 商品出品(登録)
	router.PUT("/products/:id", productService.UpdateProduct)                        // 商品情報編集(更新)
	router.DELETE("/products/:id", productService.DeleteProduct)                     // 商品削除
	router.GET("/products", productService.GetProducts)                              // 商品一覧取得
	router.GET("/products/:id", productService.GetProductsById)                      // 商品情報取得
	router.POST("/users", userService.RegisterUser)                                  // ユーザー登録
	router.PUT("/users/:id", userService.UpdateUser)                                 // ユーザー情報編集(更新)
	router.DELETE("/users/:id", userService.DeleteUser)                              // ユーザー削除
	router.GET("/users", userService.GetUsers)                                       // ユーザー一覧取得
	router.GET("/users/:id", userService.GetUsersById)                               // ユーザー情報取得
	router.POST("/review", reviewService.RegisterReview)                             // レビュー投稿
	router.POST("/purhistory/:id", purHistoryService.RegisterPurHistory)             // 購入履歴登録
	router.GET("/purhistory/:id", purHistoryService.GetPurHistorysById)              // ユーザー購入履歴取得
	router.PUT("/purhistory/:userId/*productId", purHistoryService.UpdatePurHistory) // 購入履歴更新

	router.Run(":8080")
}
