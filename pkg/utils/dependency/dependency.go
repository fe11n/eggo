/******************************************************************************
 * Copyright (c) Huawei Technologies Co., Ltd. 2021. All rights reserved.
 * eggo licensed under the Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 * PURPOSE.
 * See the Mulan PSL v2 for more details.
 * Author: zhangxiaoyu
 * Create: 2021-06-25
 * Description: eggo install implement
 ******************************************************************************/

package dependency

import (
	"fmt"
	"strings"

	"isula.org/eggo/pkg/api"
	"isula.org/eggo/pkg/utils/runner"
	"isula.org/eggo/pkg/utils/template"
)

type dependency interface {
	Install(r runner.Runner) error
	Remove(r runner.Runner) error
}

// install dependency by repo
func runRepoCommand(r runner.Runner, software []*api.PackageConfig, command string) error {
	join := ""
	for _, s := range software {
		join += s.Name + " "
	}
	_, err := r.RunCommand(fmt.Sprintf("sudo -E /bin/sh -c \"%s -y %s\"", command, join))
	return err
}

type dependencyApt struct {
	software []*api.PackageConfig
}

func (da *dependencyApt) Install(r runner.Runner) error {
	command := "apt install"
	if err := runRepoCommand(r, da.software, command); err != nil {
		return fmt.Errorf("apt install failed: %v", err)
	}

	return nil
}

func (da *dependencyApt) Remove(r runner.Runner) error {
	command := "apt remove"
	if err := runRepoCommand(r, da.software, command); err != nil {
		return fmt.Errorf("apt remove failed: %v", err)
	}

	return nil
}

type dependencyYum struct {
	software []*api.PackageConfig
}

func (dy *dependencyYum) Install(r runner.Runner) error {
	command := "yum install"
	if err := runRepoCommand(r, dy.software, command); err != nil {
		return fmt.Errorf("yum install by yum failed: %v", err)
	}

	return nil
}

func (dy *dependencyYum) Remove(r runner.Runner) error {
	command := "yum remove"
	if err := runRepoCommand(r, dy.software, command); err != nil {
		return fmt.Errorf("yum remove failed: %v", err)
	}

	return nil
}

// install dependency by pkg
func runPkgCommand(r runner.Runner, software []*api.PackageConfig, srcPath, command string) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("sudo -E /bin/sh -c \"cd %s && %s ", srcPath, command))
	for _, s := range software {
		sb.WriteString(fmt.Sprintf("%s* ", s.Name))
	}
	sb.WriteString("\"")

	_, err := r.RunCommand(sb.String())
	return err
}

type dependencyRpm struct {
	srcPath  string
	software []*api.PackageConfig
}

func (dr *dependencyRpm) Install(r runner.Runner) error {
	command := "rpm -ivh --force --nodeps"
	if err := runPkgCommand(r, dr.software, dr.srcPath, command); err != nil {
		return fmt.Errorf("rpm install failed: %v", err)
	}

	return nil
}

func (dr *dependencyRpm) Remove(r runner.Runner) error {
	command := "yum remove -y"
	if err := runPkgCommand(r, dr.software, dr.srcPath, command); err != nil {
		return fmt.Errorf("yum remove rpm pkgs failed: %v", err)
	}

	return nil
}

type dependencyDeb struct {
	srcPath  string
	software []*api.PackageConfig
}

func (dd *dependencyDeb) Install(r runner.Runner) error {
	command := "dpkg --force-all -i"
	if err := runPkgCommand(r, dd.software, dd.srcPath, command); err != nil {
		return fmt.Errorf("dpkg install failed: %v", err)
	}

	return nil
}

func (dd *dependencyDeb) Remove(r runner.Runner) error {
	command := "apt remove -y"
	if err := runPkgCommand(r, dd.software, dd.srcPath, command); err != nil {
		return fmt.Errorf("apt remove deb pkgs failed: %v", err)
	}

	return nil
}

// install file and dir
type dependencyFileDir struct {
	executable bool
	srcPath    string
	software   []*api.PackageConfig
}

func (df *dependencyFileDir) Install(r runner.Runner) error {
	shell := `
#!/bin/bash
cd {{ .srcPath }}

{{- if .executable }}
{{- range $i, $v := .software }}
chmod +x {{ $v.Name }}
{{- end }}
{{- end }}

{{- range $i, $v := .software }}
if [ ! -e {{ JoinPath $v.Dst $v.Name }} ]; then
    mkdir -p {{ $v.Dst }} && cp -r {{ $v.Name }} {{ $v.Dst }}
fi
{{- end }}
`
	datastore := make(map[string]interface{})
	datastore["srcPath"] = df.srcPath
	datastore["software"] = df.software
	datastore["executable"] = df.executable

	shellStr, err := template.TemplateRender(shell, datastore)
	if err != nil {
		return err
	}

	_, err = r.RunShell(shellStr, "install_FD")
	if err != nil {
		return err
	}

	return nil
}

