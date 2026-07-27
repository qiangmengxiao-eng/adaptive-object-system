package main

import (
	"fmt"
	"os"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/filesystem"
	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

func main() {
	system := createSystem()

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("usage: aos <command>")
		return
	}

	switch args[0] {

	case "object":
		handleObject(system, args[1:])

	default:
		fmt.Println("unknown command:", args[0])
	}
}

func createSystem() *repository.ObjectSystem {
	fs := filesystem.NewLocal("./data")

	repo := repository.New(fs)

	return repository.NewObjectSystem(repo)
}

func handleObject(system *repository.ObjectSystem, args []string) {

	if len(args) == 0 {
		fmt.Println("usage: aos object <command>")
		return
	}

	switch args[0] {

	case "list":
		objects, err := system.Registry.List()
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println(objects)

	case "get":
		if len(args) < 2 {
			fmt.Println("object name required")
			return
		}

		object, err := system.Registry.Get(args[1])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Printf("name: %s\n", object.Definition.Name)
		fmt.Printf("type: %s\n", object.Definition.Type)
		fmt.Printf("version: %d\n", object.Metadata.Version)

	case "create":
		if len(args) < 2 {
			fmt.Println("object name required")
			return
		}

		name := args[1]

		err := system.Registry.Register(
			name,
			[]byte("name: "+name),
			&repository.ObjectMetadata{
				Version: 1,
			},
		)

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("created object:", name)

	default:
		fmt.Println("unknown object command:", args[0])
	}
}
