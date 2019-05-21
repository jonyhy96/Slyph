package deployment

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"Slyph/models"
	"net/http"
)

// Updatedeployment updates deployment
func Updatedeployment(req *models.Deployment) (int, error) {
	log := models.Log{}
	deploymentsClient := models.GetClientset().AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	// Update Deployment
	result, getErr := deploymentsClient.Get(req.Name, metav1.GetOptions{}) //Check if deployment already exist
	if getErr != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to Get deployment " + req.Name + " " + getErr.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, getErr
	}
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result.Spec.Replicas = int32Ptr(req.Replicas)
		for index := range req.Containers {
			result.Spec.Template.Spec.Containers[index].Image = req.Containers[index].Image // change deployment version
			result.Spec.Template.Spec.Containers[index].Env = req.Containers[index].Env
			result.Spec.Template.Spec.Containers[index].Ports[index].ContainerPort = req.Containers[index].ContainerPort
		}
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to Update deployment " + req.Name + " " + retryErr.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, retryErr
	}
	err := models.Updateimage(req.Name, req.Containers[0].Image)
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to Update deployment " + req.Name + " " + err.Error()
		go models.Insertlog(&log)
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Update deployment " + req.Name + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
