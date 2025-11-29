grammar Blunderbuss;

program: (func | extern)+ EOF;


extern
    : EXTERN func
    | EXTERN SYM;

// todo: arrays and indexing arrays

// func
func: CACHE? FUNC ID args TYPE block ;
args: LPAREN (param (COMMA param)*)? RPAREN ;
param: TYPE ID ;

// func call
call_args: LPAREN (expr (COMMA expr)*)? RPAREN ;
func_call: ID call_args ;

block: LBRACE (effect_block|stmt)* RBRACE ;


expr
    : EXCL expr
    | expr op=(AND | OR | EQUAL | NOT_EQUAL | LE | GE | LT | GT) expr
    | expr op=(MULT | DIV) expr
    | expr op=(PLUS | MINUS ) expr
    | expr LBRACKET expr RBRACKET
    | LPAREN expr RPAREN
    | NUM
    | STRING
    | ID
    | func_call
    ;

// stmts
stmt
    : TYPE ID ASSIGN expr SEMI
    | TYPE ID SEMI
    | ID ASSIGN expr SEMI
    | ID LBRACKET expr RBRACKET ASSIGN expr SEMI
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
    // LPAREN (TYPE ID ASSIGN expr)?
    LPAREN stmt
    // loop condition expression, expression can always be somehow evaluated to bool
    expr SEMI
    // end loop statement, 
    // SEMI (ID ASSIGN expr)?
    stmt

    RPAREN block ;



// string
fragment ESCAPE: '\\' [btnr"'\\] ;
STRING
    : SQUOTE (ESCAPE | ~['\\])* SQUOTE
    | DQUOTE (ESCAPE | ~["\\])* DQUOTE
    ;

// expr tokens
PLUS: '+' ;
MINUS: '-' ;
MULT: '*' ;
DIV: '/' ;
OR: '||' ;
AND: '&&' ;
EXCL: '!' ;
// comparison tokens
NOT_EQUAL: 'ne' ;
EQUAL: 'eq' ;
LE: 'le' ;
GE: 'ge' ;
LT: 'lt' ;
GT: 'gt' ;
// tokens
NUM: '-'? [0-9]+ ('.' [0-9]*)? ;
RBRACE: '}' ;
LBRACE: '{' ;
RBRACKET: ']' ;
LBRACKET: '[' ;
LPAREN: '(' ;
RPAREN: ')' ;
COMMA: ',' ;
ASSIGN: '=' ;

EXTERN: 'extern' ;
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
FUNC: 'func' ;
SQUOTE: '\'' ;
DQUOTE: '"' ;
SEMI: ';' ;

// types
TYPE: PTR | INT | STR | BYTE | ANY ;
ANY:  'any' ;
PTR: 'ptr' ;
INT: 'int' ;
BYTE: 'byte' ;
STR: 'str' ;

WS  : [ \t\n\r]+ -> skip ;
LINE_COMMENT: '/' '/' ~[\n\r]* -> skip ;
BLOCK_COMMENT: '/' '*' .*? '*' '/' -> skip ;
ID  : [a-zA-Z_] [a-zA-Z0-9_]* ;
SYM : [a-zA-Z_] [a-zA-Z0-9_]* ;
