// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Control
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ControlParser.
type ControlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ControlParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ControlParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ControlParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by ControlParser#reasignacion.
	VisitReasignacion(ctx *ReasignacionContext) interface{}

	// Visit a parse tree produced by ControlParser#asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by ControlParser#asignacionNoExpr.
	VisitAsignacionNoExpr(ctx *AsignacionNoExprContext) interface{}

	// Visit a parse tree produced by ControlParser#asignacionNoPrimitive.
	VisitAsignacionNoPrimitive(ctx *AsignacionNoPrimitiveContext) interface{}

	// Visit a parse tree produced by ControlParser#incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by ControlParser#decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}

	// Visit a parse tree produced by ControlParser#asignacionVectorVacio.
	VisitAsignacionVectorVacio(ctx *AsignacionVectorVacioContext) interface{}

	// Visit a parse tree produced by ControlParser#asignacionVector.
	VisitAsignacionVector(ctx *AsignacionVectorContext) interface{}

	// Visit a parse tree produced by ControlParser#reasignacionVector.
	VisitReasignacionVector(ctx *ReasignacionVectorContext) interface{}

	// Visit a parse tree produced by ControlParser#reasignacionMatrixTwoD.
	VisitReasignacionMatrixTwoD(ctx *ReasignacionMatrixTwoDContext) interface{}

	// Visit a parse tree produced by ControlParser#vectorAppend.
	VisitVectorAppend(ctx *VectorAppendContext) interface{}

	// Visit a parse tree produced by ControlParser#vectorRemoveLast.
	VisitVectorRemoveLast(ctx *VectorRemoveLastContext) interface{}

	// Visit a parse tree produced by ControlParser#vectorRemoveAt.
	VisitVectorRemoveAt(ctx *VectorRemoveAtContext) interface{}

	// Visit a parse tree produced by ControlParser#matrixTwoD.
	VisitMatrixTwoD(ctx *MatrixTwoDContext) interface{}

	// Visit a parse tree produced by ControlParser#matrixThreeD.
	VisitMatrixThreeD(ctx *MatrixThreeDContext) interface{}

	// Visit a parse tree produced by ControlParser#defMatrix.
	VisitDefMatrix(ctx *DefMatrixContext) interface{}

	// Visit a parse tree produced by ControlParser#listaValores_Mat.
	VisitListaValores_Mat(ctx *ListaValores_MatContext) interface{}

	// Visit a parse tree produced by ControlParser#lista_Expresiones.
	VisitLista_Expresiones(ctx *Lista_ExpresionesContext) interface{}

	// Visit a parse tree produced by ControlParser#listaValores_Mat2.
	VisitListaValores_Mat2(ctx *ListaValores_Mat2Context) interface{}

	// Visit a parse tree produced by ControlParser#func_sinRetorno.
	VisitFunc_sinRetorno(ctx *Func_sinRetornoContext) interface{}

	// Visit a parse tree produced by ControlParser#func_conRetorno_conTipo.
	VisitFunc_conRetorno_conTipo(ctx *Func_conRetorno_conTipoContext) interface{}

	// Visit a parse tree produced by ControlParser#callFunction.
	VisitCallFunction(ctx *CallFunctionContext) interface{}

	// Visit a parse tree produced by ControlParser#listaParamsCall.
	VisitListaParamsCall(ctx *ListaParamsCallContext) interface{}

	// Visit a parse tree produced by ControlParser#oneParam.
	VisitOneParam(ctx *OneParamContext) interface{}

	// Visit a parse tree produced by ControlParser#numParams.
	VisitNumParams(ctx *NumParamsContext) interface{}

	// Visit a parse tree produced by ControlParser#oneExpr.
	VisitOneExpr(ctx *OneExprContext) interface{}

	// Visit a parse tree produced by ControlParser#numExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by ControlParser#returnstmt.
	VisitReturnstmt(ctx *ReturnstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#printlnstmt.
	VisitPrintlnstmt(ctx *PrintlnstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#printstmt.
	VisitPrintstmt(ctx *PrintstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#else_if.
	VisitElse_if(ctx *Else_ifContext) interface{}

	// Visit a parse tree produced by ControlParser#else.
	VisitElse(ctx *ElseContext) interface{}

	// Visit a parse tree produced by ControlParser#ifNormal.
	VisitIfNormal(ctx *IfNormalContext) interface{}

	// Visit a parse tree produced by ControlParser#switchstmt.
	VisitSwitchstmt(ctx *SwitchstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#cases.
	VisitCases(ctx *CasesContext) interface{}

	// Visit a parse tree produced by ControlParser#unCase.
	VisitUnCase(ctx *UnCaseContext) interface{}

	// Visit a parse tree produced by ControlParser#unDefault.
	VisitUnDefault(ctx *UnDefaultContext) interface{}

	// Visit a parse tree produced by ControlParser#whilestmt.
	VisitWhilestmt(ctx *WhilestmtContext) interface{}

	// Visit a parse tree produced by ControlParser#forNormal.
	VisitForNormal(ctx *ForNormalContext) interface{}

	// Visit a parse tree produced by ControlParser#forEach.
	VisitForEach(ctx *ForEachContext) interface{}

	// Visit a parse tree produced by ControlParser#guardstmt.
	VisitGuardstmt(ctx *GuardstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by ControlParser#toInt.
	VisitToInt(ctx *ToIntContext) interface{}

	// Visit a parse tree produced by ControlParser#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by ControlParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by ControlParser#toFloat.
	VisitToFloat(ctx *ToFloatContext) interface{}

	// Visit a parse tree produced by ControlParser#vectorIsEmpty.
	VisitVectorIsEmpty(ctx *VectorIsEmptyContext) interface{}

	// Visit a parse tree produced by ControlParser#accesoMatrixTwoD.
	VisitAccesoMatrixTwoD(ctx *AccesoMatrixTwoDContext) interface{}

	// Visit a parse tree produced by ControlParser#nilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by ControlParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}

	// Visit a parse tree produced by ControlParser#CharExpr.
	VisitCharExpr(ctx *CharExprContext) interface{}

	// Visit a parse tree produced by ControlParser#vectorCount.
	VisitVectorCount(ctx *VectorCountContext) interface{}

	// Visit a parse tree produced by ControlParser#negExpr.
	VisitNegExpr(ctx *NegExprContext) interface{}

	// Visit a parse tree produced by ControlParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by ControlParser#callFuncAsExpr.
	VisitCallFuncAsExpr(ctx *CallFuncAsExprContext) interface{}

	// Visit a parse tree produced by ControlParser#vectorGetElement.
	VisitVectorGetElement(ctx *VectorGetElementContext) interface{}

	// Visit a parse tree produced by ControlParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by ControlParser#toString.
	VisitToString(ctx *ToStringContext) interface{}

	// Visit a parse tree produced by ControlParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by ControlParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by ControlParser#primitivo.
	VisitPrimitivo(ctx *PrimitivoContext) interface{}

	// Visit a parse tree produced by ControlParser#transfer_sentence.
	VisitTransfer_sentence(ctx *Transfer_sentenceContext) interface{}

	// Visit a parse tree produced by ControlParser#var_type.
	VisitVar_type(ctx *Var_typeContext) interface{}
}
