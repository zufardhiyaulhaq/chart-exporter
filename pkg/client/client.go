package client

import (
	"context"
	"flag"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"

	"github.com/zufardhiyaulhaq/chart-exporter/pkg/model"
	"github.com/zufardhiyaulhaq/chart-exporter/pkg/settings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubernetesClient struct {
	client *kubernetes.Clientset
}

func (c *KubernetesClient) Start(settings settings.Settings) {
	var config *rest.Config
	var err error

	if settings.UseServiceAccount {
		log.Println("Using serviceaccount credential")
		config, err = rest.InClusterConfig()
	} else {
		log.Println("Using kubeconfig file credential")
		var kubeConfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		}
		flag.Parse()
		config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
	}

	if err != nil {
		log.Errorln(err)
	}

	c.client, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorln(err)
	}
}

func (c *KubernetesClient) GetDeployments() []model.DeploymentInfo {
	var deploymentData []model.DeploymentInfo

	deployments, err := c.client.ExtensionsV1beta1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Errorln(err)
	}

	for _, deployment := range deployments.Items {
		value, exist := deployment.ObjectMeta.Labels["app.kubernetes.io/managed-by"]
		if exist {
			if isManagedByHelm(value) {
				var chartName string
				var chartVersion string

				chart, exist := deployment.ObjectMeta.Labels["helm.sh/chart"]
				if exist {
					chartRegex := regexp.MustCompile(`([a-z]+(-[a-z]+)*)-(\d.*)`)
					chartData := chartRegex.FindStringSubmatch(chart)

					if len(chartData) == 4 {
						chartName = chartData[1]
						chartVersion = chartData[3]

						deploymentData = append(deploymentData, model.DeploymentInfo{
							Name:         deployment.Name,
							Namespace:    deployment.Namespace,
							ChartName:    chartName,
							ChartVersion: chartVersion,
						})
					}
				}
			}
		}
	}

	return deploymentData
}

func (c *KubernetesClient) GetStatus() (bool, error) {
	api, err := c.client.ServerVersion()
	if err != nil {
		return false, err
	}

	log.Println(api.String())
	return true, nil
}

func isManagedByHelm(value string) bool {
	if value == "Helm" {
		return true
	}

	if value == "Tiller" {
		return true
	}

	return false
}
