package docker

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/buildbeaver/buildbeaver/runner/runtime"
)

func TestPrepareLinuxContainerConfigDockerSocketBind(t *testing.T) {
	for _, dockerInDocker := range []bool{false, true} {
		r := &Runtime{
			config: Config{
				Config: runtime.Config{
					RuntimeID:    "test-runtime",
					StagingDir:   t.TempDir(),
					WorkspaceDir: t.TempDir(),
				},
				DockerInDocker: dockerInDocker,
			},
		}
		config, err := r.prepareLinuxContainerConfig(context.Background())
		require.NoError(t, err)
		require.Equal(t, dockerInDocker, containsBind(config.Binds, "/var/run/docker.sock:/var/run/docker.sock"))
	}
}

func TestPrepareWindowsContainerConfigDockerPipeBind(t *testing.T) {
	for _, dockerInDocker := range []bool{false, true} {
		r := &Runtime{
			config: Config{
				Config: runtime.Config{
					RuntimeID:    "test-runtime",
					StagingDir:   t.TempDir(),
					WorkspaceDir: t.TempDir(),
				},
				DockerInDocker: dockerInDocker,
			},
		}
		config, err := r.prepareWindowsContainerConfig(context.Background())
		require.NoError(t, err)
		require.Equal(t, dockerInDocker, containsBind(config.Binds, "\\\\.\\pipe\\docker_engine:\\\\.\\pipe\\docker_engine"))
	}
}

func containsBind(binds []string, bind string) bool {
	for _, b := range binds {
		if b == bind {
			return true
		}
	}
	return false
}
