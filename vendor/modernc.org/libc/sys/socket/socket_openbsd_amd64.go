// Code generated by 'ccgo sys/socket/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o sys/socket/socket_openbsd_amd64.go -pkgname socket', DO NOT EDIT.

package socket

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	AF_APPLETALK                   = 16
	AF_BLUETOOTH                   = 32
	AF_CCITT                       = 10
	AF_CHAOS                       = 5
	AF_CNT                         = 21
	AF_COIP                        = 20
	AF_DATAKIT                     = 9
	AF_DECnet                      = 12
	AF_DLI                         = 13
	AF_E164                        = 26
	AF_ECMA                        = 8
	AF_ENCAP                       = 28
	AF_HYLINK                      = 15
	AF_IMPLINK                     = 3
	AF_INET                        = 2
	AF_INET6                       = 24
	AF_IPX                         = 23
	AF_ISDN                        = 26
	AF_ISO                         = 7
	AF_KEY                         = 30
	AF_LAT                         = 14
	AF_LINK                        = 18
	AF_LOCAL                       = 1
	AF_MAX                         = 36
	AF_MPLS                        = 33
	AF_NATM                        = 27
	AF_NS                          = 6
	AF_OSI                         = 7
	AF_PUP                         = 4
	AF_ROUTE                       = 17
	AF_SIP                         = 29
	AF_SNA                         = 11
	AF_UNIX                        = 1
	AF_UNSPEC                      = 0
	BIG_ENDIAN                     = 4321
	BYTE_ORDER                     = 1234
	LITTLE_ENDIAN                  = 1234
	MSG_BCAST                      = 0x100
	MSG_CMSG_CLOEXEC               = 0x800
	MSG_CTRUNC                     = 0x20
	MSG_DONTROUTE                  = 0x4
	MSG_DONTWAIT                   = 0x80
	MSG_EOR                        = 0x8
	MSG_MCAST                      = 0x200
	MSG_NOSIGNAL                   = 0x400
	MSG_OOB                        = 0x1
	MSG_PEEK                       = 0x2
	MSG_TRUNC                      = 0x10
	MSG_WAITALL                    = 0x40
	NET_BPF_BUFSIZE                = 1
	NET_BPF_MAXBUFSIZE             = 2
	NET_BPF_MAXID                  = 3
	NET_KEY_MAXID                  = 3
	NET_KEY_SADB_DUMP              = 1
	NET_KEY_SPD_DUMP               = 2
	NET_LINK_IFRXQ                 = 1
	NET_LINK_IFRXQ_MAXID           = 3
	NET_LINK_IFRXQ_PRESSURE_DROP   = 2
	NET_LINK_IFRXQ_PRESSURE_RETURN = 1
	NET_LINK_MAXID                 = 2
	NET_MAXID                      = 36
	NET_PFLOW_MAXID                = 2
	NET_PFLOW_STATS                = 1
	NET_RT_DUMP                    = 1
	NET_RT_FLAGS                   = 2
	NET_RT_IFLIST                  = 3
	NET_RT_IFNAMES                 = 6
	NET_RT_MAXID                   = 8
	NET_RT_SOURCE                  = 7
	NET_RT_STATS                   = 4
	NET_RT_TABLE                   = 5
	NET_UNIX_DEFERRED              = 7
	NET_UNIX_INFLIGHT              = 6
	NET_UNIX_MAXID                 = 8
	NET_UNIX_PROTO_MAXID           = 3
	PDP_ENDIAN                     = 3412
	PF_APPLETALK                   = 16
	PF_BLUETOOTH                   = 32
	PF_BPF                         = 31
	PF_CCITT                       = 10
	PF_CHAOS                       = 5
	PF_CNT                         = 21
	PF_COIP                        = 20
	PF_DATAKIT                     = 9
	PF_DECnet                      = 12
	PF_DLI                         = 13
	PF_ECMA                        = 8
	PF_ENCAP                       = 28
	PF_HYLINK                      = 15
	PF_IMPLINK                     = 3
	PF_INET                        = 2
	PF_INET6                       = 24
	PF_IPX                         = 23
	PF_ISDN                        = 26
	PF_ISO                         = 7
	PF_KEY                         = 30
	PF_LAT                         = 14
	PF_LINK                        = 18
	PF_LOCAL                       = 1
	PF_MAX                         = 36
	PF_MPLS                        = 33
	PF_NATM                        = 27
	PF_NS                          = 6
	PF_OSI                         = 7
	PF_PFLOW                       = 34
	PF_PIP                         = 25
	PF_PIPEX                       = 35
	PF_PUP                         = 4
	PF_ROUTE                       = 17
	PF_RTIP                        = 22
	PF_SIP                         = 29
	PF_SNA                         = 11
	PF_UNIX                        = 1
	PF_UNSPEC                      = 0
	PF_XTP                         = 19
	RT_TABLEID_BITS                = 8
	RT_TABLEID_MASK                = 0xff
	RT_TABLEID_MAX                 = 255
	SCM_RIGHTS                     = 0x01
	SCM_TIMESTAMP                  = 0x04
	SHUT_RD                        = 0
	SHUT_RDWR                      = 2
	SHUT_WR                        = 1
	SOCK_CLOEXEC                   = 0x8000
	SOCK_DGRAM                     = 2
	SOCK_DNS                       = 0x1000
	SOCK_NONBLOCK                  = 0x4000
	SOCK_RAW                       = 3
	SOCK_RDM                       = 4
	SOCK_SEQPACKET                 = 5
	SOCK_STREAM                    = 1
	SOL_SOCKET                     = 0xffff
	SOMAXCONN                      = 128
	SO_ACCEPTCONN                  = 0x0002
	SO_BINDANY                     = 0x1000
	SO_BROADCAST                   = 0x0020
	SO_DEBUG                       = 0x0001
	SO_DOMAIN                      = 0x1024
	SO_DONTROUTE                   = 0x0010
	SO_ERROR                       = 0x1007
	SO_KEEPALIVE                   = 0x0008
	SO_LINGER                      = 0x0080
	SO_NETPROC                     = 0x1020
	SO_OOBINLINE                   = 0x0100
	SO_PEERCRED                    = 0x1022
	SO_PROTOCOL                    = 0x1025
	SO_RCVBUF                      = 0x1002
	SO_RCVLOWAT                    = 0x1004
	SO_RCVTIMEO                    = 0x1006
	SO_REUSEADDR                   = 0x0004
	SO_REUSEPORT                   = 0x0200
	SO_RTABLE                      = 0x1021
	SO_SNDBUF                      = 0x1001
	SO_SNDLOWAT                    = 0x1003
	SO_SNDTIMEO                    = 0x1005
	SO_SPLICE                      = 0x1023
	SO_TIMESTAMP                   = 0x0800
	SO_TYPE                        = 0x1008
	SO_USELOOPBACK                 = 0x0040
	SO_ZEROIZE                     = 0x2000
	UIO_MAXIOV                     = 1024
	UNPCTL_RECVSPACE               = 1
	UNPCTL_SENDSPACE               = 2
	X_BIG_ENDIAN                   = 4321
	X_BYTE_ORDER                   = 1234
	X_CLOCKID_T_DEFINED_           = 0
	X_CLOCK_T_DEFINED_             = 0
	X_FILE_OFFSET_BITS             = 64
	X_INT16_T_DEFINED_             = 0
	X_INT32_T_DEFINED_             = 0
	X_INT64_T_DEFINED_             = 0
	X_INT8_T_DEFINED_              = 0
	X_LITTLE_ENDIAN                = 1234
	X_LP64                         = 1
	X_MACHINE_CDEFS_H_             = 0
	X_MACHINE_ENDIAN_H_            = 0
	X_MACHINE__TYPES_H_            = 0
	X_MAX_PAGE_SHIFT               = 12
	X_OFF_T_DEFINED_               = 0
	X_PDP_ENDIAN                   = 3412
	X_PID_T_DEFINED_               = 0
	X_QUAD_HIGHWORD                = 1
	X_QUAD_LOWWORD                 = 0
	X_RET_PROTECTOR                = 1
	X_SA_FAMILY_T_DEFINED_         = 0
	X_SIZE_T_DEFINED_              = 0
	X_SOCKLEN_T_DEFINED_           = 0
	X_SSIZE_T_DEFINED_             = 0
	X_STACKALIGNBYTES              = 15
	X_SYS_CDEFS_H_                 = 0
	X_SYS_ENDIAN_H_                = 0
	X_SYS_SOCKET_H_                = 0
	X_SYS_TYPES_H_                 = 0
	X_SYS_UIO_H_                   = 0
	X_SYS__ENDIAN_H_               = 0
	X_SYS__TYPES_H_                = 0
	X_TIMER_T_DEFINED_             = 0
	X_TIMEVAL_DECLARED             = 0
	X_TIME_T_DEFINED_              = 0
	X_UINT16_T_DEFINED_            = 0
	X_UINT32_T_DEFINED_            = 0
	X_UINT64_T_DEFINED_            = 0
	X_UINT8_T_DEFINED_             = 0
	Pseudo_AF_HDRCMPLT             = 31
	Pseudo_AF_PFLOW                = 34
	Pseudo_AF_PIP                  = 25
	Pseudo_AF_PIPEX                = 35
	Pseudo_AF_RTIP                 = 22
	Pseudo_AF_XTP                  = 19
	Unix                           = 1
)

