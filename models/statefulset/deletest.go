package statefulset

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"Slyph/models"
	"net/http"
	"time"
)

// DeleteStatefulset delete a Statefulset
func DeleteStatefulset(req *models.Delete) (int, error) { // just require req.Name(service name)
	var dp metav1.DeletionPropagation
	dp = metav1.DeletePropagationBackground
	log := models.Log{}
	stclient := models.GetClientset().AppsV1beta2().StatefulSets(apiv1.NamespaceDefault)
	stimp, err := stclient.Get(req.Name, metav1.GetOptions{})
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to get latest version of Statefulset " + err.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	// delete service
	if stimp.Name != " " {
		if err := models.GetClientset().AppsV1beta2().StatefulSets(req.Namespace).Delete(req.Name, &metav1.DeleteOptions{}); err != nil {
			log.Time = models.Gettime()
			log.Etype = "error"
			log.Event = "Failed to delete Statefulset " + req.Name + " " + err.Error()
			go models.Insertlog(&log)
			return http.StatusInternalServerError, err
		}
	}
	for i := 0; i < int(*stimp.Spec.Replicas); i++ {
		if err := models.GetClientset().CoreV1().Pods(req.Namespace).Delete(fmt.Sprintf("%s-%d", req.Name, i), &metav1.DeleteOptions{
			PropagationPolicy: &dp,
		}); err != nil {
			log.Time = models.Gettime()
			log.Etype = "error"
			log.Event = "Failed to delete pod " + req.Name + " " + err.Error()
			go models.Insertlog(&log)
			return http.StatusInternalServerError, err
		}
		time.Sleep(time.Microsecond * 200)
		if err := models.GetClientset().CoreV1().PersistentVolumeClaims(req.Namespace).Delete(fmt.Sprintf("%s-%s-%d", stimp.Spec.VolumeClaimTemplates[0].GetName(), req.Name, i), &metav1.DeleteOptions{}); err != nil {
			log.Time = models.Gettime()
			log.Etype = "error"
			log.Event = "Failed to Delete for PVC " + req.Name + " " + err.Error()
			go models.Insertlog(&log)
			return http.StatusInternalServerError, err
		}
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Delete Statefulset " + req.Name + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
