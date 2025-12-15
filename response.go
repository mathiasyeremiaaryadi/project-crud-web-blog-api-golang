package main

import "time"

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type PostResponseFormat struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewPostResponseFormat(post Post) PostResponseFormat {
	var resultPost PostResponseFormat
	resultPost.ID = post.ID
	resultPost.Title = post.Title
	resultPost.Content = post.Content
	resultPost.Category = post.Category
	resultPost.CreatedAt = post.CreatedAt
	resultPost.UpdatedAt = post.UpdatedAt

	for _, tag := range post.Tags {
		resultPost.Tags = append(resultPost.Tags, tag.Name)
	}

	return resultPost
}

func NewPostsResponseFormat(posts []Post) []PostResponseFormat {
	var resultPosts []PostResponseFormat
	for _, post := range posts {
		var resultPost PostResponseFormat
		resultPost.ID = post.ID
		resultPost.Title = post.Title
		resultPost.Content = post.Content
		resultPost.Category = post.Category
		resultPost.CreatedAt = post.CreatedAt
		resultPost.UpdatedAt = post.UpdatedAt

		for _, tag := range post.Tags {
			resultPost.Tags = append(resultPost.Tags, tag.Name)
		}

		resultPosts = append(resultPosts, resultPost)
	}

	return resultPosts
}

func NewSuccessResponse(data interface{}, message string) Response {
	return Response{
		Status:  "success",
		Data:    data,
		Message: message,
	}
}

func NewNotFoundResponse() Response {
	return Response{
		Status:  "error",
		Data:    nil,
		Message: "data not found",
	}
}

func NewBadRequestResponse() Response {
	return Response{
		Status:  "error",
		Data:    nil,
		Message: "request invalid",
	}
}

func NewInternalErrorResponse() Response {
	return Response{
		Status:  "error",
		Data:    nil,
		Message: "something wrong on the server",
	}
}