func (df *dependencyFileDir) Remove(r runner.Runner) error {
	var sb strings.Builder
	sb.WriteString("sudo -E /bin/sh -c \"")
	for _, s := range df.software {
		path := fmt.Sprintf("%s/%s", s.Dst, s.Name)
		if !CheckPath(path) {
			return fmt.Errorf("path %s not in White List and cannot remove", path)
		}

		sb.WriteString(fmt.Sprintf("rm -rf %s ; ", path))
	}
	sb.WriteString("\"")

	if _, err := r.RunCommand(sb.String()); err != nil {
		return fmt.Errorf("rm dependency failed: %v", err)
	}

	return nil
}

type dependencyImage struct {
	srcPath string
	client  string
	command string
	image   []*api.PackageConfig
}

func NewDependencyImage(srcPath, client, command string, image []*api.PackageConfig) *dependencyImage {
	return &dependencyImage{
		srcPath: srcPath,
		client:  client,
		command: command,
		image:   image,
	}
}

func (di *dependencyImage) Install(r runner.Runner) error {
	var sb strings.Builder
	sb.WriteString("sudo -E /bin/sh -c \"")
	for _, i := range di.image {
		sb.WriteString(fmt.Sprintf("%s %s/%s && ", di.command, di.srcPath, i.Name))
	}
	sb.WriteString("echo success\"")

	if _, err := r.RunCommand(sb.String()); err != nil {
		return fmt.Errorf("%s load image failed: %v", di.client, err)
	}

	return nil
}

func (di *dependencyImage) Remove(r runner.Runner) error {
	// nothing to do
	return nil
}

type dependencyYaml struct {
	srcPath    string
	kubeconfig string
	yaml       []*api.PackageConfig
}

func NewDependencyYaml(srcPath, kubeconfig string, yaml []*api.PackageConfig) *dependencyYaml {
	return &dependencyYaml{
		srcPath:    srcPath,
		kubeconfig: kubeconfig,
		yaml:       yaml,
	}
}

func (dy *dependencyYaml) Install(r runner.Runner) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("sudo -E /bin/sh -c \"export KUBECONFIG=%s ", dy.kubeconfig))
	for _, y := range dy.yaml {
		if strings.HasPrefix(y.Name, "http://") || strings.HasPrefix(y.Name, "https://") {
			sb.WriteString(fmt.Sprintf("&& kubectl apply -f %s ", y.Name))
			continue
		}
		sb.WriteString(fmt.Sprintf("&& kubectl apply -f %s/%s ", dy.srcPath, y.Name))
	}
	sb.WriteString("\"")

	if _, err := r.RunCommand(sb.String()); err != nil {
		return fmt.Errorf("kubectl apply yaml failed: %v", err)
	}

	return nil
}

func (dy *dependencyYaml) Remove(r runner.Runner) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("sudo -E /bin/sh -c \"export KUBECONFIG=%s ", dy.kubeconfig))
	for _, y := range dy.yaml {
		if strings.HasPrefix(y.Name, "http://") || strings.HasPrefix(y.Name, "https://") {
			sb.WriteString(fmt.Sprintf("&& kubectl delete -f %s ", y.Name))
			continue
		}
		sb.WriteString(fmt.Sprintf("&& kubectl delete -f %s/%s ", dy.srcPath, y.Name))
	}
	sb.WriteString("\"")

	if _, err := r.RunCommand(sb.String()); err != nil {
		return fmt.Errorf("kubectl delete yaml failed: %v", err)
	}

	return nil
}

type dependencyShell struct {
	envs    []string
	srcPath string
	shell   []*api.PackageConfig
}

func NewDependencyShell(srcPath string, shell []*api.PackageConfig) *dependencyShell {
	return &dependencyShell{
		srcPath: srcPath,
		shell:   shell,
	}
}

func (ds *dependencyShell) Install(r runner.Runner) error {
	shellTemplate := `
#!/bin/bash
{{- range $i, $v := .Envs }}
export {{ $v }}
{{- end }}

{{- $tout := .Timeouts }}
{{- range $i, $v := .Shells }}
chmod +x {{ $v }} && timeout -s SIGKILL {{index $tout $i}} {{ $v }} > /dev/null
if [ $? -ne 0 ]; then
	echo "run {{ $v }} failed"
	exit 1
fi
{{- end }}

exit 0
`
	datastore := map[string]interface{}{}
	datastore["Envs"] = ds.envs
	var shells []string
	var timeouts []string
	for _, s := range ds.shell {
		shells = append(shells, fmt.Sprintf("%s/%s", ds.srcPath, s.Name))
		timeout := s.TimeOut
		if timeout == "" {
			timeout = "30s"
		}
		timeouts = append(timeouts, timeout)
	}
	datastore["Shells"] = shells
	datastore["Timeouts"] = timeouts

	parsedShell, err := template.TemplateRender(shellTemplate, datastore)
	if err != nil {
		return err
	}

	if _, err := r.RunShell(parsedShell, "exechook"); err != nil {
		return fmt.Errorf("hook execute failed: %v", err)
	}

	return nil
}

func (ds *dependencyShell) Remove(r runner.Runner) error {
	// nothing to do
	return nil
}
