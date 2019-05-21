package service

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"Slyph/models"
)

// GetServiceinfo get the services info
func GetServiceinfo() *v1.ServiceList {
	log := models.Log{}
	services, err := models.GetClientset().CoreV1().Services("default").List(metav1.ListOptions{})
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to get service " + err.Error()
		go models.Insertlog(&log)
	}
	return services
}
