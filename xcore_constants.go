/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst /Users/ben/src/capstone/bindings/python/capstone/
	2014-11-21T09:34:57+12:00

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [xcore_const.py]
// Operand type for instruction's operands
const (
	XCORE_OP_INVALID = C.XCORE_OP_INVALID
	XCORE_OP_REG     = C.XCORE_OP_REG
	XCORE_OP_IMM     = C.XCORE_OP_IMM
	XCORE_OP_MEM     = C.XCORE_OP_MEM
)

// XCore registers
const (
	XCORE_REG_INVALID = C.XCORE_REG_INVALID
	XCORE_REG_CP      = C.XCORE_REG_CP
	XCORE_REG_DP      = C.XCORE_REG_DP
	XCORE_REG_LR      = C.XCORE_REG_LR
	XCORE_REG_SP      = C.XCORE_REG_SP
	XCORE_REG_R0      = C.XCORE_REG_R0
	XCORE_REG_R1      = C.XCORE_REG_R1
	XCORE_REG_R2      = C.XCORE_REG_R2
	XCORE_REG_R3      = C.XCORE_REG_R3
	XCORE_REG_R4      = C.XCORE_REG_R4
	XCORE_REG_R5      = C.XCORE_REG_R5
	XCORE_REG_R6      = C.XCORE_REG_R6
	XCORE_REG_R7      = C.XCORE_REG_R7
	XCORE_REG_R8      = C.XCORE_REG_R8
	XCORE_REG_R9      = C.XCORE_REG_R9
	XCORE_REG_R10     = C.XCORE_REG_R10
	XCORE_REG_R11     = C.XCORE_REG_R11
)

// pseudo registers
const (
	XCORE_REG_PC     = C.XCORE_REG_PC
	XCORE_REG_SCP    = C.XCORE_REG_SCP
	XCORE_REG_SSR    = C.XCORE_REG_SSR
	XCORE_REG_ET     = C.XCORE_REG_ET
	XCORE_REG_ED     = C.XCORE_REG_ED
	XCORE_REG_SED    = C.XCORE_REG_SED
	XCORE_REG_KEP    = C.XCORE_REG_KEP
	XCORE_REG_KSP    = C.XCORE_REG_KSP
	XCORE_REG_ID     = C.XCORE_REG_ID
	XCORE_REG_ENDING = C.XCORE_REG_ENDING
)

