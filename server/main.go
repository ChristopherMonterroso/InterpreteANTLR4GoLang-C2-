package main

import (
	"fmt"
	"strconv"
	"strings"

	"Server/parser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Value struct {
	ambit         string
	PrimitiveType string
	value         interface{}
	isList        bool
}

type Visitor struct {
	parser.ControlVisitor
	memory        map[string]Value
	exprToSwitch  interface{}
	exitSwitch    bool
	outputBuilder strings.Builder
	Errors        strings.Builder
	continueFlag  bool
	breakFlag     bool
	tmpList       []interface{}
}

func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StmtContext:
		return v.VisitStmt(val)
	case *parser.AsignacionContext:
		return v.VisitAsignacion(val)
	case *parser.AsignacionNoExprContext:
		return v.VisitAsignacionNoExpr(val)
	case *parser.AsignacionNoPrimitiveContext:
		return v.VisitAsignacionNoPrimitive(val)
	case *parser.ReasignacionContext:
		return v.VisitReasignacion(val)
	case *parser.IncrementoContext:
		return v.visitIncremento(val)
	case *parser.DecrementoContext:
		return v.VisitDecremento(val)
	case *parser.IfNormalContext:
		return v.VisitIfNormal(val)

	case *parser.AsignacionVectorVacioContext:
		return v.VisitAsignacionVectorVacio(val)
	case *parser.AsignacionVectorContext:
		return v.VisitAsignacionVector(val)
	case *parser.OneExprContext:
		return v.VisitOneExpr(val)
	case *parser.NumExprContext:
		return v.VisitNumExpr(val)
	case *parser.VectorAppendContext:
		return v.VisitVectorAppend(val)
	case *parser.VectorCountContext:
		return v.visitVectorCount(val)
	case *parser.VectorIsEmptyContext:
		return v.VisitVectorIsEmpty(val)
	case *parser.VectorGetElementContext:
		return v.VisitVectorGetElement(val)

	case *parser.ElseContext:
		return v.VisitIfElse(val)
	case *parser.Else_ifContext:
		return v.VisitElse_If(val)

	case *parser.SwitchstmtContext:
		return v.VisitSwitchstmt(val)
	case *parser.CasesContext:
		return v.VisitCases(val)
	/*case * parser.CaseblockContext:
	return v.VisitCaseblock(val)*/
	case *parser.UnCaseContext:
		return v.VisitUnCase(val)
	case *parser.UnDefaultContext:
		return v.VisitUnDefault(val)
	case *parser.PrintlnstmtContext:
		return v.VisitPrintlnstmt(val)
	case *parser.WhilestmtContext:
		return v.VisitWhilestmt(val)

	case *parser.ForNormalContext:
		return v.VisitForNormal(val)
	case *parser.ForEachContext:
		return v.VisitForEach(val)

	case *parser.GuardstmtContext:
		return v.VisitGuardstmt(val)

	case *parser.PrintstmtContext:
		return v.VisitPrintstmt(val)

	case *parser.OpExprContext:
		return v.VisitOpExpr(val)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val)
	case *parser.FloatExprContext:
		return v.VisitFloatExpr(val)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val)
	case *parser.CharExprContext:
		return v.VisitCharExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	case *parser.Var_typeContext:
		return v.VisitVar_type(val)
	case *parser.PrimitivoContext:
		return v.VisitPrimitivo(val)
	case *parser.Transfer_sentenceContext:
		return v.VisitTransferSentence(val)
	default:
		panic("Unknown context " + val.GetText())
	}

}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) Value {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value {
	for i := 0; ctx.Stmt(i) != nil; i++ {
		v.Visit(ctx.Stmt(i))
		if v.continueFlag || v.breakFlag {
			v.continueFlag = false
			break
		}
	}
	return Value{value: true}
}

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) Value {
	if ctx.Assignstmt() != nil {
		return v.Visit(ctx.Assignstmt())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Printlnstmt() != nil {
		return v.Visit(ctx.Printlnstmt())
	}
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}
	if ctx.Whilestmt() != nil {
		return v.Visit(ctx.Whilestmt())
	}
	if ctx.Switchstmt() != nil {
		return v.Visit(ctx.Switchstmt())
	}
	if ctx.Forstmt() != nil {
		return v.Visit(ctx.Forstmt())
	}
	if ctx.Guardstmt() != nil {
		return v.Visit(ctx.Guardstmt())
	}
	if ctx.VectorPpts() != nil {
		return v.Visit(ctx.VectorPpts())
	}
	return Value{value: true}
}

