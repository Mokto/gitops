package helm

import (
	"fmt"
	"gitops/backend/utils"
	"io/ioutil"
	"k8s.io/helm/pkg/helm"
)

// InstallOrUpgradeRelease install a new Release or upgrade it if it exists
func InstallOrUpgradeRelease(projectPath string, releaseName string, namespace string) (err error) {
	c := helm.NewClient(helm.Host("tiller-deploy:44134"))
	fmt.Println("Installing now...")

	valuesLocation := utils.ComposeStrings(projectPath, "/chart/values.yaml")
	chartsDir := utils.ComposeStrings(projectPath, "/chart")

	fmt.Println(valuesLocation)
	values, err := ioutil.ReadFile(valuesLocation)
	if err != nil {
		return err
	}
	fmt.Println("1")

	_, err = c.ReleaseStatus(releaseName)
	if err != nil {
		fmt.Println("install")
		_, err = c.InstallRelease(
			chartsDir,
			namespace,
			helm.ReleaseName(releaseName),
			helm.ValueOverrides(values),
		)
		if err != nil {
			return err
		}
		fmt.Println("New release created.")
	} else {
		fmt.Println("upgrade")
		_, err = c.UpdateRelease(
			releaseName,
			chartsDir,
			helm.UpgradeForce(true),
			helm.UpdateValueOverrides(values),
		)
		fmt.Println("upgrade done")
		if err != nil {
			fmt.Println("err")
			return err
		}
		fmt.Println("Successfully updated chart.")
	}

	return
}
