package env

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

func execCmd(cmdName string, cmdArgs []string, cmdExec string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf(cmdExec+" cmd out | %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}

}

func (env *Env) setGSettingsHost() {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "host", env.Host}
	execCmd(cmdName, cmdArgs, "gsettingsHost")
}

func (env *Env) setGSettingsPort() {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "port", env.Port}
	execCmd(cmdName, cmdArgs, "gsettingsPort")
}

func (env *Env) setGSettingsMode() {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy", "mode", "manual"}
	execCmd(cmdName, cmdArgs, "gsettingsMode")
}

func (env *Env) setEnvVariableHttpProxy() {
	url := "http://" + env.Host + ":" + env.Port
	os.Setenv("http_proxy", url)
	os.Setenv("HTTP_PROXY", url)

}
func (env *Env) setEnvVariableHttpsProxy() {
	url := "https://" + env.Host + ":" + env.Port
	os.Setenv("https_proxy", url)
	os.Setenv("HTTPS_PROXY", url)
}

func (env *UnSetEnv) unsetEnvVariableHttpProxy() {
	os.Unsetenv("http_proxy")
	os.Unsetenv("HTTP_PROXY")

}
func (env *UnSetEnv) unsetEnvVariableHttpsProxy() {
	os.Unsetenv("https_proxy")
	os.Unsetenv("HTTPS_PROXY")
}

func (env *UnSetEnv) unsetGSettingsMode() {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy", "mode", "none"}
	execCmd(cmdName, cmdArgs, "gsettingsMode")
}

func (env *UnSetEnv) unsetGSettingsHost() {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "host", ""}
	execCmd(cmdName, cmdArgs, "gsettingsHost")
}
func (env *UnSetEnv) unsetGSettingsPort() {
	cmdName := "gsettings"
	cmdArgs := []string{"set", "org.gnome.system.proxy.http", "port", "0"}
	execCmd(cmdName, cmdArgs, "gsettingsPort")
}

func (env *Env) gBashConfigTmp() {
	cmdName := "echo"
	cmdArgs := []string{"-n", "", ">", "bash_config.tmp"}
	execCmd(cmdName, cmdArgs, "gBashConfigTmp")
}

func (env *Env) GSettings() {
	env.setGSettingsHost()
	env.setGSettingsPort()
	env.setGSettingsMode()
	env.gBashConfigTmp()
	env.setEnvVariableHttpProxy()
	env.setEnvVariableHttpsProxy()
}

func (env *UnSetEnv) UnsetGSettings() {
	env.unsetGSettingsHost()
	env.unsetGSettingsPort()
	env.unsetGSettingsMode()
}