// XCore instruction
const (
	XCORE_INS_INVALID = C.XCORE_INS_INVALID
	XCORE_INS_ADD     = C.XCORE_INS_ADD
	XCORE_INS_ANDNOT  = C.XCORE_INS_ANDNOT
	XCORE_INS_AND     = C.XCORE_INS_AND
	XCORE_INS_ASHR    = C.XCORE_INS_ASHR
	XCORE_INS_BAU     = C.XCORE_INS_BAU
	XCORE_INS_BITREV  = C.XCORE_INS_BITREV
	XCORE_INS_BLA     = C.XCORE_INS_BLA
	XCORE_INS_BLAT    = C.XCORE_INS_BLAT
	XCORE_INS_BL      = C.XCORE_INS_BL
	XCORE_INS_BF      = C.XCORE_INS_BF
	XCORE_INS_BT      = C.XCORE_INS_BT
	XCORE_INS_BU      = C.XCORE_INS_BU
	XCORE_INS_BRU     = C.XCORE_INS_BRU
	XCORE_INS_BYTEREV = C.XCORE_INS_BYTEREV
	XCORE_INS_CHKCT   = C.XCORE_INS_CHKCT
	XCORE_INS_CLRE    = C.XCORE_INS_CLRE
	XCORE_INS_CLRPT   = C.XCORE_INS_CLRPT
	XCORE_INS_CLRSR   = C.XCORE_INS_CLRSR
	XCORE_INS_CLZ     = C.XCORE_INS_CLZ
	XCORE_INS_CRC8    = C.XCORE_INS_CRC8
	XCORE_INS_CRC32   = C.XCORE_INS_CRC32
	XCORE_INS_DCALL   = C.XCORE_INS_DCALL
	XCORE_INS_DENTSP  = C.XCORE_INS_DENTSP
	XCORE_INS_DGETREG = C.XCORE_INS_DGETREG
	XCORE_INS_DIVS    = C.XCORE_INS_DIVS
	XCORE_INS_DIVU    = C.XCORE_INS_DIVU
	XCORE_INS_DRESTSP = C.XCORE_INS_DRESTSP
	XCORE_INS_DRET    = C.XCORE_INS_DRET
	XCORE_INS_ECALLF  = C.XCORE_INS_ECALLF
	XCORE_INS_ECALLT  = C.XCORE_INS_ECALLT
	XCORE_INS_EDU     = C.XCORE_INS_EDU
	XCORE_INS_EEF     = C.XCORE_INS_EEF
	XCORE_INS_EET     = C.XCORE_INS_EET
	XCORE_INS_EEU     = C.XCORE_INS_EEU
	XCORE_INS_ENDIN   = C.XCORE_INS_ENDIN
	XCORE_INS_ENTSP   = C.XCORE_INS_ENTSP
	XCORE_INS_EQ      = C.XCORE_INS_EQ
	XCORE_INS_EXTDP   = C.XCORE_INS_EXTDP
	XCORE_INS_EXTSP   = C.XCORE_INS_EXTSP
	XCORE_INS_FREER   = C.XCORE_INS_FREER
	XCORE_INS_FREET   = C.XCORE_INS_FREET
	XCORE_INS_GETD    = C.XCORE_INS_GETD
	XCORE_INS_GET     = C.XCORE_INS_GET
	XCORE_INS_GETN    = C.XCORE_INS_GETN
	XCORE_INS_GETR    = C.XCORE_INS_GETR
	XCORE_INS_GETSR   = C.XCORE_INS_GETSR
	XCORE_INS_GETST   = C.XCORE_INS_GETST
	XCORE_INS_GETTS   = C.XCORE_INS_GETTS
	XCORE_INS_INCT    = C.XCORE_INS_INCT
	XCORE_INS_INIT    = C.XCORE_INS_INIT
	XCORE_INS_INPW    = C.XCORE_INS_INPW
	XCORE_INS_INSHR   = C.XCORE_INS_INSHR
	XCORE_INS_INT     = C.XCORE_INS_INT
	XCORE_INS_IN      = C.XCORE_INS_IN
	XCORE_INS_KCALL   = C.XCORE_INS_KCALL
	XCORE_INS_KENTSP  = C.XCORE_INS_KENTSP
	XCORE_INS_KRESTSP = C.XCORE_INS_KRESTSP
	XCORE_INS_KRET    = C.XCORE_INS_KRET
	XCORE_INS_LADD    = C.XCORE_INS_LADD
	XCORE_INS_LD16S   = C.XCORE_INS_LD16S
	XCORE_INS_LD8U    = C.XCORE_INS_LD8U
	XCORE_INS_LDA16   = C.XCORE_INS_LDA16
	XCORE_INS_LDAP    = C.XCORE_INS_LDAP
	XCORE_INS_LDAW    = C.XCORE_INS_LDAW
	XCORE_INS_LDC     = C.XCORE_INS_LDC
	XCORE_INS_LDW     = C.XCORE_INS_LDW
	XCORE_INS_LDIVU   = C.XCORE_INS_LDIVU
	XCORE_INS_LMUL    = C.XCORE_INS_LMUL
	XCORE_INS_LSS     = C.XCORE_INS_LSS
	XCORE_INS_LSUB    = C.XCORE_INS_LSUB
	XCORE_INS_LSU     = C.XCORE_INS_LSU
	XCORE_INS_MACCS   = C.XCORE_INS_MACCS
	XCORE_INS_MACCU   = C.XCORE_INS_MACCU
	XCORE_INS_MJOIN   = C.XCORE_INS_MJOIN
	XCORE_INS_MKMSK   = C.XCORE_INS_MKMSK
	XCORE_INS_MSYNC   = C.XCORE_INS_MSYNC
	XCORE_INS_MUL     = C.XCORE_INS_MUL
	XCORE_INS_NEG     = C.XCORE_INS_NEG
	XCORE_INS_NOT     = C.XCORE_INS_NOT
	XCORE_INS_OR      = C.XCORE_INS_OR
	XCORE_INS_OUTCT   = C.XCORE_INS_OUTCT
	XCORE_INS_OUTPW   = C.XCORE_INS_OUTPW
	XCORE_INS_OUTSHR  = C.XCORE_INS_OUTSHR
	XCORE_INS_OUTT    = C.XCORE_INS_OUTT
	XCORE_INS_OUT     = C.XCORE_INS_OUT
	XCORE_INS_PEEK    = C.XCORE_INS_PEEK
	XCORE_INS_REMS    = C.XCORE_INS_REMS
	XCORE_INS_REMU    = C.XCORE_INS_REMU
	XCORE_INS_RETSP   = C.XCORE_INS_RETSP
	XCORE_INS_SETCLK  = C.XCORE_INS_SETCLK
	XCORE_INS_SET     = C.XCORE_INS_SET
	XCORE_INS_SETC    = C.XCORE_INS_SETC
	XCORE_INS_SETD    = C.XCORE_INS_SETD
	XCORE_INS_SETEV   = C.XCORE_INS_SETEV
	XCORE_INS_SETN    = C.XCORE_INS_SETN
	XCORE_INS_SETPSC  = C.XCORE_INS_SETPSC
	XCORE_INS_SETPT   = C.XCORE_INS_SETPT
	XCORE_INS_SETRDY  = C.XCORE_INS_SETRDY
	XCORE_INS_SETSR   = C.XCORE_INS_SETSR
	XCORE_INS_SETTW   = C.XCORE_INS_SETTW
	XCORE_INS_SETV    = C.XCORE_INS_SETV
	XCORE_INS_SEXT    = C.XCORE_INS_SEXT
	XCORE_INS_SHL     = C.XCORE_INS_SHL
	XCORE_INS_SHR     = C.XCORE_INS_SHR
	XCORE_INS_SSYNC   = C.XCORE_INS_SSYNC
	XCORE_INS_ST16    = C.XCORE_INS_ST16
	XCORE_INS_ST8     = C.XCORE_INS_ST8
	XCORE_INS_STW     = C.XCORE_INS_STW
	XCORE_INS_SUB     = C.XCORE_INS_SUB
	XCORE_INS_SYNCR   = C.XCORE_INS_SYNCR
	XCORE_INS_TESTCT  = C.XCORE_INS_TESTCT
	XCORE_INS_TESTLCL = C.XCORE_INS_TESTLCL
	XCORE_INS_TESTWCT = C.XCORE_INS_TESTWCT
	XCORE_INS_TSETMR  = C.XCORE_INS_TSETMR
	XCORE_INS_START   = C.XCORE_INS_START
	XCORE_INS_WAITEF  = C.XCORE_INS_WAITEF
	XCORE_INS_WAITET  = C.XCORE_INS_WAITET
	XCORE_INS_WAITEU  = C.XCORE_INS_WAITEU
	XCORE_INS_XOR     = C.XCORE_INS_XOR
	XCORE_INS_ZEXT    = C.XCORE_INS_ZEXT
	XCORE_INS_ENDING  = C.XCORE_INS_ENDING
)

// Group of XCore instructions
const (
	XCORE_GRP_INVALID = C.XCORE_GRP_INVALID
)

// Generic groups
const (
	XCORE_GRP_JUMP   = C.XCORE_GRP_JUMP
	XCORE_GRP_ENDING = C.XCORE_GRP_ENDING
)
