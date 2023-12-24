package generate

import (
	"gopkg.in/yaml.v3"
	"os"
)

type HelmChart struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   Metadata
	Spec       Spec
}

type Metadata struct {
	Name string
}

type Spec struct {
	Interval string
	Chart    ChartSpec
	Values   Values
}

type ChartSpec struct {
	Spec ChartSpecValues
}

type ChartSpecValues struct {
	Chart     string
	Version   string
	Interval  string
	SourceRef SourceRef `yaml:"sourceRef"`
}

type SourceRef struct {
	Kind      string
	Name      string
	Namespace string
}

type Values struct {
	Controllers Controllers
	Service     string `yaml:"service,omitempty"`
	Ingress     string `yaml:"ingress,omitempty"`
	Persistence string `yaml:"persistence,omitempty"`
}

type Controllers struct {
	Main Controller
}

type Controller struct {
	Annotations string `yaml:"annotations,omitempty"`
	Pod         string `yaml:"pod,omitempty"`
	Containers  Containers
}

type Containers struct {
	Main Container
}

type Container struct {
	Image ContainerImage
}

type ContainerImage struct {
	Repository string
	Tag        string
}

func NewChart() HelmChart {

	chart := HelmChart{
		ApiVersion: "helm.toolkit.fluxcd.io/v2beta2",
		Kind:       "HelmRelease",
		Metadata:   Metadata{},
		Spec: Spec{
			Interval: "30m",
			Chart: ChartSpec{
				ChartSpecValues{
					Chart:    "app-template",
					Version:  "2.4.0",
					Interval: "30m",
					SourceRef: SourceRef{
						Kind:      "HelmRepository",
						Name:      "bjw-s",
						Namespace: "flux-system",
					},
				},
			},
			Values: Values{},
		},
	}

	return chart
}

func Build(chart HelmChart, namespace string) {
	appsDir := os.Getenv("BOSUN_FLUX_DIR") + "/kubernetes/apps"
	nsDir := appsDir + "/" + namespace
	if _, err := os.Stat(nsDir); os.IsNotExist(err) {
		_ = os.Mkdir(nsDir, 0755)
	}
	releaseDir := nsDir + "/" + chart.Metadata.Name
	_ = os.Mkdir(releaseDir, 0755)

	releaseAppDir := releaseDir + "/" + "app"
	_ = os.Mkdir(releaseAppDir, 0755)
	chartYaml, _ := yaml.Marshal(chart)
	os.WriteFile(releaseAppDir+"/"+"helmrelease.yaml", chartYaml, 0755)
}
