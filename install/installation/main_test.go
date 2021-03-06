package installation

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kyma-incubator/hydroform/install/scheme"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	tillerYamlContent      string
	installerYamlContent   string
	installerCRYamlContent string

	upgradeTillerYamlContent      string
	upgradeInstallerYamlContent   string
	upgradeInstallerCRYamlContent string

	resourcesSchema *runtime.Scheme

	decoder runtime.Decoder
)

func TestMain(m *testing.M) {
	tillerYamlBytes, err := ioutil.ReadFile("testdata/tiller.yaml")
	logAndExitOnError(err)
	tillerYamlContent = string(tillerYamlBytes)

	installerYamlBytes, err := ioutil.ReadFile("testdata/kyma-installer.yaml")
	logAndExitOnError(err)
	installerYamlContent = string(installerYamlBytes)

	installerCRYamlBytes, err := ioutil.ReadFile("testdata/kyma-installer-cr.yaml")
	logAndExitOnError(err)
	installerCRYamlContent = string(installerCRYamlBytes)

	tillerUpgradeYamlBytes, err := ioutil.ReadFile("testdata/tiller-upgrade.yaml")
	logAndExitOnError(err)
	upgradeTillerYamlContent = string(tillerUpgradeYamlBytes)

	installerUpgradeYamlBytes, err := ioutil.ReadFile("testdata/kyma-installer-upgrade.yaml")
	logAndExitOnError(err)
	upgradeInstallerYamlContent = string(installerUpgradeYamlBytes)

	installerUpgradeCRYamlBytes, err := ioutil.ReadFile("testdata/kyma-installer-upgrade-cr.yaml")
	logAndExitOnError(err)
	upgradeInstallerCRYamlContent = string(installerUpgradeCRYamlBytes)

	resourcesSchema, err = scheme.DefaultScheme()
	logAndExitOnError(err)

	decoder, err = scheme.DefaultDecoder()
	logAndExitOnError(err)

	code, err := runTests(m)
	logAndExitOnError(err)

	os.Exit(code)
}

func runTests(m *testing.M) (int, error) {
	codecs := serializer.NewCodecFactory(resourcesSchema)
	decoder = codecs.UniversalDeserializer()

	code := m.Run()

	return code, nil
}

func logAndExitOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
