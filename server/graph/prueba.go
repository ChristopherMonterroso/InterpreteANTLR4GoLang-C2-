package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CreateGraph(input string) string {
	input = strings.ReplaceAll(input, "\"", "\\\"")
	url := "http://lab.antlr.org/parse/"
	payload := []byte(`{
	"grammar" : "grammar Control;\r\n\r\n// tokens\r\nNIL : 'nil';\r\nINT     : [0-9]+ ;\r\nID      : [a-zA-Z_][a-zA-Z0-9_]* ;\r\nFLOAT     : [0-9]+'.'[0-9]+ ;\r\nSTRING  : '\"' (~[\"\\r\\n] | '\"\"')* '\"' ;\r\nCHAR       : '\"' ~['\\r\\n] '\"' ;\r\n//skip\r\n\r\nWHITESPACE: [ \\\\\\r\\n\\t]+ -> skip;\r\nCOMMENT : '/*' .*? '*/' -> skip;\r\nLINE_COMMENT : '//' ~[\\r\\n]* -> skip;\r\n// rules\r\n\r\nprog: block EOF;\r\n\r\nblock: (stmt)*\r\n     ;\r\n\r\nstmt: assignstmt\r\n    | printlnstmt\r\n    | printstmt\r\n    | ifstmt\r\n    | whilestmt\r\n    | switchstmt\r\n    | forstmt\r\n    | guardstmt\r\n    | vectorPpts\r\n    | matrixstmt\r\n    | funcstmt\r\n    | callFunction\r\n    | returnstmt\r\n    ;\r\n\r\nassignstmt\r\n    : ID  '=' expr                       # reasignacion\r\n    | var_type ID ':' primitivo '=' expr # asignacion\r\n    | var_type ID ':' primitivo '?'      # asignacionNoExpr\r\n    | var_type ID '=' expr               # asignacionNoPrimitive\r\n    | ID '+=' expr                       # incremento\r\n    | ID '-=' expr                       # decremento\r\n    | var_type ID ':' '[' primitivo ']' '=' '['']'          # asignacionVectorVacio\r\n    | var_type ID ':' '[' primitivo ']' '=' '['listaExp']'          # asignacionVector\r\n    | ID '['expr']' '=' expr # reasignacionVector\r\n    | ID '['expr']''['expr']' '=' expr # reasignacionMatrixTwoD\r\n    ;\r\n\r\nvectorPpts\r\n    : ID '.' 'append' '('expr')' #vectorAppend\r\n    | ID '.' 'removeLast' '('')' #vectorRemoveLast\r\n    | ID '.' 'remove' '(' 'at' ':' expr ')' #vectorRemoveAt\r\n    ;\r\nmatrixstmt\r\n    : var_type ID ':' '[''['primitivo']'']' '=' defMatrix #matrixTwoD\r\n    | var_type ID ':' '[''[''['primitivo']'']'']' '=' defMatrix #matrixThreeD\r\n    ;\r\ndefMatrix\r\n    : listaValores_Mat\r\n    ;\r\nlistaValores_Mat\r\n    : '[' listaValores_Mat2 ']'\r\n    ;\r\nlista_Expresiones\r\n    : expr (','expr)*\r\n    ;\r\nlistaValores_Mat2\r\n    : listaValores_Mat2 ',' listaValores_Mat\r\n    | listaValores_Mat\r\n    | lista_Expresiones\r\n    ;\r\n\r\n\r\nfuncstmt\r\n\r\n    : 'func' ID '('listaParams?')' '{' block '}' #func_sinRetorno\r\n    | 'func' ID '('listaParams?')' '->' primitivo '{'block ret='return'? exp=expr? '}' #func_conRetorno_conTipo\r\n    ;\r\n callFunction   \r\n    :  ID '('listaParamsCall?')' \r\n    ;\r\n\r\nlistaParamsCall\r\n    : ( ID ':')? ref='&'? expr (',' ( ID ':')? ref='&'? expr)*\r\n    ;\r\n\r\nlistaParams\r\n    : ext=(ID | '_' )? ID ':' ino='inout'? '['? primitivo isVector=']'? # oneParam\r\n    |  ext=(ID | '_' )? ID ':' ino='inout'? '['? primitivo isVector=']'? ',' listaParams #numParams\r\n    ;\r\n\r\nlistaExp\r\n    : expr #oneExpr\r\n    | expr ',' listaExp # numExpr\r\n    ;\r\nreturnstmt\r\n    : 'return' expr?\r\n    ;\r\nprintlnstmt\r\n    : 'println' '(' expr ')'\r\n    ;\r\n\r\nprintstmt\r\n    : 'print' '(' expr (',' expr)*')'\r\n    ;\r\n\r\nifstmt\r\n    :'if' expr '{' block (transfer_sentence)? '}' 'else' ifstmt       #else_if  \r\n    | 'if' expr '{' block (transfer_sentence)? '}' 'else' '{' block (transfer_sentence)? '}'  #else\r\n    | 'if'  expr  '{' block (transfer_sentence)? '}'                     #ifNormal\r\n    \r\n    \r\n    ;\r\n\r\nswitchstmt\r\n    : 'switch'  expr  '{' cases '}' ;\r\n\r\ncases\r\n    : (caseblock)* ;\r\n\r\ncaseblock\r\n    : 'case' expr ':' block ('break')? # unCase\r\n    | 'default' ':' block  # unDefault\r\n    ;\r\n\r\n\r\nwhilestmt\r\n    : 'while'  expr  '{' block '}'\r\n    ;\r\nforstmt\r\n    : 'for' ID 'in' expr '...' expr '{' block '}' #forNormal\r\n    | 'for' ID 'in' expr '{'block '}'             #forEach\r\n    ;\r\nguardstmt\r\n    : 'guard' expr 'else''{' block (transfer_sentence)? '}'\r\n    ;\r\n\r\nexpr\r\n    : '!' right=expr                        # NotExpr\r\n    | '-'   right=expr                      #negExpr\r\n    | left=expr op='%' right=expr           # OpExpr\r\n    | left=expr op=('*'|'/') right=expr     # OpExpr    \r\n    | left=expr op=('+'|'-') right=expr     # OpExpr\r\n    | left=expr op=('>='|'>') right=expr    # OpExpr\r\n    | left=expr op=('<='|'<') right=expr    # OpExpr\r\n    | left=expr op=('=='|'!=') right=expr   # OpExpr\r\n    | left=expr op=('&&'|'||') right=expr   # OpExpr\r\n    | '(' expr ')'                          # ParExpr\r\n    | NIL                                   # nilExpr\r\n    | INT                                   # IntExpr\r\n    | STRING                                # StrExpr    \r\n    | ('true' | 'false')                    # BoolExpr\r\n    | FLOAT                                 # FloatExpr\r\n    | CHAR                                  # CharExpr\r\n    | ID '('listaParamsCall?')' #callFuncAsExpr\r\n    | ID                                    # IdExpr\r\n    | ID '.' 'isEmpty' #vectorIsEmpty\r\n    | ID '.' 'count' #vectorCount\r\n    | ID '['expr']'  #vectorGetElement\r\n    | ID '['expr']''['expr']' #accesoMatrixTwoD\r\n    | 'String' '('expr')' #toString\r\n    | 'Int' '('expr')'  #toInt\r\n    | 'Float' '('expr')' #toFloat\r\n    \r\n    ;\r\n\r\nprimitivo\r\n    : 'Int'\r\n    | 'String'\r\n    | 'Float'\r\n    | 'Bool'\r\n    | 'Character'\r\n    ;\r\ntransfer_sentence\r\n    : 'continue'\r\n    | 'break'\r\n    ;\r\n\r\nvar_type\r\n    : 'let'\r\n    | 'var'\r\n    ;",
	"input": "` + input + `", 
	"lexgrammar": "",
	"start": "prog"
	}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()
	fmt.Println("Res Status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}
	result := data["result"].(map[string]interface{})
	err = os.WriteFile("cst_graph_tree.svg", []byte(result["svgtree"].(string)), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}
	fmt.Println("it's ok")
	return result["svgtree"].(string)

}
