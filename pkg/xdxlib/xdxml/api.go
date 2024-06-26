/*
 * Copyright (c) 2024, XDXCT CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package xdxml

import (
	"github.com/chen-mao/go-xdxml/pkg/xdxml"
)

//go:generate moq -out xdxml_mock.go . Interface
type Interface interface {
	DeviceGetCount() (int, Return)
	DeviceGetHandleByIndex(Index int) (Device, Return)
	Init() Return
	Shutdown() Return
}

type Device interface {
	GetArchitecture() (string, Return)
	GetMinorNumber() (int, Return)
	GetPciInfo() (PciInfo, Return)
	GetUUID() (string, Return)
	GetMemoryInfo() (Memory, Return)
}

type Return xdxml.Return

type PciInfo xdxml.Pci_info

type Memory xdxml.Memory
