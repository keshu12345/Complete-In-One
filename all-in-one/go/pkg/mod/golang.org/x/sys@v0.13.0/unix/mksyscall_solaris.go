// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

/*
 This program reads a file containing function prototypes
 (like syscall_solaris.go) and generates system call bodies.
 The prototypes are marked by lines beginning with "//sys"
 and read like func declarations if //sys is replaced by func, but:
	* The parameter lists must give a name for each argument.
	  This includes return parameters.
	* The parameter lists must give a type for each argument:
	  the (x, y, z int) shorthand is not allowed.
	* If the return parameter is an error number, it must be named err.
	* If go func name needs to be different than its libc name,
	* or the function is not in libc, name could be specified
	* at the end, after "=" sign, like
	  //sys	getsockopt(s int, level int, name int, val uintptr, vallen *_Socklen) (err error) = libsocket.getsockopt
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	b32     = flag.Bool("b32", false, "32bit big-endian")
	l32     = flag.Bool("l32", false, "32bit little-endian")
	tags    = flag.String("tags", "", "build tags")
	illumos = flag.Bool("illumos", false, "illumos specific code generation")
)

// cmdLine returns this programs's commandline arguments
func cmdLine() string {
	return "go run mksyscall_solaris.go " + strings.Join(os.Args[1:], " ")
}

// goBuildTags returns build tags in the go:build format.
func goBuildTags() string {
	return strings.ReplaceAll(*tags, ",", " && ")
}

// plusBuildTags returns build tags in the +build format.
func plusBuildTags() string {
	return *tags
}

// Param is function parameter
type Param struct {
	Name string
	Type string
}

// usage prints the program usage
func usage() {
	fmt.Fprintf(os.Stderr, "usage: go run mksyscall_solaris.go [-b32 | -l32] [-tags x,y] [file ...]\n")
	os.Exit(1)
}

// parseParamList parses parameter list and returns a slice of parameters
func parseParamList(list string) []string {
	list = strings.TrimSpace(list)
	if list == "" {
		return []string{}
	}
	return regexp.MustCompile(`\s*,\s*`).Split(list, -1)
}

// parseParam splits a parameter into name and type
func parseParam(p string) Param {
	ps := regexp.MustCompile(`^(\S*) (\S*)$`).FindStringSubmatch(p)
	if ps == nil {
		fmt.Fprintf(os.Stderr, "malformed parameter: %s\n", p)
		os.Exit(1)
	}
	return Param{ps[1], ps[2]}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) <= 0 {
		fmt.Fprintf(os.Stderr, "no files to parse provided\n")
		usage()
	}

	endianness := ""
	if *b32 {
		endianness = "big-endian"
	} else if *l32 {
		endianness = "little-endian"
	}

	pack := ""
	text := ""
	dynimports := ""
	linknames := ""
	var vars []string
	for _, path := range flag.Args() {
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		s := bufio.NewScanner(file)
		for s.Scan() {
			t := s.Text()
			if p := regexp.MustCompile(`^package (\S+)$`).FindStringSubmatch(t); p != nil && pack == "" {
				pack = p[1]
			}
			nonblock := regexp.MustCompile(`^\/\/sysnb\t`).FindStringSubmatch(t)
			if regexp.MustCompile(`^\/\/sys\t`).FindStringSubmatch(t) == nil && nonblock == nil {
				continue
			}

			// Line must be of the form
			//	func Open(path string, mode int, perm int) (fd int, err error)
			// Split into name, in params, out params.
			f := regexp.MustCompile(`^\/\/sys(nb)?\t(\w+)\(([^()]*)\)\s*(?:\(([^()]+)\))?\s*(?:=\s*(?:(\w*)\.)?(\w*))?$`).FindStringSubmatch(t)
			if f == nil {
				fmt.Fprintf(os.Stderr, "%s:%s\nmalformed //sys declaration\n", path, t)
				os.Exit(1)
			}
			funct, inps, outps, modname, sysname := f[2], f[3], f[4], f[5], f[6]

			// Split argument lists on comma.
			in := parseParamList(inps)
			out := parseParamList(outps)

			inps = strings.Join(in, ", ")
			outps = strings.Join(out, ", ")

			// Try in vain to keep people from editing this file.
			// The theory is that they jump into the middle of the file
			// without reading the header.
			text += "// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT\n\n"

			// So file name.
			if modname == "" {
				modname = "libc"
			}

			// System call name.
			if sysname == "" {
				sysname = funct
			}

			// System call pointer variable name.
			sysvarname := fmt.Sprintf("proc%s", sysname)

			strconvfunc := "BytePtrFromString"
			strconvtype := "*byte"

			sysname = strings.ToLower(sysname) // All libc functions are lowercase.

			if funct != "ioctlPtrRet" {
				// Runtime import of function to allow cross-platform builds.
				dynimports += fmt.Sprintf("//go:cgo_import_dynamic libc_%s %s \"%s.so\"\n", sysname, sysname, modname)
				// Link symbol to proc address variable.
				linknames += fmt.Sprintf("//go:linkname %s libc_%s\n", sysvarname, sysname)
				// Library proc address variable.
				vars = append(vars, sysvarname)
			}

			// Go function header.
			outlist := strings.Join(out, ", ")
			if outlist != "" {
				outlist = fmt.Sprintf(" (%s)", outlist)
			}
			if text != "" {
				text += "\n"
			}
			text += fmt.Sprintf("func %s(%s)%s {\n", funct, strings.Join(in, ", "), outlist)

			// Check if err return available
			errvar := ""
			for _, param := range out {
				p := parseParam(param)
				if p.Type == "error" {
					errvar = p.Name
					continue
				}
			}

			// Prepare arguments to Syscall.
			var args []string
			n := 0
			for _, param := range in {
				p := parseParam(param)
				if regexp.MustCompile(`^\*`).FindStringSubmatch(p.Type) != nil {
					args = append(args, "uintptr(unsafe.Pointer("+p.Name+"))")
				} else if p.Type == "string" && errvar != "" {
					text += fmt.Sprintf("\tvar _p%d %s\n", n, strconvtype)
					text += fmt.Sprintf("\t_p%d, %s = %s(%s)\n", n, errvar, strconvfunc, p.Name)
					text += fmt.Sprintf("\tif %s != nil {\n\t\treturn\n\t}\n", errvar)
					args = append(args, fmt.Sprintf("uintptr(unsafe.Pointer(_p%d))", n))
					n++
				} else if p.Type == "string" {
					fmt.Fprintf(os.Stderr, path+":"+funct+" uses string arguments, but has no error return\n")
					text += fmt.Sprintf("\tvar _p%d %s\n", n, strconvtype)
					text += fmt.Sprintf("\t_p%d, _ = %s(%s)\n", n, strconvfunc, p.Name)
					args = append(args, fmt.Sprintf("uintptr(unsafe.Pointer(_p%d))", n))
					n++
				} else if s := regexp.MustCompile(`^\[\](.*)`).FindStringSubmatch(p.Type); s != nil {
					// Convert slice into pointer, length.
					// Have to be careful not to take address of &a[0] if len == 0:
					// pass nil in that case.
					text += fmt.Sprintf("\tvar _p%d *%s\n", n, s[1])
					text += fmt.Sprintf("\tif len(%s) > 0 {\n\t\t_p%d = &%s[0]\n\t}\n", p.Name, n, p.Name)
					args = append(args, fmt.Sprintf("uintptr(unsafe.Pointer(_p%d))", n), fmt.Sprintf("uintptr(len(%s))", p.Name))
					n++
				} else if p.Type == "int64" && endianness != "" {
					if endianness == "big-endian" {
						args = append(args, fmt.Sprintf("uintptr(%s>>32)", p.Name), fmt.Sprintf("uintptr(%s)", p.Name))
					} else {
						args = append(args, fmt.Sprintf("uintptr(%s)", p.Name), fmt.Sprintf("uintptr(%s>>32)", p.Name))
					}
				} else if p.Type == "bool" {
					text += fmt.Sprintf("\tvar _p%d uint32\n", n)
					text += fmt.Sprintf("\tif %s {\n\t\t_p%d = 1\n\t} else {\n\t\t_p%d = 0\n\t}\n", p.Name, n, n)
					args = append(args, fmt.Sprintf("uintptr(_p%d)", n))
					n++
				} else {
					args = append(args, fmt.Sprintf("uintptr(%s)", p.Name))
				}
			}
			nargs := len(args)

			// Determine which form to use; pad args with zeros.
			asm := "sysvicall6"
			if nonblock != nil {
				asm = "rawSysvicall6"
			}
			if len(args) <= 6 {
				for len(args) < 6 {
					args = append(args, "0")
				}
			} else {
				fmt.Fprintf(os.Stderr, "%s: too many arguments to system call\n", path)
				os.Exit(1)
			}

			// Actual call.
			arglist := strings.Join(args, ", ")
			call := fmt.Sprintf("%s(uintptr(unsafe.Pointer(&%s)), %d, %s)", asm, sysvarname, nargs, arglist)

			// Assign return values.
			body := ""
			ret := []string{"_", "_", "_"}
			doErrno := false
			for i := 0; i < len(out); i++ {
				p := parseParam(out[i])
				reg := ""
				if p.Name == "err" {
					reg = "e1"
					ret[2] = reg
					doErrno = true
				} else {
					reg = fmt.Sprintf("r%d", i)
					ret[i] = reg
				}
				if p.Type == "bool" {
					reg = fmt.Sprintf("%d != 0", reg)
				}
				if p.Type == "int64" && endianness != "" {
					// 64-bit number in r1:r0 or r0:r1.
					if i+2 > len(out) {
						fmt.Fprintf(os.Stderr, "%s: not enough registers for int64 return\n", path)
						os.Exit(1)
					}
					if endianness == "big-endian" {
						reg = fmt.Sprintf("int64(r%d)<<32 | int64(r%d)", i, i+1)
					} else {
						reg = fmt.Sprintf("int64(r%d)<<32 | int64(r%d)", i+1, i)
					}
					ret[i] = fmt.Sprintf("r%d", i)
					ret[i+1] = fmt.Sprintf("r%d", i+1)
				}
				if reg != "e1" {
					body += fmt.Sprintf("\t%s = %s(%s)\n", p.Name, p.Type, reg)
				}
			}
			if ret[0] == "_" && ret[1] == "_" && ret[2] == "_" {
				text += fmt.Sprintf("\t%s\n", call)
			} else {
				text += fmt.Sprintf("\t%s, %s, %s := %s\n", ret[0], ret[1], ret[2], call)
			}
			text += body

			if doErrno {
				text += "\tif e1 != 0 {\n"
				text += "\t\terr = errnoErr(e1)\n"
				text += "\t}\n"
			}
			text += "\treturn\n"
			text += "}\n"
		}
		if err := s.Err(); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		file.Close()
	}
	imp := ""
	if pack != "unix" {
		imp = "import \"golang.org/x/sys/unix\"\n"
	}

	syscallimp := ""
	if !*illumos {
		syscallimp = "\"syscall\""
	}

	vardecls := "\t" + strings.Join(vars, ",\n\t")
	vardecls += " syscallFunc"
	fmt.Printf(srcTemplate, cmdLine(), goBuildTags(), plusBuildTags(), pack, syscallimp, imp, dynimports, linknames, vardecls, text)
}

const srcTemplate = `// %s
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build %s
// +build %s

package %s

import (
        "unsafe"
        %s
)
%s
%s
%s
var (
%s	
)

%s
`
