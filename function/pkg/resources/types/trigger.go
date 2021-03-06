package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Attributes struct {
	Eventtypeversion string `json:"eventtypeversion"`
	Source           string `json:"source"`
	Type             string `json:"type"`
}

type TriggerFilter struct {
	Attributes Attributes `json:"attributes"`
}

type TriggerReference struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type TriggerSubscriber struct {
	Reference TriggerReference `json:"ref"`
}

type TriggerSpec struct {
	Filter     TriggerFilter     `json:"filter"`
	Subscriber TriggerSubscriber `json:"subscriber"`
}

type Trigger struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec `json:"spec"`
}

func (t Trigger) IsReference(name, namespace string) bool {
	return t.Spec.Subscriber.Reference.Kind == "Service" &&
		t.Spec.Subscriber.Reference.Name == name &&
		t.Spec.Subscriber.Reference.Namespace == namespace
}
