// Code generated by "stringer -type=Op"; DO NOT EDIT.

package gno

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpInvalid-0]
	_ = x[OpHalt-1]
	_ = x[OpNoop-2]
	_ = x[OpExec-3]
	_ = x[OpPrecall-4]
	_ = x[OpCall-5]
	_ = x[OpCallNativeBody-6]
	_ = x[OpReturn-7]
	_ = x[OpReturnFromBlock-8]
	_ = x[OpReturnToBlock-9]
	_ = x[OpDefer-10]
	_ = x[OpGo-11]
	_ = x[OpSelectCase-12]
	_ = x[OpSwitchCase-13]
	_ = x[OpTypeSwitchCase-14]
	_ = x[OpForLoop1-15]
	_ = x[OpIfCond-16]
	_ = x[OpPopValue-17]
	_ = x[OpPopResults-18]
	_ = x[OpPopBlock-19]
	_ = x[OpUpos-32]
	_ = x[OpUneg-33]
	_ = x[OpUnot-34]
	_ = x[OpUxor-35]
	_ = x[OpUrecv-37]
	_ = x[OpLor-38]
	_ = x[OpLand-39]
	_ = x[OpEql-40]
	_ = x[OpNeq-41]
	_ = x[OpLss-42]
	_ = x[OpLeq-43]
	_ = x[OpGtr-44]
	_ = x[OpGeq-45]
	_ = x[OpAdd-46]
	_ = x[OpSub-47]
	_ = x[OpBor-48]
	_ = x[OpXor-49]
	_ = x[OpMul-50]
	_ = x[OpQuo-51]
	_ = x[OpRem-52]
	_ = x[OpShl-53]
	_ = x[OpShr-54]
	_ = x[OpBand-55]
	_ = x[OpBandn-56]
	_ = x[OpEval-64]
	_ = x[OpBinary1-65]
	_ = x[OpIndex-66]
	_ = x[OpSelector-67]
	_ = x[OpSlice-68]
	_ = x[OpStar-69]
	_ = x[OpRef-70]
	_ = x[OpTypeAssert1-71]
	_ = x[OpTypeAssert2-72]
	_ = x[OpTypeOf-73]
	_ = x[OpCompositeLit-74]
	_ = x[OpArrayLit-75]
	_ = x[OpSliceLit-76]
	_ = x[OpMapLit-77]
	_ = x[OpStructLit-78]
	_ = x[OpFuncLit-79]
	_ = x[OpConvert-80]
	_ = x[OpStructLitGoNative-96]
	_ = x[OpCallGoNative-97]
	_ = x[OpFieldType-112]
	_ = x[OpArrayType-113]
	_ = x[OpSliceType-114]
	_ = x[OpPointerType-115]
	_ = x[OpInterfaceType-116]
	_ = x[OpChanType-117]
	_ = x[OpFuncType-118]
	_ = x[OpMapType-119]
	_ = x[OpStructType-120]
	_ = x[OpAssign-128]
	_ = x[OpAddAssign-129]
	_ = x[OpSubAssign-130]
	_ = x[OpMulAssign-131]
	_ = x[OpQuoAssign-132]
	_ = x[OpRemAssign-133]
	_ = x[OpBandAssign-134]
	_ = x[OpBandnAssign-135]
	_ = x[OpBorAssign-136]
	_ = x[OpXorAssign-137]
	_ = x[OpShlAssign-138]
	_ = x[OpShrAssign-139]
	_ = x[OpDefine-140]
	_ = x[OpInc-141]
	_ = x[OpDec-142]
	_ = x[OpValueDecl-144]
	_ = x[OpSticky-208]
	_ = x[OpForLoop2-208]
	_ = x[OpRangeIter-209]
	_ = x[OpReturnCallDefers-210]
}

