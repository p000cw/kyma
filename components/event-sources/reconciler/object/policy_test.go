package object

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	authenticationv1alpha1api "istio.io/api/authentication/v1alpha1"
	authenticationv1alpha1 "istio.io/client-go/pkg/apis/authentication/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	tTarget     = "tRev-private"
	tTargetPort = "http-usermetric"
)

func TestNewPolicy(t *testing.T) {
	policy := NewPolicy(tNs, tName,
		WithTarget(tTarget),
		WithPermissiveMode())

	expectPolicy := &authenticationv1alpha1.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: tNs,
			Name:      tName,
		},
		Spec: authenticationv1alpha1api.Policy{
			Targets: []*authenticationv1alpha1api.TargetSelector{
				{
					Name: tTarget,
					Ports: []*authenticationv1alpha1api.PortSelector{
						{
							Port: &authenticationv1alpha1api.PortSelector_Name{
								Name: tTargetPort,
							},
						},
					},
				},
			},
			Peers: []*authenticationv1alpha1api.PeerAuthenticationMethod{
				{
					Params: &authenticationv1alpha1api.PeerAuthenticationMethod_Mtls{
						Mtls: &authenticationv1alpha1api.MutualTls{
							Mode: authenticationv1alpha1api.MutualTls_PERMISSIVE,
						},
					},
				},
			},
		},
	}

	if d := cmp.Diff(expectPolicy, policy); d != "" {
		t.Errorf("Unexpected diff: (-:expect, +:got) %s", d)
	}
}

func TestApplyExistingPolicyAttributes(t *testing.T) {
	existingPolicy := &authenticationv1alpha1.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       tNs,
			Name:            tName,
			ResourceVersion: "100",
		},
		Spec: authenticationv1alpha1api.Policy{},
	}

	desiredPolicy := NewPolicy(tNs, tName)

	ApplyExistingPolicyAttributes(existingPolicy, desiredPolicy)
	expectedPolicy := &authenticationv1alpha1.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       tNs,
			Name:            tName,
			ResourceVersion: "100",
		},
		Spec: authenticationv1alpha1api.Policy{},
	}

	if d := cmp.Diff(desiredPolicy, expectedPolicy); d != "" {
		t.Errorf("Unexpected diff: (-:expect, +:got) %s", d)
	}
}
