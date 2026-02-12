package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/parmeet20/bloghive/internal/model"
	"github.com/parmeet20/bloghive/internal/service"
)

type BlogHandler struct {
	service service.BlogService
}

func NewBlogHandler(r *gin.Engine, service service.BlogService) {
	h := &BlogHandler{service}

	blogs := r.Group("/api/v1/blogs")
	{
		blogs.POST("", h.Create)
		blogs.GET("", h.GetAll)
		blogs.GET("/:id", h.GetByID)
		blogs.PUT("/:id", h.Update)
	}
}

func (h *BlogHandler) Create(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, blog)
}

func (h *BlogHandler) GetAll(c *gin.Context) {
	blogs, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

func (h *BlogHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	blog, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(id, &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}
