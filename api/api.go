package api

import (
	"api-geteway/api/middleware"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-geteway/api/docs"
	"api-geteway/api/handler"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	h := handler.NewHandler(conn)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/", middleware.JWTMiddleware())

	// Author routes
	author := protected.Group("/authors")
	author.POST("/", h.CreateAuthor)
	author.GET("/:id", h.GetAuthor)
	author.GET("/", h.GetAllAuthor)
	author.PUT("/:id", h.UpdateAuthor)
	author.DELETE("/:id", h.DeleteAuthor)

	// Book routes
	book := protected.Group("/books")
    book.POST("/", h.CreateBook)
    book.GET("/:id", h.GetBook)
    book.GET("/", h.GetAllBook)
    book.PUT("/:id", h.UpdateBook)
    book.DELETE("/:id", h.DeleteBook)
	book.GET("/search", h.SearchBook)

    // Genre routes
    genre := protected.Group("/genres")
    genre.POST("/", h.CreateGenre)
    genre.GET("/:id", h.GetGenre)
    genre.GET("/", h.GetAllGenre)
    genre.PUT("/:id", h.UpdateGenre)
    genre.DELETE("/:id", h.DeleteGenre)

    // Borrower routes
    borrower := protected.Group("/borrowers")
    borrower.POST("/", h.CreateBorrower)
	borrower.GET("/:id", h.GetBorrower)
	borrower.GET("/", h.GetAllBorrower)
	borrower.PUT("/:id", h.UpdateBorrower)
	borrower.DELETE("/:id", h.DeleteBorrower)
	borrower.GET("/overdue_book", h.GetOverdueBooks)

	// Users routes
	user := protected.Group("/users")
    user.GET("/:id/borrowed_books", h.BorrowedBooks)
    user.GET("/:id/borrowing_history", h.BorrowingHistory)

	return router
}
