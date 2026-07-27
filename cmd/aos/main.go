package main

import (
	"fmt"
	"os"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/filesystem"
	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

const version = "v0.1.0"

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("usage: aos <command>")
		return
	}

	fs := filesystem.NewLocal("./data")

	repo := repository.New(fs)

	system := repository.NewObjectSystem(repo)

	switch args[0] {

	case "version":

		fmt.Println(
			"adaptive-object-system",
			version,
		)

	case "object":

		handleObject(
			system,
			args[1:],
		)

	case "behavior":

		handleBehavior(
			system,
			args[1:],
		)

	case "graph":

		handleGraph(
			system,
			args[1:],
		)

	default:

		fmt.Println(
			"unknown command:",
			args[0],
		)
	}
}

func handleObject(
	system *repository.ObjectSystem,
	args []string,
) {

	if len(args) == 0 {
		fmt.Println(
			"usage: aos object <command>",
		)
		return
	}

	switch args[0] {

	case "list":

		objects, err :=
			system.Registry.List()

		if err != nil {
			fmt.Println(
				"error:",
				err,
			)
			return
		}

		fmt.Println(objects)

	case "create":

		if len(args) < 2 {

			fmt.Println(
				"object name required",
			)

			return
		}

		name := args[1]

		definition := []byte(
			"name: " + name + "\n" +
				"type: object\n" +
				"version: 1\n",
		)

		err :=
			system.Registry.Register(
				name,
				definition,
				&repository.ObjectMetadata{
					Version: 1,
				},
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(
			"created object:",
			name,
		)

	case "get":

		if len(args) < 2 {

			fmt.Println(
				"object name required",
			)

			return
		}

		object, err :=
			system.Registry.Get(
				args[1],
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(object)

	case "inspect":

		if len(args) < 2 {

			fmt.Println(
				"object name required",
			)

			return
		}

		definition, err :=
			system.Repository.ReadObjectDefinition(
				args[1],
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(
			"Object:",
			definition.Name,
		)

		fmt.Println(
			"Type:",
			definition.Type,
		)

		fmt.Println(
			"Version:",
			definition.Version,
		)

	case "delete":

		if len(args) < 2 {

			fmt.Println(
				"object name required",
			)

			return
		}

		err :=
			system.Repository.DeleteObject(
				args[1],
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(
			"deleted object:",
			args[1],
		)

	case "migrate":

		if len(args) < 2 {

			fmt.Println(
				"usage: aos object migrate <name>",
			)

			return
		}

		name := args[1]

		definition, err :=
			system.Repository.ReadObjectDefinition(
				name,
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		oldVersion :=
			definition.Version

		err =
			system.Repository.MigrateObject(
				name,
				system.MigrationService,
				oldVersion+1,
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Printf(
			"migrated object: %s %d -> %d\n",
			name,
			oldVersion,
			oldVersion+1,
		)

	default:

		fmt.Println(
			"unknown object command:",
			args[0],
		)
	}
}

func handleBehavior(
	system *repository.ObjectSystem,
	args []string,
) {

	if len(args) == 0 {

		fmt.Println(
			"usage: aos behavior <command>",
		)

		return
	}

	switch args[0] {

	case "list":

		fmt.Println(
			system.BehaviorService.List(),
		)

	case "run":

		if len(args) < 2 {

			fmt.Println(
				"behavior name required",
			)

			return
		}

		err :=
			system.BehaviorService.Run(
				args[1],
				nil,
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(
			"executed behavior:",
			args[1],
		)
	}
}

func handleGraph(
	system *repository.ObjectSystem,
	args []string,
) {

	if len(args) == 0 {

		fmt.Println(
			"usage: aos graph <command>",
		)

		return
	}

	switch args[0] {

	case "list":

		relations, err :=
			system.GraphService.List()

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(
			relations,
		)

	case "add":

		if len(args) < 4 {

			fmt.Println(
				"usage: aos graph add <from> <to> <type>",
			)

			return
		}

		err :=
			system.GraphService.AddRelation(
				repository.ObjectRelation{
					From: args[1],
					To:   args[2],
					Type: args[3],
				},
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(
			"relation added",
		)

	case "query":

		if len(args) < 2 {

			fmt.Println(
				"usage: aos graph query <object>",
			)

			return
		}

		relations, err :=
			system.GraphService.QueryRelations(
				args[1],
			)

		if err != nil {

			fmt.Println(
				"error:",
				err,
			)

			return
		}

		fmt.Println(
			relations,
		)

	default:

		fmt.Println(
			"unknown graph command:",
			args[0],
		)
	}
}
