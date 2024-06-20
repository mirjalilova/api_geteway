package handler

import (
	"api-geteway/genproto/library_service"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// CreateAuthor
// @Summary Create author item
// @Description Create a new author item
// @Tags Author
// @Accept json
// @Produce json
// @Param author body library_service.AuthorCreateReq true "Author data"
// @Success 200 {object} library_service.AuthorRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /authors [post]
func (h *HTTPHandler) CreateAuthor(ctx *gin.Context) {
	var req library_service.AuthorCreateReq
	if err := ctx.BindJSON(&req); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Author.Create(context.Background(), &req)
	if err != nil {
		slog.Error("failed to create author: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetAuthor handles getting an author item by ID.
// @Summary Get author item
// @Description Get an author item by ID
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} library_service.AuthorRes
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /authors/{id} [get]
func (h *HTTPHandler) GetAuthor(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	res, err := h.Author.Get(context.Background(), &req)
	if err != nil {
		slog.Error("failed to get author: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetAllAuthor handles getting all menu items.
// @Summary Get all author items
// @Description Get all author items
// @Tags Author
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} library_service.AuthorGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /authors [get]
func (h *HTTPHandler) GetAllAuthor(ctx *gin.Context) {

	name := ctx.Query("name")
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

	res, err := h.Author.GetAll(context.Background(), &library_service.AuthorGetAllReq{
		Name: name,
		Filter: &library_service.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		slog.Error("failed to get author: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UpdateAuthor handles updating an existing author item.
// @Summary Update author item
// @Description Update an existing author item
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Param menu body library_service.AuthorCreateReq true "Updated author item data"
// @Success 200 {object} library_service.AuthorRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /authors/{id} [put]
func (h *HTTPHandler) UpdateAuthor(ctx *gin.Context) {
	var req library_service.AuthorUpdateReq

	id := ctx.Param("id")

	req.Id = &library_service.GetByIdReq{
		Id: id,
	}

	body := library_service.AuthorCreateReq{}

	if err := ctx.BindJSON(&body); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.UpdateAuthor = &body

	res, err := h.Author.Update(context.Background(), &req)
	if err != nil {
		slog.Error("failed to update author: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// DeleteAuthor handles deleting an author item by ID.
// @Summary Delete menu author
// @Description Delete an author item by ID
// @Tags Author
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} string "Author item deleted"
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /authors/{id} [delete]
func (h *HTTPHandler) DeleteAuthor(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	_, err := h.Author.Delete(context.Background(), &req)
	if err != nil {
		slog.Error("failed to delete author: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Successfully deleted"})
}