// Funciones de asignacion
func (v *Visitor) VisitAsignacion(ctx *parser.AsignacionContext) Value {
	ambit := v.Visit(ctx.Var_type()).value //Local o global
	varID := ctx.ID().GetText()
	primitivetype := v.Visit(ctx.Primitivo())
	value := v.Visit(ctx.Expr())
	if v.DoesVarExist(varID) {
		v.outputBuilder.WriteString(fmt.Sprintf("%v\n", "La variable "+varID+" ya existe"))
		return Value{value: false}
	} else {
		if primitivetype.value.(string) == getType(value.value) { // Confirmar que el tipo primitivo asignado sea igual al tipo primitivo de la expresión
			v.memory[varID] = Value{
				ambit:         ambit.(string),
				PrimitiveType: primitivetype.value.(string),
				value:         value.value,
			}
			return Value{value: true} //Continua el flujo del interprete
		} else {

			return Value{value: false}
		}
	}

}

func (v *Visitor) VisitAsignacionNoExpr(ctx *parser.AsignacionNoExprContext) Value {
	ambit := v.Visit(ctx.Var_type()).value //Local o global
	varID := ctx.ID().GetText()
	primitivetype := v.Visit(ctx.Primitivo())

	v.memory[varID] = Value{
		ambit:         ambit.(string),
		PrimitiveType: primitivetype.value.(string),
		value:         nil,
	}
	return Value{value: true}

}

func (v *Visitor) VisitAsignacionNoPrimitive(ctx *parser.AsignacionNoPrimitiveContext) Value {
	ambit := v.Visit(ctx.Var_type()).value //Local o global
	varID := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.memory[varID] = Value{
		ambit:         ambit.(string),
		PrimitiveType: getType(value.value),
		value:         value.value,
	}
	return Value{value: true}

}

func (v *Visitor) VisitReasignacion(ctx *parser.ReasignacionContext) Value {
	varID := ctx.ID().GetText()
	ambit := v.memory[varID].ambit
	primitivetype := v.memory[varID].PrimitiveType

	currentValue := v.memory[varID].value
	valueToAssign := v.Visit(ctx.Expr())
	fmt.Println(varID, ambit, primitivetype, valueToAssign)
	fmt.Println(getType(valueToAssign.value))
	if v.DoesVarExist(varID) {
		if ambit == "let" {
			return Value{value: false}

		}
		if primitivetype == getType(valueToAssign.value) || currentValue == nil {

			v.memory[varID] = Value{
				ambit:         ambit,
				PrimitiveType: primitivetype,
				value:         valueToAssign.value,
			}
		} else {
			return Value{value: false}
		}
	} else {
		return Value{value: false}
	}
	return Value{value: true}

}

func (v *Visitor) VisitAsignacionVectorVacio(ctx *parser.AsignacionVectorVacioContext) Value {
	varID := ctx.ID().GetText()
	primitivetype := v.Visit(ctx.Primitivo())

	// Crear un slice vacío basado en el tipo primitivo
	var emptySlice interface{}
	switch primitivetype.value.(string) {
	case "Int":
		emptySlice = []int64{}
	case "Float":
		emptySlice = []float64{}
	case "String":
		emptySlice = []string{}
	// Agrega más tipos aquí si es necesario
	default:
		// Tipo desconocido, manejar el error apropiadamente
	}

	v.memory[varID] = Value{
		ambit:         "var",
		PrimitiveType: primitivetype.value.(string),
		value:         emptySlice,
	}
	return Value{value: true}
}

