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
	case "behavior":

		if len(args) < 2 {
			fmt.Println("usage: aos behavior <command>")
			return
		}

		switch args[1] {

		case "list":

			fmt.Println(system.BehaviorService.List())

		case "run":

			if len(args) < 3 {
				fmt.Println("behavior name required")
				return
			}

			err := system.BehaviorService.Run(
				args[2],
				nil,
			)

			if err != nil {
				fmt.Println("error:", err)
				return
			}

			fmt.Println("executed behavior:", args[2])

		default:

			fmt.Println("unknown behavior command")
		}
	case "version":
		fmt.Println("adaptive-object-system v0.1.0")

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
	case "search":
		if len(args) < 2 {
			fmt.Println("query required")
			return
		}

		objects, err := system.Repository.QueryObjects(
			repository.ObjectQuery{
				Name: args[1],
			},
		)

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

	case "inspect":
		if len(args) < 2 {
			fmt.Println("object name required")
			return
		}

		object, err := system.Registry.Get(args[1])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Printf("Object: %s\n", object.Definition.Name)
		fmt.Printf("Type: %s\n", object.Definition.Type)
		fmt.Printf("Version: %d\n", object.Metadata.Version)

	case "delete":
		if len(args) < 2 {
			fmt.Println("object name required")
			return
		}

		err := system.Repository.DeleteObject(args[1])

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("deleted object:", args[1])
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
