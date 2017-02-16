package main

import "fmt"
import "os"
import "bufio"

var AC, PC, ZERO, NEGATIVE uint8
var MEM [255]uint8
var execops, run int

const NOP = 0   //00000000
const STA = 16  //00010000
const LDA = 32  //00100000
const ADD = 48  //00110000
const OR = 64   //01000000
const AND = 80  //01010000
const NOT = 96  //01100000
const JMP = 128 //10000000
const JN = 144  //10010000
const JZ = 160  //10100000
const HLT = 240 //11110000

func wait() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func nop() {
	fmt.Println("nop")
}
func sta(end uint8) {
	MEM[end] = AC
	fmt.Println(MEM[PC])
	PC += 1
}

func lda(end uint8) {
	AC = MEM[end]
	PC += 1
	estado()
}

func add(end uint8) {
	AC = MEM[end] + AC
	PC += 1
	estado()
}

func or(end uint8) {
	AC = MEM[end] | AC
	PC += 1
	estado()
}

func and(end uint8) {
	AC = MEM[end] & AC
	PC += 1
	estado()
}

func not() {
	AC = ^AC
	PC += 1
	estado()
}

func jmp(end uint8) {
	PC = end
}

func jn(end uint8) {
	if NEGATIVE == 1 {
		PC = end
	} else {
		PC += 1
	}
}

func jz(end uint8) {
	fmt.Print(" JZ! ")
	fmt.Println(ZERO)
	fmt.Println(PC)
	if ZERO == 1 {
		PC = end
	} else {
		PC += 1
	}
}

func hlt() {
	fmt.Println("HALT and Catch Fire!!")
	run = 0
}

func estado() {
	if AC == 0 {
		ZERO = 1
	} else {
		ZERO = 0
	}

	if AC > 127 {
		NEGATIVE = 1
	} else {
		NEGATIVE = 0
	}
	//registradores()
}

func registradores() {
	fmt.Print("AC: ")
	fmt.Println(AC)
	fmt.Print("PC: ")
	fmt.Println(PC)
	fmt.Print("NEGATIVE: ")
	fmt.Println(NEGATIVE)
	fmt.Print("ZERO: ")
	fmt.Println(ZERO)
	fmt.Println(MEM)
}

func decoder() {
	wait()
	execops += 1
	/* ** PRECISA SER CORRIGIDO, O NEANDER ORIGINAL VOLTA AO ENDEREÇO 00** */
	if PC == 255 {
		hlt()
		return
	}
	switch MEM[PC] {
	case NOP:
		fmt.Println("   nop")
		PC += 1
		nop()
	case STA:
		fmt.Println("   sta")
		PC += 1
		sta(MEM[PC])
	case LDA:
		fmt.Println("   lda")
		PC += 1
		lda(MEM[PC])
	case ADD:
		fmt.Println("   add")
		PC += 1
		add(MEM[PC])
	case OR:
		fmt.Println("   or")
		PC += 1
		or(MEM[PC])
	case AND:
		fmt.Println("   and")
		PC += 1
		and(MEM[PC])
	case NOT:
		fmt.Println("   not")
		not()
	case JMP:
		fmt.Println("   jmp")
		PC += 1
		jmp(MEM[PC])
	case JN:
		fmt.Println("   jn")
		PC += 1
		jn(MEM[PC])
	case JZ:
		fmt.Println("   jz")
		PC += 1
		jz(MEM[PC])
	case HLT:
		fmt.Println("   hlt")
		//fmt.Println(MEM[PC])
		hlt()
	default:
		fmt.Print("default")
		fmt.Println(MEM[PC])
		nop()
	}
	//PC +=1
	fmt.Print("Decoded, PC: ")
	fmt.Println(PC)
}

func programa() {
	MEM[0] = NOP
	MEM[1] = 128
	MEM[2] = JZ
	MEM[3] = 140
	MEM[4] = NOT
	MEM[5] = ADD
	MEM[6] = 133
	MEM[7] = ADD
	MEM[8] = 133
	MEM[9] = JZ
	MEM[10] = 40
	MEM[11] = 0
	MEM[12] = 0
	MEM[13] = 0
	MEM[14] = 0
	MEM[15] = 0
	MEM[128] = 0
	MEM[129] = 0
	MEM[130] = 0
	MEM[131] = 0
	MEM[132] = 0
	MEM[133] = 1
	MEM[134] = 255
	MEM[135] = 0

	MEM[140] = LDA
	MEM[141] = 129
	MEM[142] = JN
	MEM[143] = 154 //UmMaior
	MEM[144] = LDA
	MEM[145] = 130
	MEM[146] = JN
	MEM[147] = 158 //SoUmMaior
	MEM[148] = ADD
	MEM[149] = 129
	MEM[150] = STA
	MEM[151] = 131
	MEM[152] = JMP
	//MEM[153] = ??? //NotOver
	MEM[154] = LDA //UmMaior
	MEM[155] = 130
	MEM[156] = JN
	//MEM[157] = ??? //DoisMaior
	MEM[158] = ADD //SoUmMaior
	MEM[159] = 129
	MEM[160] = STA
	MEM[161] = 131
	MEM[162] = JN
	//MEM[163] = ??? //NotOver
	MEM[164] = JMP
	//MEM[165] = ??? //Over
	/*MEM[166] = 
	MEM[167] = 
	MEM[168] = 
	MEM[169] = 
	*/
	/*
	//OVER
	LDA
	133
	STA
	132
	HLT
	//NOTOVER
	LDA
	134
	STA
	132
	HLT*/
}

