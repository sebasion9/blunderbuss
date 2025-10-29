grammar main;

// todo: for, break, if, else, effect, safe, lazy
program: func+ ;


// func
func: (CACHE|LAZY)? FUNC ID args TYPE block ;
args: LPAREN (param (COMMA param)*)? RPAREN ;
param: TYPE ID ;

// func call
call_args: LPAREN (ID (COMMA ID)*)? RPAREN ;
func_call: ID call_args ;

block: LBRACE body RBRACE ;
body: statement* ;

expression
    : NUM
    | STRING
    | ID
    | func_call
    | expression BIN_OP expression
    | SIN_OP expression
    ;

statement
    : LAZY? TYPE ID ASSIGN expression SEMI
    | RETURN expression SEMI
    | SAFE? func_call SEMI
    ;

// operators
SIN_OP: EXCL ;

BIN_OP
    : PLUS 
    | MINUS
    | PWR
    | MULT
    | DIV
    | MOD
    ;

// string
fragment ESCAPE: '\\' [btnr"'\\] ;
STRING
    : SQUOTE (ESCAPE | ~['\\])* SQUOTE
    | DQUOTE (ESCAPE | ~["\\])* DQUOTE
    ;

// lexer tokens
NUM: '-'? [0-9]+ ('.' [0-9]*)? ;
RBRACE: '}' ;
LBRACE: '{' ;
LPAREN: '(' ;
RPAREN: ')' ;
COMMA: ',' ;
ASSIGN: '=' ;
// expression tokens
PLUS: '+' ;
MINUS: '-' ;
PWR: '**';
MULT: '*' ;
DIV: '/' ;
MOD: '%';
EXCL: '!' ;

RETURN: 'return' ;
CACHE: 'cache' ;
SAFE: 'safe' ;
EFFECT: 'effect' ;
LAZY: 'lazy' ;
FUNC: 'func' ;
SQUOTE: '\'' ;
DQUOTE: '"' ;
SEMI: ';' ;
WS  : [ \t\n\r]+ -> skip ;
LINE_COMMENT: '/' '/' ~[\n\r]* -> skip ;
BLOCK_COMMENT: '/' '*' .*? '*' '/' -> skip ;

// types
TYPE: INT | DOUBLE | STR ;
INT: 'int' ;
DOUBLE: 'double' ;
STR: 'string' ;

ID  : [a-zA-Z_] [a-zA-Z0-9_]* ;
