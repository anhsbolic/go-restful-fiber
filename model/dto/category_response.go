package dto

import "go-restful-fiber/model/domain"

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponse(category domain.Category) CategoryResponse {
	return CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []CategoryResponse {
	var categoryResponses []CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
