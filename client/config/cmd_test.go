package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/flags"
)

var ErrWrongNumberOfArgs = fmt.Errorf("wrong number of arguments")

func Test_runConfigCmdTwiceWithShorterNodeValue(t *testing.T) {
	// Prepare environment
	t.Parallel()
	configHome, cleanup := tmpDir(t)
	defer cleanup()
	_ = os.RemoveAll(filepath.Join(configHome, "config"))
	viper.Set(flags.FlagHome, configHome)

	// Init command config
	cmd := Cmd()
	assert.NotNil(t, cmd)

	err := cmd.RunE(cmd, []string{"node", "tcp://localhost:26657"})
	assert.Nil(t, err)

	err = cmd.RunE(cmd, []string{"node", "--get"})
	assert.Nil(t, err)

	err = cmd.RunE(cmd, []string{"node", "tcp://local:26657"})
	assert.Nil(t, err)

	err = cmd.RunE(cmd, []string{"node", "--get"})
	assert.Nil(t, err)

	err = cmd.RunE(cmd, nil)
	assert.Nil(t, err)

	//err = cmd.RunE(cmd, []string{"invalidKey", "--get"})
	//require.Equal(t, err, errUnknownConfigKey("invalidKey"))

	err = cmd.RunE(cmd, []string{"invalidArg1"})
	require.Equal(t, err, ErrWrongNumberOfArgs)

	//err = cmd.RunE(cmd, []string{"invalidKey", "invalidValue"})
	//require.Equal(t, err, errUnknownConfigKey("invalidKey"))

}

func tmpDir(t *testing.T) (string, func()) {
	dir, err := ioutil.TempDir("", t.Name()+"_")
	require.NoError(t, err)
	return dir, func() { _ = os.RemoveAll(dir) }
}