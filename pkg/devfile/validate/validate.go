package validate

import (
	"reflect"

	"github.com/devfile/kubernetes-api/pkg/apis/workspaces/v1alpha1"
	v200 "github.com/ranakan19/parser/pkg/devfile/parser/data/2.0.0"
	"k8s.io/klog"
)

// ValidateDevfileData validates whether sections of devfile are odo compatible
func ValidateDevfileData(data interface{}) error {
	var components []v1alpha1.Component

	typeData := reflect.TypeOf(data)

	// if typeData == reflect.TypeOf(&v100.Devfile100{}) {
	// 	d := data.(*v100.Devfile100)
	// 	components = d.GetComponents()
	// }

	if typeData == reflect.TypeOf(&v200.Devfile200{}) {
		d := data.(*v200.Devfile200)
		components = d.GetComponents()
	}

	// if typeData == reflect.TypeOf(&v210.Devfile210{}) {
	// 	d := data.(*v210.Devfile210)
	// 	components = d.GetComponents()
	// }

	// Validate Components
	if err := ValidateComponents(components); err != nil {
		return err
	}

	// Successful
	klog.V(4).Info("Successfully validated devfile sections")
	return nil

}
