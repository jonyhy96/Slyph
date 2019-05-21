package pod

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"Slyph/models"
	"net/http"
)

func PodDelete(req *models.Delete) (int, error) {
	log := new(models.Log)
	if err := models.GetClientset().CoreV1().Pods(req.Namespace).Delete(req.Name, &metav1.DeleteOptions{}); err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to delete pod " + req.Name + " " + err.Error()
		models.Insertlog(log)
		return http.StatusInternalServerError, err
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Delete pod " + req.Name + " success"
	models.Insertlog(log)
	return http.StatusOK, nil
}
