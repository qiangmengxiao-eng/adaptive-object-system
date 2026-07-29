package main

import (
	"fmt"
	"os"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/filesystem"
	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

const version = "v0.1.0"

func main() {

	args :=
		os.Args[1:]

	if len(args) == 0 {

		fmt.Println(
			"usage: aos <command>",
		)

		return
	}

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

	switch args[0] {

	case "version":

		fmt.Println(
			"adaptive-object-system",
			version,
		)

	case "system":

		handleSystem(
			system,
			args[1:],
		)

	case "object":

		handleObject(
			system,
			args[1:],
		)

	case "graph":

		handleGraph(
			system,
			args[1:],
		)

	case "event":

		handleEvent(
			system,
			args[1:],
		)

	case "behavior":

		handleBehavior(
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

// system commands

func handleSystem(
	system *repository.ObjectSystem,
	args []string,
) {

	if len(args) == 0 {

		fmt.Println(
			"usage: aos system status",
		)

		return
	}

	switch args[0] {

	case "status":

		status :=
			system.StatusService.Get()

		fmt.Println(
			"Adaptive Object System",
		)

		fmt.Println()

		fmt.Println(
			"Objects:",
			status.Objects,
		)

		fmt.Println(
			"Runtime:",
			status.Runtime,
		)

		fmt.Println(
			"Events:",
			status.Events,
		)

		fmt.Println(
			"Behaviors:",
			status.Behaviors,
		)

		fmt.Println(
			"Audit:",
			status.Audit,
		)
	}
}

// object commands

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

	case "query":

		objects, err :=
			system.QueryService.ListObjects()

		if err != nil {

			fmt.Println(
				err,
			)

			return
		}

		fmt.Println(
			objects,
		)

	case "list":

		objects, err :=
			system.Registry.List()

		if err != nil {

			fmt.Println(
				err,
			)

			return
		}

		fmt.Println(
			objects,
		)

	case "create":

		if len(args) < 2 {

			fmt.Println(
				"object name required",
			)

			return
		}

		name :=
			args[1]

		definition :=
			[]byte(
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
				err,
			)

			return
		}

		objectDefinition, err :=
			system.Repository.ReadObjectDefinition(
				name,
			)

		if err == nil {

			system.Runtime.Start(
				*objectDefinition,
			)

			event :=
				repository.NewObjectEvent(
					"object.created",
					name,
					"create",
					"",
				)

			_ =
				system.EventBus.Publish(
					event,
				)

			_ =
				system.Runtime.AddEvent(
					name,
					event,
				)
		}

		fmt.Println(
			"created object:",
			name,
		)

	case "runtime":

		if len(args) < 2 {

			fmt.Println(
				"object name required",
			)

			return
		}

		object, ok :=
			system.Runtime.Get(
				args[1],
			)

		if !ok {

			fmt.Println(
				"runtime object not found",
			)

			return
		}

		fmt.Println(
			"Object:",
			object.Name,
		)

		fmt.Println(
			"State:",
			object.State.Status,
		)

		fmt.Println(
			"Events:",
			len(object.Events),
		)

	case "inspect":

		if len(args) < 2 {

			fmt.Println(
				"object name required",
			)

			return
		}

		view, err :=
			system.ObjectViewService.Inspect(
				args[1],
			)

		if err != nil {

			fmt.Println(
				err,
			)

			return
		}

		fmt.Println(
			"Object:",
			view.Name,
		)

		fmt.Println(
			"Type:",
			view.Definition.Type,
		)

		fmt.Println(
			"Version:",
			view.Definition.Version,
		)

		if view.Runtime != nil {

			fmt.Println(
				"State:",
				view.Runtime.State.Status,
			)

			fmt.Println(
				"Events:",
				len(view.Events),
			)
		}

		fmt.Println(
			"Relations:",
			view.Relations,
		)

		fmt.Println(
			"Behaviors:",
			view.Behaviors,
		)

		fmt.Println(
			"Audit:",
			view.Audit,
		)
	}
}

// graph commands

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

	case "add":

		if len(args) < 4 {

			fmt.Println(
				"usage: aos graph add <from> <to> <relation>",
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
				err,
			)

			return
		}

		fmt.Println(
			relations,
		)

	case "inspect":

		if len(args) < 2 {

			fmt.Println(
				"object required",
			)

			return
		}

		view, err :=
			system.ObjectViewService.Inspect(
				args[1],
			)

		if err != nil {

			fmt.Println(
				err,
			)

			return
		}

		fmt.Println(
			"Object:",
			view.Name,
		)

		fmt.Println(
			"Relations:",
			view.Relations,
		)
	}
}

// event commands

func handleEvent(
	system *repository.ObjectSystem,
	args []string,
) {

	if len(args) == 0 {

		fmt.Println(
			"usage: aos event list",
		)

		return
	}

	switch args[0] {

	case "list":

		events, err :=
			system.EventService.List()

		if err != nil {

			fmt.Println(
				err,
			)

			return
		}

		fmt.Println(
			events,
		)
	}
}

// behavior commands

func handleBehavior(
	system *repository.ObjectSystem,
	args []string,
) {

	fmt.Println(
		"behavior command",
		args,
	)
}
