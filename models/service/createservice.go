package service

import (
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
	"strings"
	// "k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"Slyph/models"
	"net/http"
)

// AddService add a service
func AddService(req *models.Service) (int, error) {
	log := models.Log{}
	service := v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			Labels: map[string]string{
				"name": req.Name,
			},
		},
	}
	service.Spec = v1.ServiceSpec{
		Type: req.Type,
		Selector: map[string]string{
			"name": req.Selector["name"],
		},
	}
	for _, item := range req.Ports {
		if len(item.Protocol) == 0 {
			item.Protocol = "TCP"
		}
		svcPort := v1.ServicePort{
			Name:       strconv.Itoa(item.ContainerPort),
			Protocol:   v1.Protocol(strings.ToUpper(item.Protocol)),
			TargetPort: intstr.FromInt(item.ContainerPort),
			Port:       item.ServicePort,
			NodePort:   item.NodePort,
		}
		service.Spec.Ports = append(service.Spec.Ports, svcPort)
	}
	// Create service
	if _, err := models.GetClientset().CoreV1().Services(service.Namespace).Create(&service); err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Created service " + req.Name + " error"
		go models.Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Created service " + req.Name + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
