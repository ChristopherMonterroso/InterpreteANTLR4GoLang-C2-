grammar Control;

// tokens

INT     : [0-9]+ ;
ID      : [a-zA-Z_][a-zA-Z0-9_]* ;
FLOAT     : [0-9]+'.'[0-9]+ ;
STRING  : '"' (~["\r\n] | '""')* '"' ;
CHAR       : '"' ~['\r\n] '"' ;
//skip

WS      : [ \n\t]+ -> skip ;

// rules

prog: block EOF;

block: (stmt)*
     ;

stmt: assignstmt
    | printlnstmt
    | printstmt
    | ifstmt
    | whilestmt
    | switchstmt
    | forstmt
    | guardstmt
    | vectorPpts
    ;

assignstmt
    : ID  '=' expr                       # reasignacion
    | var_type ID ':' primitivo '=' expr # asignacion
    | var_type ID ':' primitivo '?'      # asignacionNoExpr
    | var_type ID '=' expr               # asignacionNoPrimitive
    | ID '+=' expr                       # incremento
    | ID '-=' expr                       # decremento
    | var_type ID ':' '[' primitivo ']' '=' '['']'          # asignacionVectorVacio
    | var_type ID ':' '[' primitivo ']' '=' '['listaExp']'          # asignacionVector
    ;

vectorPpts
    : ID '.' 'append' '('expr')' #vectorAppend
    | ID '.' 'removeLast' '('expr')' #vectorRemoveLast
    | ID '.' 'remove' '(' 'at' ':' expr ')' #vectorRemoveAt
    ;
listaExp
    : expr #oneExpr
    | expr ',' listaExp # numExpr
    ;
printlnstmt
    : 'println' '(' expr ')'
    ;

printstmt
    : 'print' '(' expr ')'
    ;

ifstmt  
    : 'if'  expr  '{' block (transfer_sentence)? '}'                     #ifNormal
    | 'if' expr '{' block (transfer_sentence)? '}' 'else' '{' block (transfer_sentence)? '}'  #else
    | 'if' expr '{' block (transfer_sentence)? '}' 'else' ifstmt       #else_if
    ;

switchstmt
    : 'switch'  expr  '{' cases '}' ;

cases
    : (caseblock)* ;

caseblock
    : 'case' expr ':' block ('break')? # unCase
    | 'default' ':' block  # unDefault
    ;


whilestmt
    : 'while'  expr  '{' block '}'
    ;
forstmt
    : 'for' ID 'in' expr '...' expr '{' block '}' #forNormal
    | 'for' ID 'in' expr '{'block '}'             #forEach
    ;
guardstmt
    : 'guard' expr 'else''{' block (transfer_sentence)? '}'
    ;

expr
    : '!' right=expr                        # NotExpr
    | left=expr op='%' right=expr         # OpExpr
    | left=expr op=('*'|'/') right=expr     # OpExpr    
    | left=expr op=('+'|'-') right=expr     # OpExpr
    | left=expr op=('>='|'>') right=expr    # OpExpr
    | left=expr op=('<='|'<') right=expr    # OpExpr
    | left=expr op=('=='|'!=') right=expr   # OpExpr
    | left=expr op=('&&'|'||') right=expr   # OpExpr
    
    | '(' expr ')'                          # ParExpr
    | INT                                   # IntExpr
    | STRING                                # StrExpr    
    | ('true' | 'false')                    # BoolExpr
    | FLOAT                                 # FloatExpr
    | CHAR                                  # CharExpr
    | ID                                    # IdExpr
    | ID '.' 'isEmpty' #vectorIsEmpty
    | ID '.' 'count' #vectorCount
    | ID '['expr']'  #vectorGetElement
    ;


primitivo
    : 'Int'
    | 'String'
    | 'Float'
    | 'Bool'
    | 'Character'
    ;
transfer_sentence
    : 'continue'
    | 'break'
    | 'return'
    ;

var_type
    : 'let'
    | 'var'
    ;