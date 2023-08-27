package commands

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Noah-Wilderom/go-cli/src/utils"
	"github.com/urfave/cli/v2"
)

func RunProduction(cCtx *cli.Context) {

	pType, language := validate(cCtx)

	cwd, _ := os.Getwd()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go utils.CopyDir(utils.ResourcesFolder("production", pType, language), cwd, &wg)
	wg.Wait()

	fmt.Println("Done.")
}

func validate(cCtx *cli.Context) (string, string) {
	types := []string{
		"kubernetes",
		"docker",
		"workflows",
	}

	languages := []string{
		"laravel",
		"go",
	}

	if !cCtx.Args().Present() {
		log.Fatal("Incorrect usage, please specify the production type")
	}

	if utils.InArray(cCtx.Args().First(), types) {
		if utils.InArray(cCtx.Args().Get(1), languages) {

			return cCtx.Args().First(), cCtx.Args().Get(1)

		} else {
			log.Fatalf("Incorrect usage, language must be one of these languages: %s", strings.Join(languages, ", "))
		}
	} else {
		log.Fatalf("Incorrect usage, production type must be one of these types: %s", strings.Join(types, ", "))
	}

	log.Fatal("Unexpected error occurred...")

	return "", ""
}
