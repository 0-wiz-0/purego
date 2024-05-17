// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 The Ebitengine Authors

package purego

// Constants as defined in https://github.com/NetBSD/src/blob/trunk/include/dlfcn.h
const (
	RTLD_DEFAULT  = ^uintptr(0) - 2 // Pseudo-handle for dlsym so search for any loaded symbol
	RTLD_LAZY     = 0x00001         // Relocations are performed at an implementation-dependent time.
	RTLD_NOW      = 0x00002         // Relocations are performed when the object is loaded.
	RTLD_GLOBAL   = 0x00100         // All symbols are available for relocation processing of other modules.
	RTLD_LOCAL    = 0x00200         // All symbols are not made available for relocation processing by other modules.
	RTLD_NODELETE = 0x01000         // Do not remove members.
	RTLD_NOLOAD   = 0x02000         // Do not load if already loaded.
)
