// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Control
import "github.com/antlr4-go/antlr/v4"

type BaseControlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseControlVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitReasignacion(ctx *ReasignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion(ctx *AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacionNoExpr(ctx *AsignacionNoExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacionNoPrimitive(ctx *AsignacionNoPrimitiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIncremento(ctx *IncrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDecremento(ctx *DecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacionVectorVacio(ctx *AsignacionVectorVacioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacionVector(ctx *AsignacionVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitReasignacionVector(ctx *ReasignacionVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitVectorAppend(ctx *VectorAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitVectorRemoveLast(ctx *VectorRemoveLastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitVectorRemoveAt(ctx *VectorRemoveAtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitFuncSinTipoRetorno(ctx *FuncSinTipoRetornoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitFuncParams_sinRetorno(ctx *FuncParams_sinRetornoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitFunc_conRetorno_conTipo(ctx *Func_conRetorno_conTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitFuncParams_ConRetorno(ctx *FuncParams_ConRetornoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitCallFunction(ctx *CallFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitCallFunctionParams(ctx *CallFunctionParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitListaParamsCall(ctx *ListaParamsCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitOneParam(ctx *OneParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitNumParams(ctx *NumParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitOneExpr(ctx *OneExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitNumExpr(ctx *NumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitPrintlnstmt(ctx *PrintlnstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitPrintstmt(ctx *PrintstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIfNormal(ctx *IfNormalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitElse(ctx *ElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitElse_if(ctx *Else_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSwitchstmt(ctx *SwitchstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitCases(ctx *CasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitUnCase(ctx *UnCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitUnDefault(ctx *UnDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitWhilestmt(ctx *WhilestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitForNormal(ctx *ForNormalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitForEach(ctx *ForEachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitGuardstmt(ctx *GuardstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitVectorGetElement(ctx *VectorGetElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitStrExpr(ctx *StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitVectorIsEmpty(ctx *VectorIsEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitOpExpr(ctx *OpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitCharExpr(ctx *CharExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitVectorCount(ctx *VectorCountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitPrimitivo(ctx *PrimitivoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitTransfer_sentence(ctx *Transfer_sentenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitVar_type(ctx *Var_typeContext) interface{} {
	return v.VisitChildren(ctx)
}
