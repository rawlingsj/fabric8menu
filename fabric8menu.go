/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
	"github.com/fabric8io/gofabric8/cmds"
)

func main() {
	runtime.LockOSThread()         // must be the first statement in main - see below
	gallium.Loop(os.Args, onReady) // must be called from main function

}

func onReady(browser *gallium.App) {
	gallium.AddStatusItem(
		20,
		"fabric8",
		true,
		gallium.MenuItem{
			Title:   "Install",
			OnClick: handleInstall,
		},
		gallium.MenuItem{
			Title:   "Deploy",
			OnClick: handleDeploy,
		},
		gallium.MenuItem{
			Title:   "Start",
			OnClick: handleStart,
		},
		gallium.MenuItem{
			Title:   "Stop",
			OnClick: handleStop,
		},
		gallium.MenuItem{
			Title:   "Status",
			OnClick: handleStatus,
		},
	)
}

func handleInstall() {
	cmds.Install(false)
	log.Println("installing")
}

func handleDeploy() {
	log.Println("deploying")
}

func handleStart() {
	log.Println("starting")
}

func handleStop() {
	log.Println("stoping")
}

func handleStatus() {
	log.Println("status")
}
