package deployment

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"k8s/models"
	"net/http"
)

// RollBack rollbacks deployment
func RollBack(depname string) (int, error) {
	log := models.Log{}
	deploymentsClient := models.GetClientset().AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	// Update Deployment
	result, getErr := deploymentsClient.Get(depname, metav1.GetOptions{}) //Check if service already exist
	if getErr != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Check if" + depname + " exist failed :" + getErr.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, getErr
	}
	baseimage, err := models.Getbaseimage(depname)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result.Spec.Template.Spec.Containers[0].Image = baseimage // change nginx version
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Update " + depname + " failed :" + getErr.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, retryErr
	}
	err = models.Updateimage(depname, baseimage)
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Update " + depname + " failed :" + getErr.Error()
		go models.Insertlog(&log)
	}
	return http.StatusOK, nil
}
