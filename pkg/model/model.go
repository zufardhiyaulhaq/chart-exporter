package model

type DeploymentInfo struct {
	APIVersion   string
	Name         string
	Namespace    string
	ChartName    string
	ChartVersion string
}