func (v *Visitor) VisitAsignacionVector(ctx *parser.AsignacionVectorContext) Value {
	varID := ctx.ID().GetText()
	primitivoType := v.Visit(ctx.Primitivo()).value.(string)
	v.Visit(ctx.ListaExp())
	// Parse the tmpList based on the primitive type
	var convertedSlice interface{}
	switch primitivoType {
	case "Int":
		intSlice := make([]int64, len(v.tmpList))
		for i, val := range v.tmpList {
			intSlice[i] = val.(int64)
		}
		convertedSlice = intSlice
	case "String":
		stringSlice := make([]string, len(v.tmpList))
		for i, val := range v.tmpList {
			stringSlice[i] = val.(string)
		}
		convertedSlice = stringSlice
	case "Float":
		floatSlice := make([]float64, len(v.tmpList))
		for i, val := range v.tmpList {
			floatSlice[i] = val.(float64)
		}
		convertedSlice = floatSlice
	// Add more cases for other primitive types if needed
	default:
		// Handle unsupported types
		fmt.Println("Unsupported primitive type:", primitivoType)
		convertedSlice = nil
	}

	v.memory[varID] = Value{
		ambit:         "var",
		PrimitiveType: primitivoType,
		value:         convertedSlice,
		isList:        true,
	}
	fmt.Println("vector agregado", v.tmpList)
	// Clear tmpList after using it
	v.tmpList = nil

	return Value{value: true}
}

func (v *Visitor) visitIncremento(ctx *parser.IncrementoContext) Value {
	varID := ctx.ID().GetText()
	ambit := v.memory[varID].ambit
	primitivetype := v.memory[varID].PrimitiveType
	currentValue := v.memory[varID].value
	valueToAssign := v.Visit(ctx.Expr()).value

	if v.DoesVarExist(varID) {
		if ambit == "let" {
			return Value{value: false}

		}
		if primitivetype == getType(currentValue) {
			if primitivetype == "Int" {

				v.memory[varID] = Value{
					ambit:         ambit,
					PrimitiveType: primitivetype,
					value:         currentValue.(int64) + valueToAssign.(int64),
				}
			} else if primitivetype == "Float" {
				v.memory[varID] = Value{
					ambit:         ambit,
					PrimitiveType: primitivetype,
					value:         currentValue.(float64) + valueToAssign.(float64),
				}
			} else if primitivetype == "String" {
				v.memory[varID] = Value{
					ambit:         ambit,
					PrimitiveType: primitivetype,
					value:         currentValue.(string) + valueToAssign.(string),
				}
			}
		} else {
			return Value{value: false}
		}
	} else {
		return Value{value: false}
	}
	return Value{value: true}

}

func (v *Visitor) VisitDecremento(ctx *parser.DecrementoContext) Value {
	varID := ctx.ID().GetText()
	ambit := v.memory[varID].ambit
	primitivetype := v.memory[varID].PrimitiveType

	currentValue := v.memory[varID].value
	valueToAssign := v.Visit(ctx.Expr()).value

	if v.DoesVarExist(varID) {
		if ambit == "let" {
			return Value{value: false}

		}
		if primitivetype == getType(currentValue) {
			if primitivetype == "Int" {

				v.memory[varID] = Value{
					ambit:         ambit,
					PrimitiveType: primitivetype,
					value:         currentValue.(int64) - valueToAssign.(int64),
				}
			} else if primitivetype == "Float" {
				v.memory[varID] = Value{
					ambit:         ambit,
					PrimitiveType: primitivetype,
					value:         currentValue.(float64) - valueToAssign.(float64),
				}
			}
		} else {
			return Value{value: false}
		}
	} else {
		return Value{value: false}
	}
	return Value{value: true}
}

func (v *Visitor) VisitIfNormal(ctx *parser.IfNormalContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	if ok && value {
		return v.Visit(ctx.Block())
	}
	return Value{value: false}
}

func (v *Visitor) VisitIfElse(ctx *parser.ElseContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	if ok && value {
		return v.Visit(ctx.Block(0))
	} else {
		return v.Visit(ctx.Block(1))
	}
}
func (v *Visitor) VisitElse_If(ctx *parser.Else_ifContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	if ok && value {
		return v.Visit(ctx.Block())
	} else {
		return v.Visit(ctx.Ifstmt())
	}

}

