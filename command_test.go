/*
 * Copyright (c) 2019. aberic - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gnomon

import "testing"

func TestCommandCommon_Exec(t *testing.T) {
	if line, cmd, strArr, err := CommandExec("ls", "-l"); err != nil {
		t.Skip(err)
	} else {
		t.Log("line =", line, ", strArr =", strArr, ", pid =", cmd.Process.Pid)
	}

	_, _, _, err := CommandExec("lss", "-l")
	t.Skip(err)
}

func TestCommandCommon_ExecSilent(t *testing.T) {
	if line, cmd, strArr, err := CommandExecSilent("ls", "-l"); err != nil {
		t.Skip(err)
	} else {
		t.Log("line =", line, ", strArr =", strArr, ", pid =", cmd.Process.Pid)
	}

	_, _, _, err := CommandExecSilent("lss", "-l")
	t.Skip(err)
}

func TestCommandCommon_ExecTail(t *testing.T) {
	if line, cmd, strArr, err := CommandExecTail("ls", "-l"); err != nil {
		t.Skip(err)
	} else {
		t.Log("line =", line, ", strArr =", strArr, ", pid =", cmd.Process.Pid)
	}

	_, _, _, err := CommandExecTail("lss", "-l")
	t.Skip(err)
}

func TestCommandCommon_ExecAsync(t *testing.T) {
	commandAsync := make(chan *CommandAsync, 1)
	var keep bool
	go CommandExecAsync(commandAsync, "ls", "-l")
	keep = true
	for {
		ca := <-commandAsync
		if nil != ca.Err {
			t.Skip(ca.Err)
			keep = false
		}
		if ca.Tail == "OFF" {
			keep = false
			t.Log("command over")
		}
		t.Log("tail", ca.Tail)
		if !keep {
			break
		}
	}

	go CommandExecAsync(commandAsync, "lls", "-l")
	keep = true
	for {
		ca := <-commandAsync
		if nil != ca.Err {
			t.Skip(ca.Err)
			keep = false
		}
		if ca.Tail == "OFF" {
			keep = false
			t.Log("command over")
		}
		t.Log("tail", ca.Tail)
		if !keep {
			break
		}
	}
}
