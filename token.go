package secrets

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type GithubTokenOption func(context.Context, *githubTokenOptions) error

type githubTokenOptions struct {}

func GithubToken(ctx context.Context, opts ...GithubTokenOption) (string, error) {
	if _, err := exec.LookPath("gh"); err != nil {
		return "", fmt.Errorf("gh is not installed")
	}

	out, err := exec.CommandContext(ctx, "gh", "auth", "token").Output()
	if err != nil {
		return "", fmt.Errorf("failed to obtain token: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}
