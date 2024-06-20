package handler

import (
	"api-geteway/genproto/library_service"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// CreateBorrower
// @Summary Create borrower item
// @Description Create a new borrower item
// @Tags Borrower
// @Accept json
// @Produce json
// @Param borrower body library_service.BorrowerCreateReq true "Borrower data"
// @Success 200 {object} library_service.BorrowerRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /borrowers [post]
func (h *HTTPHandler) CreateBorrower(ctx *gin.Context) {
	var req library_service.BorrowerCreateReq

	if err := ctx.BindJSON(&req); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Borrower.Create(context.Background(), &req)
	if err != nil {
		slog.Error("failed to create borrower: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetBorrower handles getting a borrower item by ID.
// @Summary Get borrower item
// @Description Get a borrower item by ID
// @Tags Borrower
// @Accept json
// @Produce json
// @Param id path string true "Borrower ID"
// @Success 200 {object} library_service.BorrowerRes
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /borrowers/{id} [get]
func (h *HTTPHandler) GetBorrower(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	res, err := h.Borrower.Get(context.Background(), &req)
	if err != nil {
		slog.Error("failed to get borrower: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetAllBorrower handles getting all menu items.
// @Summary Get all borrower items
// @Description Get all borrower items
// @Tags Borrower
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Param book_id query string false "book_id"
// @Param borrow_date query string false "borrow_date"
// @Param return_date query string false "return_date"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} library_service.BorrowerGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /borrowers [get]
func (h *HTTPHandler) GetAllBorrower(ctx *gin.Context) {

	uid := ctx.Query("user_id")
	bid := ctx.Query("book_id")
	brid := ctx.Query("borrow_date")
	rid := ctx.Query("return_date")
	limitStr := ctx.Query("limit")
	offsetStr := ctx.Query("offset")

	var limit, offset int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			slog.Error("failed to get limit: %v", err)
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	} else {
		limit = 0
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			slog.Error("failed to get offset: %v", err)
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	} else {
		offset = 0
	}

	res, err := h.Borrower.GetAll(context.Background(), &library_service.BorrowerGetAllReq{
		UserId:     uid,
		BookId:     bid,
		BorrowDate: brid,
		ReturnDate: rid,
		Filter: &library_service.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})

	if err != nil {
		slog.Error("failed to get borrower: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UpdateBorrower handles updating an existing borrower item.
// @Summary Update borrower item
// @Description Update an existing borrower item
// @Tags Borrower
// @Accept json
// @Produce json
// @Param id path string true "Borrower ID"
// @Param menu body library_service.BorrowerUpdate true "Updated borrower item data"
// @Success 200 {object} library_service.BorrowerRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /borrowers/{id} [put]
func (h *HTTPHandler) UpdateBorrower(ctx *gin.Context) {
	var req library_service.BorrowerUpdateReq
	id := ctx.Param("id")

	req.Id = &library_service.GetByIdReq{
		Id: id,
	}

	body := library_service.BorrowerUpdate{}

	if err := ctx.BindJSON(&body); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.UpdateBorrower = &body

	res, err := h.Borrower.Update(context.Background(), &req)
	if err != nil {
		slog.Error("failed to update borrower: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteBorrower handles deleting an borrower item by ID.
// @Summary Delete menu borrower
// @Description Delete an borrower item by ID
// @Tags Borrower
// @Accept json
// @Produce json
// @Param id path string true "Borrower ID"
// @Success 200 {object} string "Borrower item deleted"
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /borrowers/{id} [delete]
func (h *HTTPHandler) DeleteBorrower(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	_, err := h.Borrower.Delete(context.Background(), &req)
	if err != nil {
		slog.Error("failed to delete borrower: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Borrower deleted"})
}

// GetOverdueBooks
// @Summary Get overdue book items
// @Description Get overdue book items
// @Tags Borrower
// @Accept json
// @Produce json
// @Success 200 {object} library_service.BorrowerGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /borrowers/overdue_book [get]
func (h *HTTPHandler) GetOverdueBooks(ctx *gin.Context) {

	res, err := h.Borrower.GetOverdueBooks(context.Background(), &library_service.Void{})
	if err != nil {
		slog.Error("Error while getting overdue book : %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// BorrowedBooks handles getting all borrowed books.
// @Summary Get all borrowed books
// @Description Get all borrowed books
// @Tags User
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Success 200 {object} library_service.BorrowedBooksRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /users/{user_id}/borrowed_books [get]
func (h *HTTPHandler) BorrowedBooks(ctx *gin.Context) {
	var req library_service.BorrowedBooksReq

	id := ctx.Query("user_id")

	req.UserId = id

	res, err := h.Borrower.GetBorrowedBooks(context.Background(), &req)
	if err != nil {
		slog.Error("failed to get borrowed books: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// BorrowingHistory handles getting all borrowed books history.
// @Summary Get all borrowed books history
// @Description Get all borrowed books history
// @Tags User
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Success 200 {object} library_service.BorrowedBooksRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /users/{user_id}/borrowing_history [get]
func (h *HTTPHandler) BorrowingHistory(ctx *gin.Context) {
	var req library_service.BorrowedBooksReq

	id := ctx.Query("user_id")

	req.UserId = id

	res, err := h.Borrower.GetBorrowingHistory(context.Background(), &req)
	if err != nil {
		slog.Error("failed to get borrowing history: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