const ( /* uio.h:57:1: */
	UIO_READ  = 0
	UIO_WRITE = 1
)

// Segment flag values.
const ( /* uio.h:60:1: */
	UIO_USERSPACE = 0 // from user data space
	UIO_SYSSPACE  = 1
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

//	$OpenBSD: socket.h,v 1.100 2021/05/17 17:54:31 claudio Exp $
//	$NetBSD: socket.h,v 1.14 1996/02/09 18:25:36 christos Exp $

// Copyright (c) 1982, 1985, 1986, 1988, 1993, 1994
//	The Regents of the University of California.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)socket.h	8.4 (Berkeley) 2/21/94

// get the definitions for struct iovec, size_t, ssize_t, and <sys/cdefs.h>
//	$OpenBSD: uio.h,v 1.19 2018/08/20 16:00:22 mpi Exp $
//	$NetBSD: uio.h,v 1.12 1996/02/09 18:25:45 christos Exp $

// Copyright (c) 1982, 1986, 1993, 1994
//	The Regents of the University of California.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)uio.h	8.5 (Berkeley) 2/22/94

//	$OpenBSD: cdefs.h,v 1.43 2018/10/29 17:10:40 guenther Exp $
//	$NetBSD: cdefs.h,v 1.16 1996/04/03 20:46:39 christos Exp $

// Copyright (c) 1991, 1993
//	The Regents of the University of California.  All rights reserved.
//
// This code is derived from software contributed to Berkeley by
// Berkeley Software Design, Inc.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)cdefs.h	8.7 (Berkeley) 1/21/94

//	$OpenBSD: cdefs.h,v 1.3 2013/03/28 17:30:45 martynas Exp $

// Written by J.T. Conklin <jtc@wimsey.com> 01/17/95.
// Public domain.

// Macro to test if we're using a specific version of gcc or later.

// The __CONCAT macro is used to concatenate parts of symbol names, e.g.
// with "#define OLD(foo) __CONCAT(old,foo)", OLD(foo) produces oldfoo.
// The __CONCAT macro is a bit tricky -- make sure you don't put spaces
// in between its arguments.  Do not use __CONCAT on double-quoted strings,
// such as those from the __STRING macro: to concatenate strings just put
// them next to each other.

// GCC1 and some versions of GCC2 declare dead (non-returning) and
// pure (no side effects) functions using "volatile" and "const";
// unfortunately, these then cause warnings under "-ansi -pedantic".
// GCC >= 2.5 uses the __attribute__((attrs)) style.  All of these
// work for GNU C++ (modulo a slight glitch in the C++ grammar in
// the distribution version of 2.5.5).

// __returns_twice makes the compiler not assume the function
// only returns once.  This affects registerisation of variables:
// even local variables need to be in memory across such a call.
// Example: setjmp()

// __only_inline makes the compiler only use this function definition
// for inlining; references that can't be inlined will be left as
// external references instead of generating a local copy.  The
// matching library should include a simple extern definition for
// the function to handle those references.  c.f. ctype.h

// GNU C version 2.96 adds explicit branch prediction so that
// the CPU back-end can hint the processor and also so that
// code blocks can be reordered such that the predicted path
// sees a more linear flow, thus improving cache behavior, etc.
//
// The following two macros provide us with a way to utilize this
// compiler feature.  Use __predict_true() if you expect the expression
// to evaluate to true, and __predict_false() if you expect the
// expression to evaluate to false.
//
// A few notes about usage:
//
//	* Generally, __predict_false() error condition checks (unless
//	  you have some _strong_ reason to do otherwise, in which case
//	  document it), and/or __predict_true() `no-error' condition
//	  checks, assuming you want to optimize for the no-error case.
//
//	* Other than that, if you don't know the likelihood of a test
//	  succeeding from empirical or other `hard' evidence, don't
//	  make predictions.
//
//	* These are meant to be used in places that are run `a lot'.
//	  It is wasteful to make predictions in code that is run
//	  seldomly (e.g. at subsystem initialization time) as the
//	  basic block reordering that this affects can often generate
//	  larger code.

// Delete pseudo-keywords wherever they are not available or needed.

// The __packed macro indicates that a variable or structure members
// should have the smallest possible alignment, despite any host CPU
// alignment requirements.
//
// The __aligned(x) macro specifies the minimum alignment of a
// variable or structure.
//
// These macros together are useful for describing the layout and
// alignment of messages exchanged with hardware or other systems.

// "The nice thing about standards is that there are so many to choose from."
// There are a number of "feature test macros" specified by (different)
// standards that determine which interfaces and types the header files
// should expose.
//
// Because of inconsistencies in these macros, we define our own
// set in the private name space that end in _VISIBLE.  These are
// always defined and so headers can test their values easily.
// Things can get tricky when multiple feature macros are defined.
// We try to take the union of all the features requested.
//
// The following macros are guaranteed to have a value after cdefs.h
// has been included:
//	__POSIX_VISIBLE
//	__XPG_VISIBLE
//	__ISO_C_VISIBLE
//	__BSD_VISIBLE

// X/Open Portability Guides and Single Unix Specifications.
// _XOPEN_SOURCE				XPG3
// _XOPEN_SOURCE && _XOPEN_VERSION = 4		XPG4
// _XOPEN_SOURCE && _XOPEN_SOURCE_EXTENDED = 1	XPG4v2
// _XOPEN_SOURCE == 500				XPG5
// _XOPEN_SOURCE == 520				XPG5v2
// _XOPEN_SOURCE == 600				POSIX 1003.1-2001 with XSI
// _XOPEN_SOURCE == 700				POSIX 1003.1-2008 with XSI
//
// The XPG spec implies a specific value for _POSIX_C_SOURCE.

// POSIX macros, these checks must follow the XOPEN ones above.
//
// _POSIX_SOURCE == 1		1003.1-1988 (superseded by _POSIX_C_SOURCE)
// _POSIX_C_SOURCE == 1		1003.1-1990
// _POSIX_C_SOURCE == 2		1003.2-1992
// _POSIX_C_SOURCE == 199309L	1003.1b-1993
// _POSIX_C_SOURCE == 199506L   1003.1c-1995, 1003.1i-1995,
//				and the omnibus ISO/IEC 9945-1:1996
// _POSIX_C_SOURCE == 200112L   1003.1-2001
// _POSIX_C_SOURCE == 200809L   1003.1-2008
//
// The POSIX spec implies a specific value for __ISO_C_VISIBLE, though
// this may be overridden by the _ISOC99_SOURCE macro later.

// _ANSI_SOURCE means to expose ANSI C89 interfaces only.
// If the user defines it in addition to one of the POSIX or XOPEN
// macros, assume the POSIX/XOPEN macro(s) should take precedence.

// _ISOC99_SOURCE, _ISOC11_SOURCE, __STDC_VERSION__, and __cplusplus
// override any of the other macros since they are non-exclusive.

// Finally deal with BSD-specific interfaces that are not covered
// by any standards.  We expose these when none of the POSIX or XPG
// macros is defined or if the user explicitly asks for them.

// Default values.

//	$OpenBSD: _types.h,v 1.9 2014/08/22 23:05:15 krw Exp $

// -
// Copyright (c) 1990, 1993
//	The Regents of the University of California.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)types.h	8.3 (Berkeley) 1/5/94

//	$OpenBSD: _types.h,v 1.17 2018/03/05 01:15:25 deraadt Exp $

// -
// Copyright (c) 1990, 1993
//	The Regents of the University of California.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)types.h	8.3 (Berkeley) 1/5/94
//	@(#)ansi.h	8.2 (Berkeley) 1/4/94

// _ALIGN(p) rounds p (pointer or byte index) up to a correctly-aligned
// value for all data types (int, long, ...).   The result is an
// unsigned long and must be cast to any desired pointer type.
//
// _ALIGNED_POINTER is a boolean macro that checks whether an address
// is valid to fetch data elements of type t from on this architecture.
// This does not reflect the optimal alignment, just the possibility
// (within reasonable limits).

// 7.18.1.1 Exact-width integer types
type X__int8_t = int8     /* _types.h:61:22 */
type X__uint8_t = uint8   /* _types.h:62:24 */
type X__int16_t = int16   /* _types.h:63:17 */
type X__uint16_t = uint16 /* _types.h:64:25 */
type X__int32_t = int32   /* _types.h:65:15 */
type X__uint32_t = uint32 /* _types.h:66:23 */
type X__int64_t = int64   /* _types.h:67:20 */
type X__uint64_t = uint64 /* _types.h:68:28 */

// 7.18.1.2 Minimum-width integer types
type X__int_least8_t = X__int8_t     /* _types.h:71:19 */
type X__uint_least8_t = X__uint8_t   /* _types.h:72:20 */
type X__int_least16_t = X__int16_t   /* _types.h:73:20 */
type X__uint_least16_t = X__uint16_t /* _types.h:74:21 */
type X__int_least32_t = X__int32_t   /* _types.h:75:20 */
type X__uint_least32_t = X__uint32_t /* _types.h:76:21 */
type X__int_least64_t = X__int64_t   /* _types.h:77:20 */
type X__uint_least64_t = X__uint64_t /* _types.h:78:21 */

// 7.18.1.3 Fastest minimum-width integer types
type X__int_fast8_t = X__int32_t    /* _types.h:81:20 */
type X__uint_fast8_t = X__uint32_t  /* _types.h:82:21 */
type X__int_fast16_t = X__int32_t   /* _types.h:83:20 */
type X__uint_fast16_t = X__uint32_t /* _types.h:84:21 */
type X__int_fast32_t = X__int32_t   /* _types.h:85:20 */
type X__uint_fast32_t = X__uint32_t /* _types.h:86:21 */
type X__int_fast64_t = X__int64_t   /* _types.h:87:20 */
type X__uint_fast64_t = X__uint64_t /* _types.h:88:21 */

// 7.18.1.4 Integer types capable of holding object pointers
type X__intptr_t = int64   /* _types.h:103:16 */
type X__uintptr_t = uint64 /* _types.h:104:24 */

// 7.18.1.5 Greatest-width integer types
type X__intmax_t = X__int64_t   /* _types.h:107:20 */
type X__uintmax_t = X__uint64_t /* _types.h:108:21 */

// Register size
type X__register_t = int64 /* _types.h:111:16 */

// VM system types
type X__vaddr_t = uint64 /* _types.h:114:24 */
type X__paddr_t = uint64 /* _types.h:115:24 */
type X__vsize_t = uint64 /* _types.h:116:24 */
type X__psize_t = uint64 /* _types.h:117:24 */

// Standard system types
type X__double_t = float64           /* _types.h:120:18 */
type X__float_t = float32            /* _types.h:121:17 */
type X__ptrdiff_t = int64            /* _types.h:122:16 */
type X__size_t = uint64              /* _types.h:123:24 */
type X__ssize_t = int64              /* _types.h:124:16 */
type X__va_list = X__builtin_va_list /* _types.h:126:27 */

// Wide character support types
type X__wchar_t = int32     /* _types.h:133:15 */
type X__wint_t = int32      /* _types.h:135:15 */
type X__rune_t = int32      /* _types.h:136:15 */
type X__wctrans_t = uintptr /* _types.h:137:14 */
type X__wctype_t = uintptr  /* _types.h:138:14 */

type X__blkcnt_t = X__int64_t    /* _types.h:39:19 */ // blocks allocated for file
type X__blksize_t = X__int32_t   /* _types.h:40:19 */ // optimal blocksize for I/O
type X__clock_t = X__int64_t     /* _types.h:41:19 */ // ticks in CLOCKS_PER_SEC
type X__clockid_t = X__int32_t   /* _types.h:42:19 */ // CLOCK_* identifiers
type X__cpuid_t = uint64         /* _types.h:43:23 */ // CPU id
type X__dev_t = X__int32_t       /* _types.h:44:19 */ // device number
type X__fixpt_t = X__uint32_t    /* _types.h:45:20 */ // fixed point number
type X__fsblkcnt_t = X__uint64_t /* _types.h:46:20 */ // file system block count
type X__fsfilcnt_t = X__uint64_t /* _types.h:47:20 */ // file system file count
type X__gid_t = X__uint32_t      /* _types.h:48:20 */ // group id
type X__id_t = X__uint32_t       /* _types.h:49:20 */ // may contain pid, uid or gid
type X__in_addr_t = X__uint32_t  /* _types.h:50:20 */ // base type for internet address
type X__in_port_t = X__uint16_t  /* _types.h:51:20 */ // IP port type
type X__ino_t = X__uint64_t      /* _types.h:52:20 */ // inode number
type X__key_t = int64            /* _types.h:53:15 */ // IPC key (for Sys V IPC)
type X__mode_t = X__uint32_t     /* _types.h:54:20 */ // permissions
type X__nlink_t = X__uint32_t    /* _types.h:55:20 */ // link count
type X__off_t = X__int64_t       /* _types.h:56:19 */ // file offset or size
type X__pid_t = X__int32_t       /* _types.h:57:19 */ // process id
type X__rlim_t = X__uint64_t     /* _types.h:58:20 */ // resource limit
type X__sa_family_t = X__uint8_t /* _types.h:59:19 */ // sockaddr address family type
type X__segsz_t = X__int32_t     /* _types.h:60:19 */ // segment size
type X__socklen_t = X__uint32_t  /* _types.h:61:20 */ // length type for network syscalls
type X__suseconds_t = int64      /* _types.h:62:15 */ // microseconds (signed)
type X__swblk_t = X__int32_t     /* _types.h:63:19 */ // swap offset
type X__time_t = X__int64_t      /* _types.h:64:19 */ // epoch time
type X__timer_t = X__int32_t     /* _types.h:65:19 */ // POSIX timer identifiers
type X__uid_t = X__uint32_t      /* _types.h:66:20 */ // user id
type X__useconds_t = X__uint32_t /* _types.h:67:20 */ // microseconds

// mbstate_t is an opaque object to keep conversion state, during multibyte
// stream conversions. The content must not be referenced by user programs.
type X__mbstate_t = struct {
	F__ccgo_pad1 [0]uint64
	F__mbstate8  [128]int8
} /* _types.h:76:3 */

type Ssize_t = X__ssize_t /* uio.h:48:19 */

type Iovec = struct {
	Fiov_base uintptr
	Fiov_len  Size_t
} /* uio.h:51:1 */

// Tell sys/endian.h we have MD variants of the swap macros.

// Note that these macros evaluate their arguments several times.

// Public names

// These are specified to be function-like macros to match the spec

// POSIX names

// original BSD names

// these were exposed here before

// ancient stuff

type U_char = uint8   /* types.h:51:23 */
type U_short = uint16 /* types.h:52:24 */
type U_int = uint32   /* types.h:53:22 */
type U_long = uint64  /* types.h:54:23 */

type Unchar = uint8  /* types.h:56:23 */ // Sys V compatibility
type Ushort = uint16 /* types.h:57:24 */ // Sys V compatibility
type Uint = uint32   /* types.h:58:22 */ // Sys V compatibility
type Ulong = uint64  /* types.h:59:23 */ // Sys V compatibility

type Cpuid_t = X__cpuid_t       /* types.h:61:19 */ // CPU id
type Register_t = X__register_t /* types.h:62:22 */ // register-sized type

// XXX The exact-width bit types should only be exposed if __BSD_VISIBLE
//     but the rest of the includes are not ready for that yet.

type Int8_t = X__int8_t /* types.h:75:19 */

type Uint8_t = X__uint8_t /* types.h:80:20 */

type Int16_t = X__int16_t /* types.h:85:20 */

type Uint16_t = X__uint16_t /* types.h:90:21 */

type Int32_t = X__int32_t /* types.h:95:20 */

type Uint32_t = X__uint32_t /* types.h:100:21 */

type Int64_t = X__int64_t /* types.h:105:20 */

type Uint64_t = X__uint64_t /* types.h:110:21 */

// BSD-style unsigned bits types
type U_int8_t = X__uint8_t   /* types.h:114:19 */
type U_int16_t = X__uint16_t /* types.h:115:20 */
type U_int32_t = X__uint32_t /* types.h:116:20 */
type U_int64_t = X__uint64_t /* types.h:117:20 */

// quads, deprecated in favor of 64 bit int types
type Quad_t = X__int64_t    /* types.h:120:19 */
type U_quad_t = X__uint64_t /* types.h:121:20 */

// VM system types
type Vaddr_t = X__vaddr_t /* types.h:125:19 */
type Paddr_t = X__paddr_t /* types.h:126:19 */
type Vsize_t = X__vsize_t /* types.h:127:19 */
type Psize_t = X__psize_t /* types.h:128:19 */

// Standard system types
type Blkcnt_t = X__blkcnt_t       /* types.h:132:20 */ // blocks allocated for file
type Blksize_t = X__blksize_t     /* types.h:133:21 */ // optimal blocksize for I/O
type Caddr_t = uintptr            /* types.h:134:14 */ // core address
type Daddr32_t = X__int32_t       /* types.h:135:19 */ // 32-bit disk address
type Daddr_t = X__int64_t         /* types.h:136:19 */ // 64-bit disk address
type Dev_t = X__dev_t             /* types.h:137:18 */ // device number
type Fixpt_t = X__fixpt_t         /* types.h:138:19 */ // fixed point number
type Gid_t = X__gid_t             /* types.h:139:18 */ // group id
type Id_t = X__id_t               /* types.h:140:17 */ // may contain pid, uid or gid
type Ino_t = X__ino_t             /* types.h:141:18 */ // inode number
type Key_t = X__key_t             /* types.h:142:18 */ // IPC key (for Sys V IPC)
type Mode_t = X__mode_t           /* types.h:143:18 */ // permissions
type Nlink_t = X__nlink_t         /* types.h:144:19 */ // link count
type Rlim_t = X__rlim_t           /* types.h:145:18 */ // resource limit
type Segsz_t = X__segsz_t         /* types.h:146:19 */ // segment size
type Swblk_t = X__swblk_t         /* types.h:147:19 */ // swap offset
type Uid_t = X__uid_t             /* types.h:148:18 */ // user id
type Useconds_t = X__useconds_t   /* types.h:149:22 */ // microseconds
type Suseconds_t = X__suseconds_t /* types.h:150:23 */ // microseconds (signed)
type Fsblkcnt_t = X__fsblkcnt_t   /* types.h:151:22 */ // file system block count
type Fsfilcnt_t = X__fsfilcnt_t   /* types.h:152:22 */ // file system file count

// The following types may be defined in multiple header files.
type Clock_t = X__clock_t /* types.h:159:19 */

type Clockid_t = X__clockid_t /* types.h:164:21 */

type Pid_t = X__pid_t /* types.h:169:18 */

type Time_t = X__time_t /* types.h:184:18 */

type Timer_t = X__timer_t /* types.h:189:19 */

type Off_t = X__off_t /* types.h:194:18 */

// Major, minor numbers, dev_t's.

type Socklen_t = X__socklen_t /* socket.h:47:21 */ // length type for network syscalls

type Sa_family_t = X__sa_family_t /* socket.h:52:23 */ // sockaddr address family type

// Definitions related to sockets: types, address families, options.

// Types

// Socket creation flags

// Option flags per-socket.

// Additional options, not kept in so_options.

// Structure used for manipulating linger option.
type Linger = struct {
	Fl_onoff  int32
	Fl_linger int32
} /* socket.h:122:1 */

type Timeval = struct {
	Ftv_sec  Time_t
	Ftv_usec Suseconds_t
} /* socket.h:131:1 */

// Structure used for manipulating splice option.
type Splice = struct {
	Fsp_fd       int32
	F__ccgo_pad1 [4]byte
	Fsp_max      Off_t
	Fsp_idle     struct {
		Ftv_sec  Time_t
		Ftv_usec Suseconds_t
	}
} /* socket.h:140:1 */

// Maximum number of alternate routing tables

// Level number for (get/set)sockopt() to apply to socket itself.

// Address families.

// Structure used by kernel to store most
// addresses.
type Sockaddr = struct {
	Fsa_len    X__uint8_t
	Fsa_family Sa_family_t
	Fsa_data   [14]int8
} /* socket.h:209:1 */

// Sockaddr type which can hold any sockaddr type available
// in the system.
//
// Note: __ss_{len,family} is defined in RFC2553.  During RFC2553 discussion
// the field name went back and forth between ss_len and __ss_len,
// and RFC2553 specifies it to be __ss_len.  openbsd picked ss_len.
// For maximum portability, userland programmer would need to
// (1) make the code never touch ss_len portion (cast it into sockaddr and
// touch sa_len), or (2) add "-Dss_len=__ss_len" into CFLAGS to unify all
// occurrences (including header file) to __ss_len.
type Sockaddr_storage = struct {
	Fss_len    X__uint8_t
	Fss_family Sa_family_t
	F__ss_pad1 [6]uint8
	F__ss_pad2 X__uint64_t
	F__ss_pad3 [240]uint8
} /* socket.h:227:1 */

// Protocol families, same as address families for now.

// These are the valid values for the "how" field used by shutdown(2).

// Read using getsockopt() with SOL_SOCKET, SO_PEERCRED
type Sockpeercred = struct {
	Fuid Uid_t
	Fgid Gid_t
	Fpid Pid_t
} /* socket.h:300:1 */

// Definitions for network related sysctl, CTL_NET.
//
// Second level is protocol family.
// Third level is protocol number.
//
// Further levels are defined by the individual families below.

// PF_ROUTE - Routing table
//
// Four additional levels are defined:
//	Fourth: address family, 0 is wildcard
//	Fifth: type of info, defined below
//	Sixth: flag(s) to mask with for NET_RT_FLAGS
//	Seventh: routing table to use (facultative, defaults to 0)
//		 NET_RT_TABLE has the table id as sixth element.

// PF_UNIX - unix socket tunables

// PF_LINK - link layer or device tunables

// PF_KEY - Key Management

// PF_BPF  not really a family, but connected under CTL_NET

// PF_PFLOW not really a family, but connected under CTL_NET

// Maximum queue length specifiable by listen(2).

// Message header for recvmsg and sendmsg calls.
// Used value-result for recvmsg, value only for sendmsg.
type Msghdr = struct {
	Fmsg_name       uintptr
	Fmsg_namelen    Socklen_t
	F__ccgo_pad1    [4]byte
	Fmsg_iov        uintptr
	Fmsg_iovlen     uint32
	F__ccgo_pad2    [4]byte
	Fmsg_control    uintptr
	Fmsg_controllen Socklen_t
	Fmsg_flags      int32
} /* socket.h:483:1 */

// Header for ancillary data objects in msg_control buffer.
// Used for additional information with/about a datagram
// not expressible by flags.  The format is a sequence
// of message elements headed by cmsghdr structures.
type Cmsghdr = struct {
	Fcmsg_len   Socklen_t
	Fcmsg_level int32
	Fcmsg_type  int32
} /* socket.h:512:1 */

var _ int8 /* gen.c:2:13: */
