// Package azure deploys a togo app to Azure (Container Apps) by driving the `az` CLI.
// Select with deploy.provider=azure; requires the az CLI authenticated.
package azure

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/togo-framework/deploy"
	"github.com/togo-framework/togo"
)

func init() { deploy.RegisterDriver("azure", New) }

// New checks the az CLI is present.
func New(_ *togo.Kernel) (deploy.Deployer, error) {
	if _, err := exec.LookPath("az"); err != nil {
		return nil, fmt.Errorf("deploy-azure: the %q CLI is required (install + authenticate it)", "az")
	}
	return &driver{}, nil
}

type driver struct{}

func run(ctx context.Context, name string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	return string(out), err
}

func region(s deploy.Spec, def string) string {
	if s.Region != "" {
		return s.Region
	}
	return def
}

func (d *driver) Provision(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	return d.Deploy(ctx, spec)
}

func (d *driver) Deploy(ctx context.Context, spec deploy.Spec) (*deploy.Result, error) {
	rg := spec.App + "-rg"
	out, err := run(ctx, "az", "containerapp", "up", "--name", spec.App, "--resource-group", rg, "--image", spec.Image, "--target-port", "8080", "--ingress", "external")
	if err != nil {
		return nil, fmt.Errorf("az containerapp up: %v: %s", err, out)
	}
	return &deploy.Result{Message: "deployed to Azure Container Apps", Raw: map[string]any{"out": out}}, nil
}

func (d *driver) Destroy(ctx context.Context, spec deploy.Spec) error {
	_, err := run(ctx, "az", "containerapp", "delete", "--name", spec.App, "--resource-group", spec.App+"-rg", "--yes")
	return err
}

func (d *driver) Status(ctx context.Context, spec deploy.Spec) (*deploy.Status, error) {
	out, err := run(ctx, "az", "containerapp", "show", "--name", spec.App, "--resource-group", spec.App+"-rg", "--query", "properties.runningStatus", "-o", "tsv")
	return &deploy.Status{Healthy: err == nil, Detail: out}, err
}
