// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package example

// Injectors from wire.go:

func GetPostService() *PostService {
	iPostRepo := NewPostRepo()
	iPostUsecase := NewPostUsecase(iPostRepo)
	postService := NewPostService(iPostUsecase)
	return postService
}
