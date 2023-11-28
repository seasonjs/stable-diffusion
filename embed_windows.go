// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build windows && amd64

package sd

import (
	_ "embed"
	"errors"
	"golang.org/x/sys/cpu"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

//go:embed deps/windows/sd-abi_avx2.dll
var libStableDiffusionAvx2 []byte

//go:embed deps/windows/sd-abi_avx.dll
var libStableDiffusionAvx []byte

//go:embed deps/windows/sd-abi_avx512.dll
var libStableDiffusionAvx512 []byte

//go:embed deps/windows/sd-abi_cuda12.dll
var libStableDiffusionCuda12 []byte

var libName = "stable-diffusion-*.dll"

func getDl(gpu bool) []byte {
	if gpu {
		info, err := getGPUInfo()
		if err != nil {
			log.Println(err)
		}
		log.Print("GPU: ", info)
	}

	if cpu.X86.HasAVX512 {
		return libStableDiffusionAvx512
	}

	if cpu.X86.HasAVX2 {
		return libStableDiffusionAvx2
	}

	return libStableDiffusionAvx
}

func runPowerShellCommand(command string) (string, error) {
	cmd := exec.Command("powershell", "-Command", command)

	// execute the command and get the output
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func getGPUInfo() (string, error) {
	psCommand := "Get-WmiObject Win32_VideoController"

	output, err := runPowerShellCommand(psCommand)
	if err != nil {
		return "", err
	}
	infos := strings.Split(output, "\r\n")
	re := regexp.MustCompile(`^Name\s+:\s+(.+)`)
	for _, info := range infos {
		match := re.FindStringSubmatch(info)
		if len(match) >= 2 {
			// get the first match Name
			gpuName := match[1]
			return gpuName, nil
		}
	}

	return "", errors.New("no gpu found")
}