func (v *Visitor) VisitSwitchstmt(ctx *parser.SwitchstmtContext) Value {
	expr := v.Visit(ctx.Expr()).value
	v.exprToSwitch = expr
	if ctx.Cases() != nil {
		return v.Visit(ctx.Cases())
	}
	return Value{value: true}

}
func (v *Visitor) VisitCases(ctx *parser.CasesContext) Value {
	for i := 0; ctx.Caseblock(i) != nil; i++ {
		v.exitSwitch = false
		v.Visit(ctx.Caseblock(i))
		if v.exitSwitch {
			v.exitSwitch = false
			break
		}
	}

	return Value{value: true}
}
func (v *Visitor) VisitUnCase(ctx *parser.UnCaseContext) Value {
	value := v.Visit(ctx.Expr()).value
	if value == v.exprToSwitch && getType(value) == getType(v.exprToSwitch) {
		v.Visit(ctx.Block())
		v.exitSwitch = true
		return Value{value: true}
	} else {
		return Value{value: false}
	}

}
func (v *Visitor) VisitUnDefault(ctx *parser.UnDefaultContext) Value {
	if ctx.Block() != nil {
		return v.Visit(ctx.Block())
	}
	return Value{value: false}
}
func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) Value {
	value := v.Visit(ctx.Expr()).value
	if getType(value) == "Character" {
		value = asciiToChar(value)
	}
	v.outputBuilder.WriteString(fmt.Sprintf("%v\n", value))
	return Value{value: true}
}

func (v *Visitor) VisitPrintlnstmt(ctx *parser.PrintlnstmtContext) Value {
	fmt.Println(v.Visit(ctx.Expr()).value)
	return Value{value: true}
}

func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	for ok && value {
		v.Visit(ctx.Block())
		if v.breakFlag { // si encuentra un break se termina el ciclo
			v.breakFlag = false
			break
		}
		value, ok = v.Visit(ctx.Expr()).value.(bool)

	}
	return Value{value: true}
}
func (v *Visitor) VisitForNormal(ctx *parser.ForNormalContext) Value {
	varID := ctx.ID().GetText()
	exp1 := v.Visit(ctx.Expr(0)).value
	exp2 := v.Visit(ctx.Expr(1)).value

	if exp1.(int64) > exp2.(int64) {
		return Value{value: false}
	} else {
		for exp1.(int64) <= exp2.(int64) {
			v.memory[varID] = Value{
				ambit:         "var",
				PrimitiveType: "Int",
				value:         exp1.(int64),
			}
			loopBlock := ctx.Block()

			v.Visit(loopBlock)

			if v.breakFlag { // si encuentra un break se termina el ciclo
				v.breakFlag = false
				break
			}
			exp1 = exp1.(int64) + 1

		}
		return Value{value: true}
	}
}
func (v *Visitor) VisitForEach(ctx *parser.ForEachContext) Value {
	varID := ctx.ID().GetText()
	expr := v.Visit(ctx.Expr()).value
	if getType(expr) == "String" {
		v.memory[varID] = Value{
			ambit:         "var",
			PrimitiveType: "Character",
			value:         nil,
		}
		for _, char := range expr.(string) {
			v.memory[varID] = Value{
				ambit:         "var",
				PrimitiveType: "Character",
				value:         char,
			}
			v.Visit(ctx.Block())
			if v.breakFlag { // si encuentra un break se termina el ciclo
				v.breakFlag = false
				break
			}

		}
		return Value{value: true}
	}
	return Value{value: false}
}
func (v *Visitor) VisitGuardstmt(ctx *parser.GuardstmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	if !value && ok {
		v.Visit(ctx.Block())
		hasTransferStc := ctx.Transfer_sentence().GetText()
		if hasTransferStc == "continue" {
			v.continueFlag = true //Se encontró el continue xd
		}
		if hasTransferStc == "break" {
			v.breakFlag = true //
		}
	}
	return Value{value: false}
}

