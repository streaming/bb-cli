package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/buildbeaver/buildbeaver/common/models"
)

const dockerInDockerYAML = `
version: "%s"
jobs:
  - name: my-job
    docker:
      image: golang:1.19
%s
    steps:
      - name: my-step
        commands:
          - echo hello
`

func TestParseDockerInDocker(t *testing.T) {
	for _, version := range []string{"0.2", "0.3"} {
		t.Run("version "+version, func(t *testing.T) {
			t.Run("defaults to false when not set", func(t *testing.T) {
				config := []byte(formatDockerInDockerYAML(version, ""))
				parser := NewBuildDefinitionParser(ParserLimits{})
				build, err := parser.Parse(config, models.ConfigTypeYAML)
				require.NoError(t, err)
				require.Len(t, build.Jobs, 1)
				require.False(t, build.Jobs[0].DockerInDocker)
			})

			t.Run("parses true when set", func(t *testing.T) {
				config := []byte(formatDockerInDockerYAML(version, "      docker_in_docker: true"))
				parser := NewBuildDefinitionParser(ParserLimits{})
				build, err := parser.Parse(config, models.ConfigTypeYAML)
				require.NoError(t, err)
				require.Len(t, build.Jobs, 1)
				require.True(t, build.Jobs[0].DockerInDocker)
			})

			t.Run("errors on non-bool value", func(t *testing.T) {
				config := []byte(formatDockerInDockerYAML(version, "      docker_in_docker: \"yes\""))
				parser := NewBuildDefinitionParser(ParserLimits{})
				_, err := parser.Parse(config, models.ConfigTypeYAML)
				require.Error(t, err)
			})
		})
	}
}

func formatDockerInDockerYAML(version string, extraDockerLine string) string {
	return fmt.Sprintf(dockerInDockerYAML, version, extraDockerLine)
}
