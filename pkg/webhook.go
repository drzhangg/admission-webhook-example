package pkg

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDeserializer()
)

const (
	admissionWebhookAnnotationValidateKey = "admission-webhook-example.qikqiak.com/validate"
	admissionWebhookAnnotationMutateKey = "admission-webhook-example.qikqiak.com/mutate"
	admissionWebhookAnnotationStatusKey = "admission-webhook-example.qikqiak.com/status"

	nameLabel = "app.kubernetes.io/name"
	instanceLabel = "app.kubernetes.io/instance"
	versionLabel = "app.kubernetes.io/version"
	componentLabel = "app.kubernetes.io/component"
	partOfLabel = "app.kubernetes.io/part-of"
	managedByLabel = "app.kubernetes.io/managed-by"

	NA = "not_available"
)

var (
	ignoredNamespaces = []string{
		metav1.NamespaceSystem,
		metav1.NamespacePublic,
	}
	requiredLabels = []string{

	}
)
