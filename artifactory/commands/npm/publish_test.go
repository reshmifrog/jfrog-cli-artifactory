package npm

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPackageInfoFromTarball(t *testing.T) {
	npmPublish := NewNpmPublishCommand()

	var testCases = []struct {
		filePath       string
		packageName    string
		packageVersion string
	}{
		{
			filePath:       filepath.Join("..", "testdata", "npm", "npm-example-0.0.3.tgz"),
			packageName:    "npm-example",
			packageVersion: "0.0.3",
		}, {
			filePath:       filepath.Join("..", "testdata", "npm", "npm-example-0.0.4.tgz"),
			packageName:    "npm-example",
			packageVersion: "0.0.4",
		}, {
			// Test case for non-standard structure where package.json is in a custom location
			filePath:       filepath.Join("..", "testdata", "npm", "node-package-1.0.0.tgz"),
			packageName:    "nonstandard-package",
			packageVersion: "1.0.0",
		},
	}
	for _, test := range testCases {
		err := npmPublish.readPackageInfoFromTarball(test.filePath)
		assert.NoError(t, err)
		assert.Equal(t, test.packageName, npmPublish.packageInfo.Name)
		assert.Equal(t, test.packageVersion, npmPublish.packageInfo.Version)
	}
}
