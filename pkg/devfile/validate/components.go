package validate

import (
	"fmt"

	v1 "github.com/devfile/kubernetes-api/pkg/apis/workspaces/v1alpha1"
)

// Errors
var (
	ErrorNoComponents         = "no components present"
	ErrorNoContainerComponent = fmt.Sprintf("odo requires atleast one component of type '%s' in devfile", v1.ContainerComponentType)
)

// ValidateComponents validates all the devfile components
func ValidateComponents(components []v1.Component) error {

	// components cannot be empty
	if len(components) < 1 {
		return fmt.Errorf(ErrorNoComponents)
	}

	// Check if component of type container  is present
	isContainerComponentPresent := false
	for _, component := range components {
		if component.Container != nil {
			isContainerComponentPresent = true
			break
		}
	}

	if !isContainerComponentPresent {
		return fmt.Errorf(ErrorNoContainerComponent)
	}

	// Successful
	return nil
}

func validatev1Components(components []v1.Component) error {
	if len(components) < 1 {
		return fmt.Errorf(ErrorNoComponents)
	}

	// Check if component of type container  is present
	isContainerComponentPresent := false
	for _, component := range components {
		if component.Container != nil {
			isContainerComponentPresent = true
			break
		}
	}

	if !isContainerComponentPresent {
		return fmt.Errorf(ErrorNoContainerComponent)
	}

	// Successful
	return nil
}
