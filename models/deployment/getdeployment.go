package deployment

import (
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s/models"
)

// Getdeploymentinfo gets deployments information
func Getdeploymentinfo() *appsv1beta1.DeploymentList {
	log := models.Log{}
	deployments, err := models.GetClientset().AppsV1beta1().Deployments("default").List(metav1.ListOptions{})
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to get deployment " + err.Error()
		go models.Insertlog(&log)
		return nil
	}
	return deployments
}
