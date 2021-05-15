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
 * Author: haozi007
 * Create: 2021-05-13
 * Description: nodemanager testcase
 ******************************************************************************/

package nodemanager

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"gitee.com/openeuler/eggo/pkg/clusterdeployment"
	"gitee.com/openeuler/eggo/pkg/utils/runner"
	"github.com/sirupsen/logrus"
)

type MockRunner struct {
}

func (m *MockRunner) Copy(src, dst string) error {
	logrus.Infof("copy %s to %s", src, dst)
	return nil
}

func (m *MockRunner) RunCommand(cmd string) error {
	logrus.Infof("run command: %s", cmd)
	return nil
}

func (m *MockRunner) Reconnect() error {
	logrus.Infof("reconnect")
	return nil
}

func (m *MockRunner) Close() {
	logrus.Infof("close")
}

type MockTask struct {
	// some need data
	name   string
	lock   sync.RWMutex
	lables map[string]string
}

func (m *MockTask) Run(r runner.Runner) error {
	rand.Seed(time.Now().UnixNano())

	err := r.Copy("/home/data", "/data")
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	err = r.RunCommand(m.name + " run 'top'")
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	r.Reconnect()

	return err
}

func (m *MockTask) Name() string {
	return m.name
}

func (m *MockTask) AddLabels(key, lable string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.lables[key] = lable
}

func (m *MockTask) GetLable(key string) string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	l, ok := m.lables[key]
	if ok {
		return l
	}

	return ""
}

func addNodes() {
	hcf1 := &clusterdeployment.HostConfig{
		Arch:     "x86_64",
		Name:     "master",
		Address:  "192.168.0.1",
		Port:     22,
		UserName: "root",
		Password: "123456",
		Type:     "master",
	}
	hcf2 := &clusterdeployment.HostConfig{
		Arch:     "arm64",
		Name:     "work",
		Address:  "192.168.0.2",
		Port:     22,
		UserName: "root",
		Password: "123456",
		Type:     "work",
	}
	r := &MockRunner{}
	RegisterNode(hcf1, r)
	RegisterNode(hcf2, r)
}

func releaseNodes(nodes []string) {
	for _, id := range nodes {
		UnRegisterNode(id)
	}
}

func TestRunTaskOnNodes(t *testing.T) {
	addNodes()
	tt := &MockTask{name: "precheck", lables: make(map[string]string)}
	nodes := []string{"192.168.0.1", "192.168.0.2"}
	err := RunTaskOnNodes(tt, nodes)
	if err != nil {
		t.Fatalf("run task on ondes failed: %v\n", err)
	}

	err = WaitTaskOnNodesFinished(tt, nodes, time.Second*30)
	if err != nil {
		t.Fatalf("run task on ondes failed: %v\n", err)
	}
	releaseNodes(nodes)
}

func TestRunTaskOnAll(t *testing.T) {
	addNodes()
	tt := &MockTask{name: "precheck", lables: make(map[string]string)}
	err := RunTaskOnAll(tt)
	if err != nil {
		t.Fatalf("run task on all node failed: %v\n", err)
	}
	err = WaitTaskOnAllFinished(tt, time.Second*30)
	if err != nil {
		t.Fatal("run task on all ondes failed\n")
	}
	UnRegisterAllNodes()
}