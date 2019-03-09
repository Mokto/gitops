package backend

import (
	"fmt"
	"gitops/backend/github"
	"gitops/backend/helm"
	"gitops/backend/templates"
)

func initGithub() {

	//configuration := config.Get()

	fmt.Println("Cloning repo...")
	_, err := github.CloneRepo()

	if err != nil {
		fmt.Println(err)
		return
	}
	secrets, err := templates.GetSecrets()

	if err != nil {
		fmt.Println(err)
		return
	}

	templateValues := templates.ValuesTemplate{
		Branch:  "develop",
		Tag:     "release-3-8-0-r1",
		Secrets: secrets,
	}

	err = templates.WriteValues(templateValues)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done writing file")

	err = helm.InstallOrUpgradeRelease("/tmp/cloned-repo/app", "release-name", "develop")

	fmt.Println("Installing release...")

	if err != nil {
		fmt.Println(err)
		return
	}
}
