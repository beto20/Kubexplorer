package knowledge

type WellKnowPodError string
type WellKnowDeploymentError string
type WellKnowServiceError string
type WellKnowSecretError string

const (
	CRASH_LOOP_BACK_OFF  WellKnowPodError        = "CrashLoopBackOff"
	OOM_KILLED           WellKnowPodError        = "OOMKilled"
	IMAGE_PULL_BACK_OFF  WellKnowPodError        = "ImagePullBackOff"
	ERR_IMAGE_PULL       WellKnowPodError        = "ErrImagePull"
	EVICTED              WellKnowPodError        = "Evicted"
	PENDING              WellKnowPodError        = "Pending"
	UNAVAILABLE_REPLICAS WellKnowDeploymentError = "UnavailableReplicas"
	DEPLOYMENT_PENDING   WellKnowDeploymentError = "DeploymentPending"
	CONNECTION_REFUSED   WellKnowServiceError    = "ConnectionRefused"
	IP_PENDING           WellKnowServiceError    = "IPPending"
	UNAUTHORIZED         WellKnowSecretError     = "Unauthorized"
	INVALID_SECRET_DATA  WellKnowSecretError     = "InvalidSecretData"
)
