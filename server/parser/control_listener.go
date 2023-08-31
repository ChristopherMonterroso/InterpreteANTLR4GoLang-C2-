// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Control
import "github.com/antlr4-go/antlr/v4"

// ControlListener is a complete listener for a parse tree produced by ControlParser.
type ControlListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterReasignacion is called when entering the reasignacion production.
	EnterReasignacion(c *ReasignacionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterAsignacionNoExpr is called when entering the asignacionNoExpr production.
	EnterAsignacionNoExpr(c *AsignacionNoExprContext)

	// EnterAsignacionNoPrimitive is called when entering the asignacionNoPrimitive production.
	EnterAsignacionNoPrimitive(c *AsignacionNoPrimitiveContext)

	// EnterIncremento is called when entering the incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the decremento production.
	EnterDecremento(c *DecrementoContext)

	// EnterAsignacionVectorVacio is called when entering the asignacionVectorVacio production.
	EnterAsignacionVectorVacio(c *AsignacionVectorVacioContext)

	// EnterAsignacionVector is called when entering the asignacionVector production.
	EnterAsignacionVector(c *AsignacionVectorContext)

	// EnterVectorAppend is called when entering the vectorAppend production.
	EnterVectorAppend(c *VectorAppendContext)

	// EnterVectorRemoveLast is called when entering the vectorRemoveLast production.
	EnterVectorRemoveLast(c *VectorRemoveLastContext)

	// EnterVectorRemoveAt is called when entering the vectorRemoveAt production.
	EnterVectorRemoveAt(c *VectorRemoveAtContext)

	// EnterOneExpr is called when entering the oneExpr production.
	EnterOneExpr(c *OneExprContext)

	// EnterNumExpr is called when entering the numExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterPrintlnstmt is called when entering the printlnstmt production.
	EnterPrintlnstmt(c *PrintlnstmtContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIfNormal is called when entering the ifNormal production.
	EnterIfNormal(c *IfNormalContext)

	// EnterElse is called when entering the else production.
	EnterElse(c *ElseContext)

	// EnterElse_if is called when entering the else_if production.
	EnterElse_if(c *Else_ifContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterCases is called when entering the cases production.
	EnterCases(c *CasesContext)

	// EnterUnCase is called when entering the unCase production.
	EnterUnCase(c *UnCaseContext)

	// EnterUnDefault is called when entering the unDefault production.
	EnterUnDefault(c *UnDefaultContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterForNormal is called when entering the forNormal production.
	EnterForNormal(c *ForNormalContext)

	// EnterForEach is called when entering the forEach production.
	EnterForEach(c *ForEachContext)

	// EnterGuardstmt is called when entering the guardstmt production.
	EnterGuardstmt(c *GuardstmtContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterFloatExpr is called when entering the FloatExpr production.
	EnterFloatExpr(c *FloatExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterVectorGetElement is called when entering the vectorGetElement production.
	EnterVectorGetElement(c *VectorGetElementContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterVectorIsEmpty is called when entering the vectorIsEmpty production.
	EnterVectorIsEmpty(c *VectorIsEmptyContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterCharExpr is called when entering the CharExpr production.
	EnterCharExpr(c *CharExprContext)

	// EnterVectorCount is called when entering the vectorCount production.
	EnterVectorCount(c *VectorCountContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// EnterTransfer_sentence is called when entering the transfer_sentence production.
	EnterTransfer_sentence(c *Transfer_sentenceContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitReasignacion is called when exiting the reasignacion production.
	ExitReasignacion(c *ReasignacionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitAsignacionNoExpr is called when exiting the asignacionNoExpr production.
	ExitAsignacionNoExpr(c *AsignacionNoExprContext)

	// ExitAsignacionNoPrimitive is called when exiting the asignacionNoPrimitive production.
	ExitAsignacionNoPrimitive(c *AsignacionNoPrimitiveContext)

	// ExitIncremento is called when exiting the incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the decremento production.
	ExitDecremento(c *DecrementoContext)

	// ExitAsignacionVectorVacio is called when exiting the asignacionVectorVacio production.
	ExitAsignacionVectorVacio(c *AsignacionVectorVacioContext)

	// ExitAsignacionVector is called when exiting the asignacionVector production.
	ExitAsignacionVector(c *AsignacionVectorContext)

	// ExitVectorAppend is called when exiting the vectorAppend production.
	ExitVectorAppend(c *VectorAppendContext)

	// ExitVectorRemoveLast is called when exiting the vectorRemoveLast production.
	ExitVectorRemoveLast(c *VectorRemoveLastContext)

	// ExitVectorRemoveAt is called when exiting the vectorRemoveAt production.
	ExitVectorRemoveAt(c *VectorRemoveAtContext)

	// ExitOneExpr is called when exiting the oneExpr production.
	ExitOneExpr(c *OneExprContext)

	// ExitNumExpr is called when exiting the numExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitPrintlnstmt is called when exiting the printlnstmt production.
	ExitPrintlnstmt(c *PrintlnstmtContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIfNormal is called when exiting the ifNormal production.
	ExitIfNormal(c *IfNormalContext)

	// ExitElse is called when exiting the else production.
	ExitElse(c *ElseContext)

	// ExitElse_if is called when exiting the else_if production.
	ExitElse_if(c *Else_ifContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitCases is called when exiting the cases production.
	ExitCases(c *CasesContext)

	// ExitUnCase is called when exiting the unCase production.
	ExitUnCase(c *UnCaseContext)

	// ExitUnDefault is called when exiting the unDefault production.
	ExitUnDefault(c *UnDefaultContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitForNormal is called when exiting the forNormal production.
	ExitForNormal(c *ForNormalContext)

	// ExitForEach is called when exiting the forEach production.
	ExitForEach(c *ForEachContext)

	// ExitGuardstmt is called when exiting the guardstmt production.
	ExitGuardstmt(c *GuardstmtContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitFloatExpr is called when exiting the FloatExpr production.
	ExitFloatExpr(c *FloatExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitVectorGetElement is called when exiting the vectorGetElement production.
	ExitVectorGetElement(c *VectorGetElementContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitVectorIsEmpty is called when exiting the vectorIsEmpty production.
	ExitVectorIsEmpty(c *VectorIsEmptyContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitCharExpr is called when exiting the CharExpr production.
	ExitCharExpr(c *CharExprContext)

	// ExitVectorCount is called when exiting the vectorCount production.
	ExitVectorCount(c *VectorCountContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)

	// ExitTransfer_sentence is called when exiting the transfer_sentence production.
	ExitTransfer_sentence(c *Transfer_sentenceContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)
}
