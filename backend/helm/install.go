package helm

import (
	"fmt"
	"io/ioutil"
	"strings"

	"k8s.io/helm/pkg/helm"
)

// InstallOrUpgradeRelease install a new Release or upgrade it if it exists
func InstallOrUpgradeRelease(chartsDir string, releaseName string, namespace string) (err error) {
	c := helm.NewClient(helm.Host("tiller-deploy:44134"))
	fmt.Println("Installing now...")

	var valuesLocation strings.Builder
	valuesLocation.WriteString(chartsDir)
	valuesLocation.WriteString("/values.yaml")

	values, err := ioutil.ReadFile(valuesLocation.String())
	if err != nil {
		return err
	}

	_, err = c.ReleaseStatus(releaseName)
	if err != nil {
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
		_, err = c.UpdateRelease(
			releaseName,
			chartsDir,
			helm.UpgradeForce(true),
			helm.UpdateValueOverrides(values),
		)
		if err != nil {
			return err
		}
		fmt.Println("Successfully updated chart.")
	}

	return
}
