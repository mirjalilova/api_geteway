package handler

import (
	"api-geteway/genproto/library_service"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// CreateGenre
// @Summary Create genre item
// @Description Create a new genre item
// @Tags Genre
// @Accept json
// @Produce json
// @Param genre body library_service.GenreCreateReq true "Genre data"
// @Success 200 {object} library_service.GenreRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /genres [post]
func (h *HTTPHandler) CreateGenre(ctx *gin.Context) {
	var req library_service.GenreCreateReq

	if err := ctx.BindJSON(&req); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Genre.Create(context.Background(), &req)
	if err != nil {
		slog.Error("failed to create genre: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetGenre handles getting a genre item by ID.
// @Summary Get genre item
// @Description Get a genre item by ID
// @Tags Genre
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Success 200 {object} library_service.GenreRes
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /genres/{id} [get]
func (h *HTTPHandler) GetGenre(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	res, err := h.Genre.Get(context.Background(), &req)
	if err != nil {
		slog.Error("failed to get genre: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// GetAllGenre handles getting all menu items.
// @Summary Get all genre items
// @Description Get all genre items
// @Tags Genre
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} library_service.GenreGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /genres [get]
func (h *HTTPHandler) GetAllGenre(ctx *gin.Context) {

	name := ctx.Query("name")
	limitStr := ctx.Query("limit")
	offsetStr := ctx.Query("offset")

	var limit int
	var offset int
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

	res, err := h.Genre.GetAll(context.Background(), &library_service.GenreGetAllReq{
		Name: name,
		Filter: &library_service.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		slog.Error("failed to get genre: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

// UpdateGenre handles updating an existing genre item.
// @Summary Update genre item
// @Description Update an existing genre item
// @Tags Genre
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Param menu body library_service.GenreCreateReq true "Updated genre item data"
// @Success 200 {object} library_service.GenreRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /genres/{id} [put]
func (h *HTTPHandler) UpdateGenre(ctx *gin.Context) {
	var req library_service.GenreUpdateReq

	id := ctx.Param("id")

	body := &library_service.GenreCreateReq{}

	if err := ctx.BindJSON(&body); err != nil {
		slog.Error("failed to bind json: %v", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.Id = &library_service.GetByIdReq{
		Id: id,
	}

	req.UpdateGenre = body

	res, err := h.Genre.Update(context.Background(), &req)
	if err != nil {
		slog.Error("failed to update genre: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("ressssssssssssssssss", res)

	ctx.JSON(200, res)
}

// DeleteGenre handles deleting an genre item by ID.
// @Summary Delete menu genre
// @Description Delete an genre item by ID
// @Tags Genre
// @Accept json
// @Produce json
// @Param id path string true "Genre ID"
// @Success 200 {object} string "Genre item deleted"
// @Failure 400 {object} string "Invalid menu item ID"
// @Failure 404 {object} string "Menu item not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /genres/{id} [delete]
func (h *HTTPHandler) DeleteGenre(ctx *gin.Context) {
	var req library_service.GetByIdReq

	id := ctx.Param("id")

	req.Id = id

	_, err := h.Genre.Delete(context.Background(), &req)
	if err != nil {
		slog.Error("failed to delete genre: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Genre deleted"})
}
