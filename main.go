package main

import (
	"log"
	"os"

	acme "github.com/cert-manager/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/cert-manager/cert-manager/pkg/acme/webhook/cmd"
	"k8s.io/client-go/rest"
)

var GroupName = os.Getenv("GROUP_NAME")

func main() {
	cmd.RunWebhookServer(GroupName,
		&exampleSolver{
			name: "pinax-webhook-solver",
		},
	)
}

type exampleSolver struct {
	name string
}

func (e *exampleSolver) Name() string {
	return e.name
}

func (e *exampleSolver) Present(ch *acme.ChallengeRequest) error {
	// Log that a challenge is being requested, but don't take any action
	log.Println("Presenting DNS challenge, but passing to k8s_gateway")
	return nil
}

func (e *exampleSolver) CleanUp(ch *acme.ChallengeRequest) error {
	log.Println("Cleaning up DNS challenge, but passing to k8s_gateway")
	return nil
}

func (e *exampleSolver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	log.Println("Initialize, doing nothing...")
	return nil
}
