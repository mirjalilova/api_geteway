package handler

import (
	"api-geteway/genproto/library_service"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// CreateBook
// @Summary Create book item
// @Description Create a new book item
// @Tags Book
// @Accept json
// @Produce json
// @Param book body library_service.BookCreateReq true "Book data"
// @Success 200 {object} library_service.BookRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /books [post]
func (h *HTTPHandler) CreateBook(ctx *gin.Context) {
	var req library_service.BookCreateReq

	if err := ctx.BindJSON(&req); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Book.Create(context.Background(), &req)
	if err != nil {
		slog.Error("failed to create book: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetBook handles getting a book item by ID.
// @Summary Get book item
// @Description Get a book item by ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} library_service.BookRes
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /books/{id} [get]
func (h *HTTPHandler) GetBook(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	res, err := h.Book.Get(context.Background(), &req)
	if err != nil {
		slog.Error("failed to get book: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetAllBook handles getting all menu items.
// @Summary Get all book items
// @Description Get all book items
// @Tags Book
// @Accept json
// @Produce json
// @Param author_id query string false "author_id"
// @Param genre_id query string false "genre_id"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} library_service.BookGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /books [get]
func (h *HTTPHandler) GetAllBook(ctx *gin.Context) {

	aid := ctx.Query("author_id")
	gid := ctx.Query("genre_id")
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

	res, err := h.Book.GetAll(context.Background(), &library_service.BookGetAllReq{
		AuthorId: aid,
		GenreId:  gid,
		Filter: &library_service.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		slog.Error("failed to get book: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UpdateBook handles updating an existing book item.
// @Summary Update book item
// @Description Update an existing book item
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param menu body library_service.BookCreateReq true "Updated book item data"
// @Success 200 {object} library_service.BookRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /books/{id} [put]
func (h *HTTPHandler) UpdateBook(ctx *gin.Context) {
	var req library_service.BookUpdateReq

	id := ctx.Param("id")

	req.Id = &library_service.GetByIdReq{
		Id: id,
	}

	body := library_service.BookCreateReq{}

	if err := ctx.BindJSON(&body); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.UpdateBook = &body

	res, err := h.Book.Update(context.Background(), &req)
	if err != nil {
		slog.Error("failed to update book: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteBook handles deleting an book item by ID.
// @Summary Delete menu book
// @Description Delete an book item by ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} string "Book item deleted"
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /books/{id} [delete]
func (h *HTTPHandler) DeleteBook(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	_, err := h.Book.Delete(context.Background(), &req)
	if err != nil {
		slog.Error("failed to delete book: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Book deleted"})
}

// SearchBook handles searching book items.
// @Summary Search book items
// @Description Search book items
// @Tags Book
// @Accept json
// @Produce json
// @Param title query string false "title"
// @Param author query string false "author"
// @Success 200 {object} library_service.BookGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /books/search [get]
func (h *HTTPHandler) SearchBook(ctx *gin.Context) {
	title := ctx.Query("title")
	author := ctx.Query("author")

	res, err := h.Book.Search(context.Background(), &library_service.BookSearchReq{
		Title:  title,
		Author: author,
	})
	if err != nil {
		slog.Error("failed to search book: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}
