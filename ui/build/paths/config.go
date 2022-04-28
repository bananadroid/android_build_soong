// Copyright 2018 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package paths

import "runtime"

type PathConfig struct {
	// Whether to create the symlink in the new PATH for this tool.
	Symlink bool

	// Whether to log about usages of this tool to the soong.log
	Log bool

	// Whether to exit with an error instead of invoking the underlying tool.
	Error bool

	// Whether we use a linux-specific prebuilt for this tool. On Darwin,
	// we'll allow the host executable instead.
	LinuxOnlyPrebuilt bool
}

// These binaries can be run from $PATH, nonhermetically. There should be as
// few as possible of these, since this means that the build depends on tools
// that are not shipped in the source tree and whose behavior is therefore
// unpredictable.
var Allowed = PathConfig{
	Symlink: true,
	Log:     false,
	Error:   false,
}

// This tool is specifically disallowed and calling it will result in an
// "executable no found" error.
var Forbidden = PathConfig{
	Symlink: false,
	Log:     true,
	Error:   true,
}

// This tool is allowed, but access to it will be logged.
var Log = PathConfig{
	Symlink: true,
	Log:     true,
	Error:   false,
}

// The configuration used if the tool is not listed in the config below.
// Currently this will create the symlink, but log and error when it's used. In
// the future, I expect the symlink to be removed, and this will be equivalent
// to Forbidden. This applies to every tool not specifically mentioned in the
// configuration.
var Missing = PathConfig{
	Symlink: true,
	Log:     true,
	Error:   true,
}

// This is used for binaries for which we have prebuilt versions, but only for
// Linux. Thus, their execution from $PATH is only allowed on Mac OS.
var LinuxOnlyPrebuilt = PathConfig{
	Symlink:           false,
	Log:               true,
	Error:             true,
	LinuxOnlyPrebuilt: true,
}

func GetConfig(name string) PathConfig {
	if config, ok := Configuration[name]; ok {
		return config
	}
	return Missing
}

