package trinity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadWebAppSetting(t *testing.T) {
	const (
		logFilePath = "./examples/config/webapp_template.yml"
		runMode     = "Local"
	)
	s := newSetting(runMode, logFilePath)
	assert.Equal(t, "trinity", s.GetProjectName(), "func exec error , test failed")
	assert.Equal(t, "HTTP", s.GetSetting().Webapp.Type, "func exec error , test failed")
	assert.Equal(t, false, s.GetSetting().ServiceMesh.AutoRegister, "func exec error , test failed")
	assert.Equal(t, "trinity", s.GetSetting().Project, "func exec error , test failed")
	assert.Equal(t, "127.0.0.1", s.GetServiceMeshAddress(), "func exec error , test failed")
	assert.Equal(t, 8500, s.GetServiceMeshPort(), "func exec error , test failed")
}
