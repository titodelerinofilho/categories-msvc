package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/titodelerinofilho/categories-msvc/internal/repositories"
	use_cases "github.com/titodelerinofilho/categories-msvc/internal/use-cases"
)

func ListCategories(context *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewListCategoryUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": true, "categories": categories})
}
