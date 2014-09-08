// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package tcp

/*
#include <netinet/tcp.h>
*/
import "C"

const (
	sysTCPI_OPT_TIMESTAMPS = C.TCPI_OPT_TIMESTAMPS
	sysTCPI_OPT_SACK       = C.TCPI_OPT_SACK
	sysTCPI_OPT_WSCALE     = C.TCPI_OPT_WSCALE
	sysTCPI_OPT_ECN        = C.TCPI_OPT_ECN
	sysTCPI_OPT_TOE        = C.TCPI_OPT_TOE

	sysSizeofTCPInfo = C.sizeof_struct_tcp_info
)

type sysTCPInfo C.struct_tcp_info