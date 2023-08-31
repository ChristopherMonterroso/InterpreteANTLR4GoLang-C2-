// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Control
import "github.com/antlr4-go/antlr/v4"

// BaseControlListener is a complete listener for a parse tree produced by ControlParser.
type BaseControlListener struct{}

var _ ControlListener = &BaseControlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseControlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseControlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseControlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseControlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseControlListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseControlListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseControlListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseControlListener) ExitBlock(ctx *BlockContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseControlListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseControlListener) ExitStmt(ctx *StmtContext) {}

// EnterReasignacion is called when production reasignacion is entered.
func (s *BaseControlListener) EnterReasignacion(ctx *ReasignacionContext) {}

// ExitReasignacion is called when production reasignacion is exited.
func (s *BaseControlListener) ExitReasignacion(ctx *ReasignacionContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseControlListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseControlListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterAsignacionNoExpr is called when production asignacionNoExpr is entered.
func (s *BaseControlListener) EnterAsignacionNoExpr(ctx *AsignacionNoExprContext) {}

// ExitAsignacionNoExpr is called when production asignacionNoExpr is exited.
func (s *BaseControlListener) ExitAsignacionNoExpr(ctx *AsignacionNoExprContext) {}

// EnterAsignacionNoPrimitive is called when production asignacionNoPrimitive is entered.
func (s *BaseControlListener) EnterAsignacionNoPrimitive(ctx *AsignacionNoPrimitiveContext) {}

// ExitAsignacionNoPrimitive is called when production asignacionNoPrimitive is exited.
func (s *BaseControlListener) ExitAsignacionNoPrimitive(ctx *AsignacionNoPrimitiveContext) {}

// EnterIncremento is called when production incremento is entered.
func (s *BaseControlListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production incremento is exited.
func (s *BaseControlListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production decremento is entered.
func (s *BaseControlListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production decremento is exited.
func (s *BaseControlListener) ExitDecremento(ctx *DecrementoContext) {}

// EnterAsignacionVectorVacio is called when production asignacionVectorVacio is entered.
func (s *BaseControlListener) EnterAsignacionVectorVacio(ctx *AsignacionVectorVacioContext) {}

// ExitAsignacionVectorVacio is called when production asignacionVectorVacio is exited.
func (s *BaseControlListener) ExitAsignacionVectorVacio(ctx *AsignacionVectorVacioContext) {}

// EnterAsignacionVector is called when production asignacionVector is entered.
func (s *BaseControlListener) EnterAsignacionVector(ctx *AsignacionVectorContext) {}

// ExitAsignacionVector is called when production asignacionVector is exited.
func (s *BaseControlListener) ExitAsignacionVector(ctx *AsignacionVectorContext) {}

// EnterVectorAppend is called when production vectorAppend is entered.
func (s *BaseControlListener) EnterVectorAppend(ctx *VectorAppendContext) {}

// ExitVectorAppend is called when production vectorAppend is exited.
func (s *BaseControlListener) ExitVectorAppend(ctx *VectorAppendContext) {}

// EnterVectorRemoveLast is called when production vectorRemoveLast is entered.
func (s *BaseControlListener) EnterVectorRemoveLast(ctx *VectorRemoveLastContext) {}

// ExitVectorRemoveLast is called when production vectorRemoveLast is exited.
func (s *BaseControlListener) ExitVectorRemoveLast(ctx *VectorRemoveLastContext) {}

// EnterVectorRemoveAt is called when production vectorRemoveAt is entered.
func (s *BaseControlListener) EnterVectorRemoveAt(ctx *VectorRemoveAtContext) {}

// ExitVectorRemoveAt is called when production vectorRemoveAt is exited.
func (s *BaseControlListener) ExitVectorRemoveAt(ctx *VectorRemoveAtContext) {}

// EnterOneExpr is called when production oneExpr is entered.
func (s *BaseControlListener) EnterOneExpr(ctx *OneExprContext) {}

// ExitOneExpr is called when production oneExpr is exited.
func (s *BaseControlListener) ExitOneExpr(ctx *OneExprContext) {}

// EnterNumExpr is called when production numExpr is entered.
func (s *BaseControlListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production numExpr is exited.
func (s *BaseControlListener) ExitNumExpr(ctx *NumExprContext) {}

// EnterPrintlnstmt is called when production printlnstmt is entered.
func (s *BaseControlListener) EnterPrintlnstmt(ctx *PrintlnstmtContext) {}

// ExitPrintlnstmt is called when production printlnstmt is exited.
func (s *BaseControlListener) ExitPrintlnstmt(ctx *PrintlnstmtContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseControlListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseControlListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterIfNormal is called when production ifNormal is entered.
func (s *BaseControlListener) EnterIfNormal(ctx *IfNormalContext) {}

// ExitIfNormal is called when production ifNormal is exited.
func (s *BaseControlListener) ExitIfNormal(ctx *IfNormalContext) {}

// EnterElse is called when production else is entered.
func (s *BaseControlListener) EnterElse(ctx *ElseContext) {}

// ExitElse is called when production else is exited.
func (s *BaseControlListener) ExitElse(ctx *ElseContext) {}

// EnterElse_if is called when production else_if is entered.
func (s *BaseControlListener) EnterElse_if(ctx *Else_ifContext) {}

// ExitElse_if is called when production else_if is exited.
func (s *BaseControlListener) ExitElse_if(ctx *Else_ifContext) {}

// EnterSwitchstmt is called when production switchstmt is entered.
func (s *BaseControlListener) EnterSwitchstmt(ctx *SwitchstmtContext) {}

// ExitSwitchstmt is called when production switchstmt is exited.
func (s *BaseControlListener) ExitSwitchstmt(ctx *SwitchstmtContext) {}

// EnterCases is called when production cases is entered.
func (s *BaseControlListener) EnterCases(ctx *CasesContext) {}

// ExitCases is called when production cases is exited.
func (s *BaseControlListener) ExitCases(ctx *CasesContext) {}

// EnterUnCase is called when production unCase is entered.
func (s *BaseControlListener) EnterUnCase(ctx *UnCaseContext) {}

// ExitUnCase is called when production unCase is exited.
func (s *BaseControlListener) ExitUnCase(ctx *UnCaseContext) {}

// EnterUnDefault is called when production unDefault is entered.
func (s *BaseControlListener) EnterUnDefault(ctx *UnDefaultContext) {}

// ExitUnDefault is called when production unDefault is exited.
func (s *BaseControlListener) ExitUnDefault(ctx *UnDefaultContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseControlListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseControlListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterForNormal is called when production forNormal is entered.
func (s *BaseControlListener) EnterForNormal(ctx *ForNormalContext) {}

// ExitForNormal is called when production forNormal is exited.
func (s *BaseControlListener) ExitForNormal(ctx *ForNormalContext) {}

// EnterForEach is called when production forEach is entered.
func (s *BaseControlListener) EnterForEach(ctx *ForEachContext) {}

// ExitForEach is called when production forEach is exited.
func (s *BaseControlListener) ExitForEach(ctx *ForEachContext) {}

// EnterGuardstmt is called when production guardstmt is entered.
func (s *BaseControlListener) EnterGuardstmt(ctx *GuardstmtContext) {}

// ExitGuardstmt is called when production guardstmt is exited.
func (s *BaseControlListener) ExitGuardstmt(ctx *GuardstmtContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseControlListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseControlListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterFloatExpr is called when production FloatExpr is entered.
func (s *BaseControlListener) EnterFloatExpr(ctx *FloatExprContext) {}

// ExitFloatExpr is called when production FloatExpr is exited.
func (s *BaseControlListener) ExitFloatExpr(ctx *FloatExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseControlListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseControlListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterParExpr is called when production ParExpr is entered.
func (s *BaseControlListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production ParExpr is exited.
func (s *BaseControlListener) ExitParExpr(ctx *ParExprContext) {}

// EnterVectorGetElement is called when production vectorGetElement is entered.
func (s *BaseControlListener) EnterVectorGetElement(ctx *VectorGetElementContext) {}

// ExitVectorGetElement is called when production vectorGetElement is exited.
func (s *BaseControlListener) ExitVectorGetElement(ctx *VectorGetElementContext) {}

// EnterStrExpr is called when production StrExpr is entered.
func (s *BaseControlListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production StrExpr is exited.
func (s *BaseControlListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterVectorIsEmpty is called when production vectorIsEmpty is entered.
func (s *BaseControlListener) EnterVectorIsEmpty(ctx *VectorIsEmptyContext) {}

// ExitVectorIsEmpty is called when production vectorIsEmpty is exited.
func (s *BaseControlListener) ExitVectorIsEmpty(ctx *VectorIsEmptyContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseControlListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseControlListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseControlListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseControlListener) ExitIntExpr(ctx *IntExprContext) {}

// EnterOpExpr is called when production OpExpr is entered.
func (s *BaseControlListener) EnterOpExpr(ctx *OpExprContext) {}

// ExitOpExpr is called when production OpExpr is exited.
func (s *BaseControlListener) ExitOpExpr(ctx *OpExprContext) {}

// EnterCharExpr is called when production CharExpr is entered.
func (s *BaseControlListener) EnterCharExpr(ctx *CharExprContext) {}

// ExitCharExpr is called when production CharExpr is exited.
func (s *BaseControlListener) ExitCharExpr(ctx *CharExprContext) {}

// EnterVectorCount is called when production vectorCount is entered.
func (s *BaseControlListener) EnterVectorCount(ctx *VectorCountContext) {}

// ExitVectorCount is called when production vectorCount is exited.
func (s *BaseControlListener) ExitVectorCount(ctx *VectorCountContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseControlListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseControlListener) ExitPrimitivo(ctx *PrimitivoContext) {}

// EnterTransfer_sentence is called when production transfer_sentence is entered.
func (s *BaseControlListener) EnterTransfer_sentence(ctx *Transfer_sentenceContext) {}

// ExitTransfer_sentence is called when production transfer_sentence is exited.
func (s *BaseControlListener) ExitTransfer_sentence(ctx *Transfer_sentenceContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseControlListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseControlListener) ExitVar_type(ctx *Var_typeContext) {}