func main() {

	run = 1
	estado()
	//registradores()

	//AC = 10
	programa()
	registradores()

	for {
		decoder()
		//fmt.Println(PC)
		registradores()
		if run == 0 {
			break
		}
	}

	//registradores()
	fmt.Print("Total de operações executadas: ")
	fmt.Println(execops)
}

/*func programa(){
	MEM[0] = LDA
    MEM[1] = 128
	MEM[2] = JZ
    MEM[3] = 10
    MEM[4] = AND
    MEM[5] = 134
    MEM[6] = JZ
    MEM[7] = 48
    MEM[8] = JMP
    MEM[9] = 90
    MEM[10] = LDA
    MEM[11] = 129
    MEM[12] = JN
    MEM[13] = 27
    MEM[14] = LDA
    MEM[15] = 130
    MEM[16] = JN
    MEM[17] = 31
    MEM[18] = ADD
    MEM[19] = 129
    MEM[20] = STA
    MEM[21] = 131
    MEM[22] = LDA
    MEM[23] = 135
    MEM[24] = STA
    MEM[25] = 132
    MEM[26] = HLT
    MEM[27] = LDA
    MEM[28] = 130
    MEM[29] = JN
    MEM[30] = 42
	MEM[31] = ADD
	MEM[32] = 129
	MEM[33] = STA
	MEM[34] = 131
	MEM[35] = JN
	MEM[36] = 22
	MEM[37] = LDA
	MEM[38] = 134
	MEM[39] = STA
	MEM[40] = 132
	MEM[41] = HLT
	MEM[42] = ADD
	MEM[43] = 129
	MEM[44] = STA
	MEM[45] = 131
	MEM[46] = JMP
	MEM[47] = 37
	MEM[48] = LDA
	MEM[49] = 129
	MEM[50] = JN
	MEM[51] = 67
	MEM[52] = LDA
	MEM[53] = 130
	MEM[54] = JN
	MEM[55] = 71
	MEM[56] = ADD
	MEM[57] = 129
	MEM[58] = STA
	MEM[59] = 131
	MEM[60] = JN
	MEM[61] = 77
	MEM[62] = LDA
	MEM[63] = 135
	MEM[64] = STA
	MEM[65] = 132
	MEM[66] = HLT
	MEM[67] = LDA
	MEM[68] = 130
	MEM[69] = JN
	MEM[70] = 82
	MEM[71] = ADD
	MEM[72] = 129
	MEM[73] = STA
	MEM[74] = 131
	MEM[75] = JMP
	MEM[76] = 62
	MEM[77] = LDA
	MEM[78] = 134
	MEM[79] = STA
	MEM[80] = 132
	MEM[81] = HLT
	MEM[82] = ADD
	MEM[83] = 129
	MEM[84] = STA
	MEM[85] = 131
	MEM[86] = JN
	MEM[87] = 62
	MEM[88] = JMP
	MEM[89] = 77
	MEM[90] = LDA
	MEM[91] = 129
	MEM[92] = JN
	MEM[93] = 114
	MEM[94] = LDA
	MEM[95] = 130
	MEM[96] = JN
	MEM[97] = 144
	MEM[98] = ADD
	MEM[99] = 129
	MEM[100] = STA
	MEM[101] = 131
	MEM[102] = JN
	MEM[103] = 109
	MEM[104] = LDA
	MEM[105] = 135
	MEM[106] = STA
	MEM[107] = 132
	MEM[108] = HLT
	MEM[109] = LDA
	MEM[110] = 134
	MEM[111] = STA
	MEM[112] = 132
	MEM[113] = HLT
	MEM[114] = LDA
	MEM[115] = 130
	MEM[116] = JN
	MEM[117] = 138
	MEM[118] = ADD
	MEM[119] = 129
	MEM[120] = JN
	MEM[121] = 104
	MEM[122] = ADD
	MEM[123] = 134
	MEM[124] = STA
	MEM[125] = 131
	MEM[126] = JMP
	MEM[127] = 109
	MEM[128] = 1
	MEM[129] = 80
	MEM[130] = 130
	MEM[131] = 0
	MEM[132] = 0
	MEM[133] = 0
	MEM[134] = 1
	MEM[135] = 255
	MEM[136] = 0
	MEM[137] = 0
	MEM[138] = ADD
	MEM[139] = 129
	MEM[140] = JN
	MEM[141] = 109
	MEM[142] = JMP
	MEM[143] = 122
	MEM[144] = ADD
	MEM[145] = 129
	MEM[146] = JN
	MEM[147] = 104
	MEM[148] = JMP
	MEM[149] = 122
}*/
