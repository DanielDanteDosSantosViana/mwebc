package env

import (
	"errors"
	"os"
	"os/exec"

	"github.com/Sirupsen/logrus"
)

type Env struct {
	Host string
	Port string
}

type UnSetEnv struct {
}

func NewEnv(host string, port string) *Env {
	return &Env{host, port}
}

func NewUnsetEnv() *UnSetEnv {
	return &UnSetEnv{}
}

func execCmd(cmdName string, cmdArgs []string, cmdExec string) error {
	cmd := exec.Command(cmdName, cmdArgs...)
	_, err := cmd.StdoutPipe()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cmd ":    cmdName,
			"cmdArgs": cmdArgs,
		}).Error(err)
		return errors.New("Error creating StdoutPipe for Cmd")
	}

	err = cmd.Start()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cmd ":    cmdName,
			"cmdArgs": cmdArgs,
		}).Error(err)
		return errors.New("Error starting Cmd")
	}

	err = cmd.Wait()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cmd ":    cmdName,
			"cmdArgs": cmdArgs,
		}).Error(err)
		return errors.New("Error waiting for Cmd")
	}
	return nil
}

func (env *Env) setGSettingsHost() error {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "host", env.Host}
	err := execCmd(cmdName, cmdArgs, "gsettingsHost")
	if err != nil {
		return err
	}
	return nil
}

func (env *Env) setGSettingsPort() error {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "port", env.Port}
	err := execCmd(cmdName, cmdArgs, "gsettingsPort")
	if err != nil {
		return err
	}
	return nil
}

func (env *Env) setGSettingsMode() error {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy", "mode", "manual"}
	err := execCmd(cmdName, cmdArgs, "gsettingsMode")
	if err != nil {
		return err
	}
	return nil
}

func (env *Env) setEnvVariableHttpProxy() error {
	url := "http://" + env.Host + ":" + env.Port
	err := os.Setenv("http_proxy", url)
	if err != nil {
		return err
	}
	err = os.Setenv("HTTP_PROXY", url)
	if err != nil {
		return err
	}
	return nil
}
func (env *Env) setEnvVariableHttpsProxy() error {
	url := "https://" + env.Host + ":" + env.Port
	err := os.Setenv("https_proxy", url)
	if err != nil {
		return err
	}
	err = os.Setenv("HTTPS_PROXY", url)
	if err != nil {
		return err
	}
	return nil
}

func (env *UnSetEnv) unsetEnvVariableHttpProxy() error {
	err := os.Unsetenv("http_proxy")
	if err != nil {
		return err
	}
	err = os.Unsetenv("HTTP_PROXY")
	if err != nil {
		return err
	}
	return nil
}
func (env *UnSetEnv) unsetEnvVariableHttpsProxy() error {
	err := os.Unsetenv("https_proxy")
	if err != nil {
		return err
	}
	err = os.Unsetenv("HTTPS_PROXY")
	if err != nil {
		return err
	}
	return nil
}

func (env *UnSetEnv) unsetGSettingsMode() error {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy", "mode", "none"}
	err := execCmd(cmdName, cmdArgs, "gsettingsMode")
	if err != nil {
		return err
	}
	return nil
}

func (env *UnSetEnv) unsetGSettingsHost() error {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "host", ""}
	err := execCmd(cmdName, cmdArgs, "gsettingsHost")
	if err != nil {
		return err
	}
	return nil
}
func (env *UnSetEnv) unsetGSettingsPort() error {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "port", "0"}
	err := execCmd(cmdName, cmdArgs, "gsettingsPort")
	if err != nil {
		return err
	}
	return nil
}

func (env *Env) gBashConfigTmp() error {
	cmdName := "echo"
	cmdArgs := []string{"-n", "", ">", "bash_config.tmp"}
	err := execCmd(cmdName, cmdArgs, "gBashConfigTmp")
	if err != nil {
		return err
	}
	return nil
}

func (env *Env) GSettings() error {
	err := env.setGSettingsHost()
	if err != nil {
		return err
	}
	err = env.setGSettingsPort()
	if err != nil {
		return err
	}
	err = env.setGSettingsMode()
	if err != nil {
		return err
	}
	err = env.gBashConfigTmp()
	if err != nil {
		return err
	}
	err = env.setEnvVariableHttpProxy()
	if err != nil {
		return err
	}
	err = env.setEnvVariableHttpsProxy()
	if err != nil {
		return err
	}
	return nil
}

func (env *UnSetEnv) UnsetGSettings() error {
	err := env.unsetGSettingsHost()
	if err != nil {
		return err
	}
	err = env.unsetGSettingsPort()
	if err != nil {
		return err
	}
	err = env.unsetGSettingsMode()
	if err != nil {
		return err
	}
	return nil
}
