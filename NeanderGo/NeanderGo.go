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
	MEM[0] = LDA
	MEM[1] = 133 // ZERO
	MEM[2] = STA
	MEM[3] = 132 // ERRO
	MEM[4] = STA
	MEM[5] = 131 // RESULTADO
	MEM[6] = LDA
	MEM[7] = 128 // OPER1
	MEM[8] = JN
	MEM[9] = 96 //erro1 // LABEL
	MEM[10] = LDA
	MEM[11] = 129 // OPER2
	MEM[12] = JN
	MEM[13] = 96 //erro1 // LABEL
	MEM[14] = LDA
	MEM[15] = 130 // OPERACAO
	MEM[16] = ADD
	MEM[17] = 137 // CONST -1
	MEM[18] = JZ
	MEM[19] = 34 //soma 	// LABEL
	MEM[20] = ADD
	MEM[21] = 137 // CONST -1
	MEM[22] = JZ
	MEM[23] = 42 //subtrai // LABEL
	MEM[24] = ADD
	MEM[25] = 137 // CONST -1
	MEM[26] = JZ
	MEM[27] = 53 //multiplica // LABEL
	MEM[28] = ADD
	MEM[29] = 137 // CONST -1
	MEM[30] = JZ
	MEM[31] = 69 //divide // LABEL
	MEM[32] = JMP
	MEM[33] = 102 //erro2 // LABEL

	//SOMA
	MEM[34] = LDA
	MEM[35] = 128
	MEM[36] = ADD
	MEM[37] = 129
	MEM[38] = STA
	MEM[39] = 131
	MEM[40] = JMP
	MEM[41] = 120 //fim // LABEL
	//SUBTRAI
	MEM[42] = LDA
	MEM[43] = 129 // OPER2
	MEM[44] = NOT
	MEM[45] = ADD
	MEM[46] = 134 // CONST1
	MEM[47] = ADD
	MEM[48] = 128 //OPER1
	MEM[49] = STA
	MEM[50] = 131 //RESULTADO
	MEM[51] = JMP
	MEM[52] = 120 //fim //LABEL
	//MULTIPLICA
	MEM[53] = LDA
	MEM[54] = 129 // OPER2
	MEM[55] = JZ
	MEM[56] = 120 //fim	// LABEL
	MEM[57] = ADD
	MEM[58] = 137 // CONST -1
	MEM[59] = STA
	MEM[60] = 129 // OPER2
	MEM[61] = LDA
	MEM[62] = 131 // RESULTADO
	MEM[63] = ADD
	MEM[64] = 128 // OPER1
	MEM[65] = STA
	MEM[66] = 131 // RESULTADO
	MEM[67] = JMP
	MEM[68] = 53 // MULTIPLICA
	//DIVIDE
	MEM[69] = LDA
	MEM[70] = 129 // OPER2
	MEM[71] = JZ
	MEM[72] = 108 //erro4	// LABEL
	MEM[73] = NOT
	MEM[74] = ADD
	MEM[75] = 134 // CONST1
	MEM[76] = STA
	MEM[77] = 129 // OPER2
	//LACOD
	MEM[78] = LDA
	MEM[79] = 128 // OPER1
	MEM[80] = JZ
	MEM[81] = 120 //fim	// LABEL
	MEM[82] = JN
	MEM[83] = 114 //subum	// LABEL
	MEM[84] = ADD
	MEM[85] = 129 // OPER2
	MEM[86] = STA
	MEM[87] = 128 // OPER1
	MEM[88] = LDA
	MEM[89] = 131 // RESULTADO
	MEM[90] = ADD
	MEM[91] = 134 // CONST1
	MEM[92] = STA
	MEM[93] = 131 // RESULTADO
	MEM[94] = JMP
	MEM[95] = 78 // LACOD
	//ERRO1
	MEM[96] = LDA
	MEM[97] = 134 // CONST1
	MEM[98] = STA
	MEM[99] = 132 // ERRO
	MEM[100] = JMP
	MEM[101] = 120 // FIM
	//ERRO2
	MEM[102] = LDA
	MEM[103] = 135 // CONST2
	MEM[104] = STA
	MEM[105] = 132 // ERRO
	MEM[106] = JMP
	MEM[107] = 120 // FIM
	//ERRO4
	MEM[108] = LDA
	MEM[109] = 136 // CONST4
	MEM[110] = STA
	MEM[111] = 132 // ERRO
	MEM[112] = JMP
	MEM[113] = 120 // FIM
	//SUBUM
	MEM[114] = LDA
	MEM[115] = 131 // RESULT
	MEM[116] = ADD
	MEM[117] = 137 // CONST -1
	MEM[118] = STA
	MEM[119] = 131 // RESULT
	//FIM
	MEM[120] = HLT

	MEM[128] = 47  // OPER1
	MEM[129] = 3   // OPER2
	MEM[130] = 3   // OPERACAO
	MEM[131] = 0   // RESULTADO
	MEM[132] = 0   // ERRO
	MEM[133] = 0   // CONST 0
	MEM[134] = 1   // CONST 1
	MEM[135] = 2   // CONST 2
	MEM[136] = 4   // CONST 4
	MEM[137] = 255 // CONST -1
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

/*
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
	HLT
}*/
