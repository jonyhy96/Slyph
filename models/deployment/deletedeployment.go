package deployment

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"Slyph/models"
	"net/http"
)

//Deletedeployment delete a deployment
func Deletedeployment(req *models.Delete) (int, error) {
	log := models.Log{}
	deletePolicy := metav1.DeletePropagationForeground
	if err := models.GetClientset().AppsV1beta1().Deployments(req.Namespace).Delete(req.Name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to delete  deployment " + req.Name + " " + err.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Delete deployment " + req.Name + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
