grammar Blunderbuss;

program: func+ ;

// todo: arrays and indexing arrays

// func
func: (CACHE|LAZY)? FUNC ID args TYPE block ;
args: LPAREN (param (COMMA param)*)? RPAREN ;
param: TYPE ID ;

// func call
call_args: LPAREN (expr (COMMA expr)*)? RPAREN ;
func_call: ID call_args ;

block: LBRACE body RBRACE ;
body: (effect_block|stmt)* ;

expr
    : NUM
    | STRING
    | ID
    | func_call
    | expr BIN_OP expr
    | SIN_OP expr
    ;

// stmts
stmt
    : LAZY? TYPE ID ASSIGN expr SEMI
    | ID ASSIGN expr SEMI
    | RETURN expr SEMI
    | SAFE? func_call SEMI
    | NEXT SEMI
    | BREAK SEMI
    | if_stmt
    | for_stmt
    ;

effect_block: EFFECT LBRACE stmt* RBRACE ;

if_stmt: IF LPAREN expr RPAREN block
    (ELSEIF LPAREN expr RPAREN block)*
    (ELSE block)?
    ;

for_stmt: FOR
    // assign stmt
    LPAREN (TYPE ID ASSIGN expr)?
    // loop condition expression, expression can always be somehow evaluated to bool
    SEMI expr?
    // end loop statement, 
    SEMI (ID ASSIGN expr)?

    RPAREN block ;


// operators
SIN_OP: EXCL ;

BIN_OP
    : PLUS 
    | MINUS
    | PWR
    | MULT
    | DIV
    | MOD
    | EQUAL
    | LE
    | GE
    | LT
    | GT
    | LSHIFT
    | RSHIFT
    | AND
    | OR
    | BIN_AND
    | BIN_OR
    ;


// string
fragment ESCAPE: '\\' [btnr"'\\] ;
STRING
    : SQUOTE (ESCAPE | ~['\\])* SQUOTE
    | DQUOTE (ESCAPE | ~["\\])* DQUOTE
    ;

// expr tokens
PLUS: '+' ;
MINUS: '-' ;
PWR: '**';
MULT: '*' ;
DIV: '/' ;
MOD: '%';
OR: '||' ;
AND: '&&' ;
BIN_OR: '|' ;
BIN_AND: '&' ;
LSHIFT: '<<';
RSHIFT: '>>';
EXCL: '!' ;
// comparison tokens
EQUAL: '==' ;
LE: '=<' ;
GE: '=>' ;
LT: '<' ;
GT: '>' ;
// tokens
NUM: '-'? [0-9]+ ('.' [0-9]*)? ;
RBRACE: '}' ;
LBRACE: '{' ;
LPAREN: '(' ;
RPAREN: ')' ;
COMMA: ',' ;
ASSIGN: '=' ;

ELSEIF: 'elseif' ;
IF: 'if' ;
ELSE: 'else' ;
FOR: 'for' ;
BREAK: 'break' ;
NEXT: 'next' ;
RETURN: 'return' ;
CACHE: 'cache' ;
SAFE: 'safe' ;
EFFECT: 'effect' ;
LAZY: 'lazy' ;
FUNC: 'func' ;
SQUOTE: '\'' ;
DQUOTE: '"' ;
SEMI: ';' ;

// types
TYPE: INT | DOUBLE | STR | BOOL | BYTE | UINT ;
INT: 'int' ;
UINT: 'uint' ;
BOOL: 'bool' ;
BYTE: 'byte' ;
DOUBLE: 'double' ;
STR: 'str' ;

WS  : [ \t\n\r]+ -> skip ;
LINE_COMMENT: '/' '/' ~[\n\r]* -> skip ;
BLOCK_COMMENT: '/' '*' .*? '*' '/' -> skip ;
ID  : [a-zA-Z_] [a-zA-Z0-9_]* ;