func (v *Visitor) VisitVectorAppend(ctx *parser.VectorAppendContext) Value {
	varID := ctx.ID().GetText()
	exprValue := v.Visit(ctx.Expr()).value
	primitive := getType(exprValue)

	if v.memory[varID].isList {
		if v.memory[varID].PrimitiveType == primitive {
			switch v.memory[varID].PrimitiveType {
			case "String":
				sliceValue, ok := v.memory[varID].value.([]string)
				if !ok {
					return Value{value: false}
				}
				newValue := exprValue.(string)
				sliceValue = append(sliceValue, newValue)

				// Actualizar el valor en el mapa
				v.memory[varID] = Value{
					ambit:         v.memory[varID].ambit,
					PrimitiveType: v.memory[varID].PrimitiveType,
					value:         sliceValue,
					isList:        v.memory[varID].isList,
				}

			case "Int":
				sliceValue, ok := v.memory[varID].value.([]int64)
				if !ok {
					return Value{value: false}
				}
				newValue := exprValue.(int64)
				sliceValue = append(sliceValue, newValue)

				// Actualizar el valor en el mapa
				v.memory[varID] = Value{
					ambit:         v.memory[varID].ambit,
					PrimitiveType: v.memory[varID].PrimitiveType,
					value:         sliceValue,
					isList:        v.memory[varID].isList,
				}

			case "Float":
				sliceValue, ok := v.memory[varID].value.([]float64)
				if !ok {
					return Value{value: false}
				}
				newValue := exprValue.(float64)
				sliceValue = append(sliceValue, newValue)

				// Actualizar el valor en el mapa
				v.memory[varID] = Value{
					ambit:         v.memory[varID].ambit,
					PrimitiveType: v.memory[varID].PrimitiveType,
					value:         sliceValue,
					isList:        v.memory[varID].isList,
				}

			default:
				fmt.Println("tipo primitivo desconocido")
				return Value{value: false}
			}

			fmt.Printf("Nuevo valor agregado al slice de %s: %v\n", v.memory[varID].PrimitiveType, v.memory[varID].value)
			return Value{value: true}
		} else {
			fmt.Println("no coincide el tipo de expresión con el tipo del vector")
			return Value{value: false}
		}
	} else {
		fmt.Println("variable no es un slice")
		return Value{value: false}
	}
}