const (
	_Op_name_0 = "OpInvalidOpHaltOpNoopOpExecOpPrecallOpCallOpCallNativeBodyOpReturnOpReturnFromBlockOpReturnToBlockOpDeferOpGoOpSelectCaseOpSwitchCaseOpTypeSwitchCaseOpForLoop1OpIfCondOpPopValueOpPopResultsOpPopBlock"
	_Op_name_1 = "OpUposOpUnegOpUnotOpUxor"
	_Op_name_2 = "OpUrecvOpLorOpLandOpEqlOpNeqOpLssOpLeqOpGtrOpGeqOpAddOpSubOpBorOpXorOpMulOpQuoOpRemOpShlOpShrOpBandOpBandn"
	_Op_name_3 = "OpEvalOpBinary1OpIndexOpSelectorOpSliceOpStarOpRefOpTypeAssert1OpTypeAssert2OpTypeOfOpCompositeLitOpArrayLitOpSliceLitOpMapLitOpStructLitOpFuncLitOpConvert"
	_Op_name_4 = "OpStructLitGoNativeOpCallGoNative"
	_Op_name_5 = "OpFieldTypeOpArrayTypeOpSliceTypeOpPointerTypeOpInterfaceTypeOpChanTypeOpFuncTypeOpMapTypeOpStructType"
	_Op_name_6 = "OpAssignOpAddAssignOpSubAssignOpMulAssignOpQuoAssignOpRemAssignOpBandAssignOpBandnAssignOpBorAssignOpXorAssignOpShlAssignOpShrAssignOpDefineOpIncOpDec"
	_Op_name_7 = "OpValueDecl"
	_Op_name_8 = "OpStickyOpRangeIterOpReturnCallDefers"
)

var (
	_Op_index_0 = [...]uint8{0, 9, 15, 21, 27, 36, 42, 58, 66, 83, 98, 105, 109, 121, 133, 149, 159, 167, 177, 189, 199}
	_Op_index_1 = [...]uint8{0, 6, 12, 18, 24}
	_Op_index_2 = [...]uint8{0, 7, 12, 18, 23, 28, 33, 38, 43, 48, 53, 58, 63, 68, 73, 78, 83, 88, 93, 99, 106}
	_Op_index_3 = [...]uint8{0, 6, 15, 22, 32, 39, 45, 50, 63, 76, 84, 98, 108, 118, 126, 137, 146, 155}
	_Op_index_4 = [...]uint8{0, 19, 33}
	_Op_index_5 = [...]uint8{0, 11, 22, 33, 46, 61, 71, 81, 90, 102}
	_Op_index_6 = [...]uint8{0, 8, 19, 30, 41, 52, 63, 75, 88, 99, 110, 121, 132, 140, 145, 150}
	_Op_index_8 = [...]uint8{0, 8, 19, 37}
)

func (i Op) String() string {
	switch {
	case 0 <= i && i <= 19:
		return _Op_name_0[_Op_index_0[i]:_Op_index_0[i+1]]
	case 32 <= i && i <= 35:
		i -= 32
		return _Op_name_1[_Op_index_1[i]:_Op_index_1[i+1]]
	case 37 <= i && i <= 56:
		i -= 37
		return _Op_name_2[_Op_index_2[i]:_Op_index_2[i+1]]
	case 64 <= i && i <= 80:
		i -= 64
		return _Op_name_3[_Op_index_3[i]:_Op_index_3[i+1]]
	case 96 <= i && i <= 97:
		i -= 96
		return _Op_name_4[_Op_index_4[i]:_Op_index_4[i+1]]
	case 112 <= i && i <= 120:
		i -= 112
		return _Op_name_5[_Op_index_5[i]:_Op_index_5[i+1]]
	case 128 <= i && i <= 142:
		i -= 128
		return _Op_name_6[_Op_index_6[i]:_Op_index_6[i+1]]
	case i == 144:
		return _Op_name_7
	case 208 <= i && i <= 210:
		i -= 208
		return _Op_name_8[_Op_index_8[i]:_Op_index_8[i+1]]
	default:
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
