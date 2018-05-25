package service

import (
	// "k8s.io/api/apps/v1beta2"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s/models"
	"net/http"
)

// DeleteService delete a service
func DeleteService(req *models.Delete) (int, error) { // just require req.Name(service name)
	log := models.Log{}
	serviceclient := models.GetClientset().CoreV1().Services(apiv1.NamespaceDefault)
	serviceimp, err := serviceclient.Get(req.Name, metav1.GetOptions{})
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to get latest version of service " + err.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	// delete service
	if serviceimp.Name != " " {
		if err := models.GetClientset().CoreV1().Services(req.Namespace).Delete(req.Name, &metav1.DeleteOptions{}); err != nil {
			log.Time = models.Gettime()
			log.Etype = "error"
			log.Event = "Failed to delete service " + req.Name + " " + err.Error()
			go models.Insertlog(&log)
			return http.StatusInternalServerError, err
		}
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Delete service " + req.Name + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
