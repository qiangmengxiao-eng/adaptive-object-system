package main

import (
	"fmt"
	"net/http"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/api"
	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/filesystem"
	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

const serverVersion = "v0.1.0"

func main() {

	fs :=
		filesystem.NewLocal(
			"./data",
		)

	repo :=
		repository.New(
			fs,
		)

	system :=
		repository.NewObjectSystem(
			repo,
		)

	server :=
		api.NewServer(
			system,
		)

	fmt.Println(
		"Adaptive Object System API Server",
	)

	fmt.Println(
		"version:",
		serverVersion,
	)

	fmt.Println(
		"listen: :8080",
	)

	err :=
		http.ListenAndServe(
			":8080",
			server.Handler(),
		)

	if err != nil {

		panic(err)
	}
}