// This list specifies whether a particular binary from $PATH is allowed to be
// run during the build. For more documentation, see path_interposer.go .
var Configuration = map[string]PathConfig{
	"bash":    Allowed,
	"bison":   Allowed,
	"brotli":  Allowed,
	"ccache":  Allowed,
	"cpio":    Allowed,
	"curl":    Allowed,
	"date":    Allowed,
	"depmod":  Allowed,
	"dd":      Allowed,
	"diff":    Allowed,
	"dlv":     Allowed,
	"expr":    Allowed,
	"flex":    Allowed,
	"flock":   Allowed,
	"fuser":   Allowed,
	"getopt":  Allowed,
	"git":     Allowed,
	"grep":    Allowed,
	"hexdump": Allowed,
	"jar":     Allowed,
	"java":    Allowed,
	"javap":   Allowed,
	"ld.lld":  Allowed,
	"llvm-ar": Allowed,
	"locale":  Allowed,
	"lsof":    Allowed,
	"m4":      Allowed,
	"nproc":   Allowed,
	"numfmt":  Allowed,
	"openssl": Allowed,
	"patch":   Allowed,
	"perl":    Allowed,
	"pstree":  Allowed,
	"python3": Allowed,
	"repo":    Allowed,
	"rsync":   Allowed,
	"sh":      Allowed,
	"soong_zip": Allowed,
	"stubby":  Allowed,
	"tar":     Allowed,
	"tr":      Allowed,
	"unzip":   Allowed,
	"zcat":    Allowed,
	"zip":     Allowed,

	"python3.6": Allowed,
	"python3.7": Allowed,
	"python3.8": Allowed,
	"python3.9": Allowed,
	"python3.10": Allowed,
	"resize2fs": Allowed,

	"aarch64-linux-android-addr2line":    Allowed,
	"aarch64-linux-android-ar":           Allowed,
	"aarch64-linux-android-as":           Allowed,
	"aarch64-linux-android-c++filt":      Allowed,
	"aarch64-linux-android-dwp":          Allowed,
	"aarch64-linux-android-elfedit":      Allowed,
	"aarch64-linux-android-gcc":          Allowed,
	"aarch64-linux-android-gcc-ar":       Allowed,
	"aarch64-linux-android-gcc-nm":       Allowed,
	"aarch64-linux-android-gcc-ranlib":   Allowed,
	"aarch64-linux-android-gcov":         Allowed,
	"aarch64-linux-android-gcov-tool":    Allowed,
	"aarch64-linux-android-gprof":        Allowed,
	"aarch64-linux-android-ld":           Allowed,
	"aarch64-linux-android-ld.bfd":       Allowed,
	"aarch64-linux-android-ld.gold":      Allowed,
	"aarch64-linux-android-nm":           Allowed,
	"aarch64-linux-android-objcopy":      Allowed,
	"aarch64-linux-android-objdump":      Allowed,
	"aarch64-linux-android-ranlib":       Allowed,
	"aarch64-linux-android-readelf":      Allowed,
	"aarch64-linux-android-size":         Allowed,
	"aarch64-linux-android-strings":      Allowed,
	"aarch64-linux-android-strip":        Allowed,
	"aarch64-linux-gnu-as":               Allowed,
	"arm-linux-androideabi-addr2line":    Allowed,
	"arm-linux-androideabi-ar":           Allowed,
	"arm-linux-androideabi-as":           Allowed,
	"arm-linux-androideabi-c++filt":      Allowed,
	"arm-linux-androideabi-cpp":          Allowed,
	"arm-linux-androideabi-dwp":          Allowed,
	"arm-linux-androideabi-elfedit":      Allowed,
	"arm-linux-androideabi-gcc":          Allowed,
	"arm-linux-androideabi-gcc-ar":       Allowed,
	"arm-linux-androideabi-gcc-nm":       Allowed,
	"arm-linux-androideabi-gcc-ranlib":   Allowed,
	"arm-linux-androideabi-gcov":         Allowed,
	"arm-linux-androideabi-gcov-tool":    Allowed,
	"arm-linux-androideabi-gprof":        Allowed,
	"arm-linux-androideabi-ld":           Allowed,
	"arm-linux-androideabi-ld.bfd":       Allowed,
	"arm-linux-androideabi-ld.gold":      Allowed,
	"arm-linux-androideabi-nm":           Allowed,
	"arm-linux-androideabi-objcopy":      Allowed,
	"arm-linux-androideabi-objdump":      Allowed,
	"arm-linux-androideabi-ranlib":       Allowed,
	"arm-linux-androideabi-readelf":      Allowed,
	"arm-linux-androideabi-size":         Allowed,
	"arm-linux-androideabi-strings":      Allowed,
	"arm-linux-androideabi-strip":        Allowed,
	"arm-linux-androidkernel-addr2line":  Allowed,
	"arm-linux-androidkernel-ar":         Allowed,
	"arm-linux-androidkernel-as":         Allowed,
	"arm-linux-androidkernel-c++filt":    Allowed,
	"arm-linux-androidkernel-cpp":        Allowed,
	"arm-linux-androidkernel-dwp":        Allowed,
	"arm-linux-androidkernel-elfedit":    Allowed,
	"arm-linux-androidkernel-gcc":        Allowed,
	"arm-linux-androidkernel-gcc-ar":     Allowed,
	"arm-linux-androidkernel-gcc-nm":     Allowed,
	"arm-linux-androidkernel-gcc-ranlib": Allowed,
	"arm-linux-androidkernel-gcov":       Allowed,
	"arm-linux-androidkernel-gcov-tool":  Allowed,
	"arm-linux-androidkernel-gprof":      Allowed,
	"arm-linux-androidkernel-ld":         Allowed,
	"arm-linux-androidkernel-ld.bfd":     Allowed,
	"arm-linux-androidkernel-ld.gold":    Allowed,
	"arm-linux-androidkernel-nm":         Allowed,
	"arm-linux-androidkernel-objcopy":    Allowed,
	"arm-linux-androidkernel-objdump":    Allowed,
	"arm-linux-androidkernel-ranlib":     Allowed,
	"arm-linux-androidkernel-readelf":    Allowed,
	"arm-linux-androidkernel-size":       Allowed,
	"arm-linux-androidkernel-strings":    Allowed,
	"arm-linux-androidkernel-strip":      Allowed,

	// Host toolchain is removed. In-tree toolchain should be used instead.
	// GCC also can't find cc1 with this implementation.
	"ar":         Forbidden,
	"as":         Forbidden,
	"cc":         Forbidden,
	"clang":      Forbidden,
	"clang++":    Forbidden,
	"gcc":        Forbidden,
	"g++":        Forbidden,
	"ld":         Forbidden,
	"ld.bfd":     Forbidden,
	"ld.gold":    Forbidden,
	"pkg-config": Allowed,

	// These are toybox tools that only work on Linux.
	"pgrep": LinuxOnlyPrebuilt,
	"pkill": LinuxOnlyPrebuilt,
	"ps":    LinuxOnlyPrebuilt,
}

func init() {
	if runtime.GOOS == "darwin" {
		Configuration["sw_vers"] = Allowed
		Configuration["xcrun"] = Allowed

		// We don't have darwin prebuilts for some tools,
		// so allow the host versions.
		for name, config := range Configuration {
			if config.LinuxOnlyPrebuilt {
				Configuration[name] = Allowed
			}
		}
	}
}
