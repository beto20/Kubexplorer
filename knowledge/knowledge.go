package knowledge

import "fmt"

func ErrorSource(source SourceObject, errorType string) {

	if source == PODS {
		podSolutions(errorType)
	} else if source == DEPLOYMENTS {
		deploymentSolutions(errorType)
	} else if source == SERVICES {
		serviceSolutions(errorType)
	} else if source == SECRETS {
		secretSolutions(errorType)
	}
}

func podSolutions(errorType string) {
	switch errorType {
	case string(CRASH_LOOP_BACK_OFF):
		rollbackImage()
	case string(OOM_KILLED):
		rollbackImage()
	case string(IMAGE_PULL_BACK_OFF):
		rollbackImage()
	case string(ERR_IMAGE_PULL):
		rollbackImage()
	case string(EVICTED):
		rollbackImage()
	case string(PENDING):
		rollbackImage()
	}
}

func deploymentSolutions(errorType string) {
	switch errorType {
	case string(UNAVAILABLE_REPLICAS):
		rollbackImage()
	case string(DEPLOYMENT_PENDING):
		rollbackImage()
	}
}

func serviceSolutions(errorType string) {
	switch errorType {
	case string(CONNECTION_REFUSED):
		rollbackImage()
	case string(IP_PENDING):
		rollbackImage()
	}
}

func secretSolutions(errorType string) {
	switch errorType {
	case string(UNAUTHORIZED):
		rollbackImage()
	case string(INVALID_SECRET_DATA):
		rollbackImage()
	}
}

func rollbackImage() {
	fmt.Print("rollback")
}

func calibrateResources() {

	// if errorType == ErrorsWellKnow(OOMKilled) {
	// 	increaseMemory()
	// } else if errorType == ErrorsWellKnow(Pending) {
	// 	increaseCPU()
	// } else {
	// 	increaseMemory()
	// 	increaseCPU()
	// }

}

func increaseMemory() {
	fmt.Print("memory increase")
}

func increaseCPU() {
	fmt.Print("cpu increase")
}