func (v *Visitor) visitVectorCount(ctx *parser.VectorCountContext) Value {
	typePrimitive := v.memory[ctx.ID().GetText()].PrimitiveType
	varID := ctx.ID().GetText()
	getSliceSize := func(v Value) (int, error) {
		switch typePrimitive {
		case "String":
			sliceValue, ok := v.value.([]string)
			if !ok {
				return 0, fmt.Errorf("tipo de slice incorrecto")
			}
			return len(sliceValue), nil
		case "Int":
			sliceValue, ok := v.value.([]int64)
			if !ok {
				return 0, fmt.Errorf("tipo de slice incorrecto")
			}
			return len(sliceValue), nil
		case "Float":
			sliceValue, ok := v.value.([]float64)
			if !ok {
				return 0, fmt.Errorf("tipo de slice incorrecto")
			}
			return len(sliceValue), nil
		default:
			return 0, fmt.Errorf("tipo primitivo desconocido")
		}
	}

	// Obtener el tamaño de los slices
	sizeSlice, errSlice := getSliceSize(v.memory[varID])
	if errSlice != nil {
		fmt.Println(errSlice)
	}

	return Value{value: sizeSlice}
}
func (v *Visitor) VisitVectorIsEmpty(ctx *parser.VectorIsEmptyContext) Value {
	typePrimitive := v.memory[ctx.ID().GetText()].PrimitiveType
	varID := ctx.ID().GetText()
	getSliceSize := func(v Value) (int, error) {
		switch typePrimitive {
		case "String":
			sliceValue, ok := v.value.([]string)
			if !ok {
				return 0, fmt.Errorf("tipo de slice incorrecto")
			}
			return len(sliceValue), nil
		case "Int":
			sliceValue, ok := v.value.([]int64)
			if !ok {
				return 0, fmt.Errorf("tipo de slice incorrecto")
			}
			return len(sliceValue), nil
		case "Float":
			sliceValue, ok := v.value.([]float64)
			if !ok {
				return 0, fmt.Errorf("tipo de slice incorrecto")
			}
			return len(sliceValue), nil
		default:
			return 0, fmt.Errorf("tipo primitivo desconocido")
		}
	}

	// Obtener el tamaño de los slices
	sizeSlice, errSlice := getSliceSize(v.memory[varID])
	if errSlice != nil {
		fmt.Println(errSlice)
	}
	if sizeSlice == 0 {
		return Value{value: true}
	}
	return Value{value: false}
}
func (v *Visitor) VisitVectorGetElement(ctx *parser.VectorGetElementContext) Value {

	typePrimitive := v.memory[ctx.ID().GetText()].PrimitiveType
	varID := ctx.ID().GetText()
	getSliceValue := func(v Value, index int) (interface{}, error) {
		switch typePrimitive {
		case "String":
			sliceValue, ok := v.value.([]string)
			if !ok {
				return nil, fmt.Errorf("tipo de slice incorrecto")
			}
			if index < 0 || index >= len(sliceValue) {
				return nil, fmt.Errorf("índice fuera de rango")
			}
			return sliceValue[index], nil
		case "Float":
			sliceValue, ok := v.value.([]float64)
			if !ok {
				return nil, fmt.Errorf("tipo de slice incorrecto")
			}
			if index < 0 || index >= len(sliceValue) {
				return nil, fmt.Errorf("índice fuera de rango")
			}
			return sliceValue[index], nil
		case "Int":
			sliceValue, ok := v.value.([]int64)
			if !ok {
				return nil, fmt.Errorf("tipo de slice incorrecto")
			}
			if index < 0 || index >= len(sliceValue) {
				return nil, fmt.Errorf("índice fuera de rango")
			}
			return sliceValue[index], nil
		default:
			return nil, fmt.Errorf("tipo primitivo desconocido")
		}
	}
	var int64Value int64 = v.Visit(ctx.Expr()).value.(int64)
	intValue := int(int64Value)

	sliceValu, errSliceValue := getSliceValue(v.memory[varID], intValue)
	if errSliceValue != nil {
		fmt.Println(errSliceValue)
	}
	return Value{value: sliceValu}
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {

	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{value: i}
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) Value {
	f, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Value{value: f}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	value := strings.Trim(ctx.GetText(), "\"")
	return Value{value: value}
}
func (v *Visitor) VisitCharExpr(ctx *parser.CharExprContext) Value {
	value := rune(ctx.GetText()[1]) // Obtiene el caracter entre las comillas simples
	return Value{value: value}      // Suponiendo que Value tiene un campo charValue de tipo rune
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()
	value, ok := v.memory[id]

	if ok {
		return value
	} else {
		panic("no such variable: " + id)

	}

}

// Funciones para asignaciones
func (v *Visitor) VisitVar_type(ctx *parser.Var_typeContext) Value {
	return Value{value: ctx.GetText()}
}
func (v *Visitor) VisitPrimitivo(ctx *parser.PrimitivoContext) Value {
	return Value{value: ctx.GetText()}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{value: value}
}
func (v *Visitor) VisitTransferSentence(ctx *parser.Transfer_sentenceContext) Value {
	return Value{value: ctx.GetText()}
}
func (v *Visitor) VisitOneExpr(ctx *parser.OneExprContext) Value {
	exprValue := v.Visit(ctx.Expr()).value.(interface{})
	v.tmpList = append(v.tmpList, exprValue)
	return Value{value: true}
}

func (v *Visitor) VisitNumExpr(ctx *parser.NumExprContext) Value {
	exprValue := v.Visit(ctx.Expr()).value.(interface{})
	v.tmpList = append(v.tmpList, exprValue) // Add value to tmpList
	if ctx.ListaExp() != nil {
		v.Visit(ctx.ListaExp())
	}
	return Value{value: true}
}

func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext) Value {
	type_left := getType(v.Visit(ctx.GetLeft()).value)
	type_right := getType(v.Visit(ctx.GetRight()).value)
	l := v.Visit(ctx.GetLeft()).value

	r := v.Visit(ctx.GetRight()).value
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		if type_left == "Float" || type_right == "Int" || type_right == "Float" || type_left == "Int" {
			return Value{value: l.(float64) + r.(float64)}
		} else if type_left == "Float" && type_right == "Float" {
			return Value{value: l.(float64) + r.(float64)}
		} else if type_left == "String" && type_right == "String" {
			return Value{value: l.(string) + r.(string)}
		} else if type_left == "Int" && type_right == "Int" {
			return Value{value: l.(int64) + r.(int64)}
		}
	case "-":
		if type_left == "Float" && type_right == "Int" || type_right == "Float" && type_left == "Int" {
			return Value{value: l.(float64) - r.(float64)}
		} else if type_left == "Float" && type_right == "Float" {
			return Value{value: l.(float64) - r.(float64)}
		} else if type_left == "Int" && type_right == "Int" {
			return Value{value: l.(int64) - r.(int64)}
		}
	case "*":
		if type_left == "Float" || type_right == "Int" || type_right == "Float" || type_left == "Int" {
			return Value{value: l.(float64) * r.(float64)}
		} else if type_left == "Float" && type_right == "Float" {
			return Value{value: l.(float64) * r.(float64)}
		} else if type_left == "Int" && type_right == "Int" {
			return Value{value: l.(int64) * r.(int64)}
		}
	case "/":
		if type_left == "Float" || type_right == "Int" || type_right == "Float" || type_left == "Int" {
			return Value{value: l.(float64) / r.(float64)}
		} else if type_left == "Float" && type_right == "Float" {
			return Value{value: l.(float64) / r.(float64)}
		} else if type_left == "Int" && type_right == "Int" {
			return Value{value: l.(int64) / r.(int64)}
		}
	case "%":
		if type_left == "Int" && type_right == "Int" {
			return Value{value: l.(int64) % r.(int64)}
		}
		//Menor que  o menor igual que
	case "<":
		if type_left == "Int" && type_right == "Int" {
			if l.(int64) < r.(int64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Float" && type_right == "Float" {
			if l.(float64) < r.(float64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "String" && type_right == "String" {
			if len(l.(string)) < len(r.(string)) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		}
	case "<=":
		if type_left == "Int" && type_right == "Int" {
			if l.(int64) <= r.(int64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Float" && type_right == "Float" {
			if l.(float64) <= r.(float64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "String" && type_right == "String" {
			if len(l.(string)) <= len(r.(string)) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		}
	case ">":
		if type_left == "Int" && type_right == "Int" {
			if l.(int64) > r.(int64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Float" && type_right == "Float" {
			if l.(float64) > r.(float64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "String" && type_right == "String" {
			if len(l.(string)) > len(r.(string)) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		}
	case ">=":
		if type_left == "Int" && type_right == "Int" {
			if l.(int64) >= r.(int64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Float" && type_right == "Float" {
			if l.(float64) >= r.(float64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "String" && type_right == "String" {
			if len(l.(string)) >= len(r.(string)) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		}
		//----------------------------------------------------------------
	case "==":
		if type_left == "Int" && type_right == "Int" {
			if l.(int64) == r.(int64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Float" && type_right == "Float" {
			if l.(float64) == r.(float64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "String" && type_right == "String" {
			if l.(string) == r.(string) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Bool" && type_right == "Bool" {
			if l.(bool) == r.(bool) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else {
			fmt.Print("Los operandos no coinciden")
			return Value{value: false}
		}
	//----------------------------------------------------------------
	case "!=":
		if type_left == "Int" && type_right == "Int" {
			if l.(int64) != r.(int64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Float" && type_right == "Float" {
			if l.(float64) != r.(float64) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "String" && type_right == "String" {
			if l.(string) != r.(string) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else if type_left == "Bool" && type_right == "Bool" {
			if l.(bool) != r.(bool) {
				return Value{value: true}
			} else {
				return Value{value: false}
			}
		} else {
			fmt.Print("Los operandos no coinciden")
			return Value{value: false}
		}
	case "&&":
		if l.(bool) && r.(bool) {
			return Value{value: true}
		} else {
			return Value{value: false}
		}
	case "||":
		if l.(bool) || r.(bool) {
			return Value{value: true}
		} else {
			return Value{value: false}
		}

	}

	return Value{value: false}

}

func getType(value interface{}) string {
	switch v := value.(type) {
	case bool:
		return "Bool"
	case int:
		return "Int"
	case int64:
		return "Int"
	case string:
		return "String"
	case float64:
		return "Float"
	case rune:
		return "Character"
	default:
		fmt.Print(v, "desconocido")
		return "Unknown"
	}
}

func (v *Visitor) DoesVarExist(varID string) bool {
	_, exists := v.memory[varID]
	return exists
}
func asciiToChar(ascii interface{}) string {
	return string(ascii.(int32))
}
func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/analyze", func(c *fiber.Ctx) error {
		prog := string(c.Body())

		input := antlr.NewInputStream(prog)
		lexer := parser.NewControlLexer(input)
		tokens := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewControlParser(tokens)
		p.BuildParseTrees = true
		tree := p.Prog()
		eval := Visitor{memory: make(map[string]Value)}
		eval.Visit(tree)

		return c.JSON(fiber.Map{
			"prints":  eval.outputBuilder.String(),
			"errores": eval.Errors.String(),
		})
	})

	app.Listen(":8080")
}
