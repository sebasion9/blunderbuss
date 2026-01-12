grammar Blunderbuss;

program: (func | extern)+ EOF;


extern
    : EXTERN func
    | EXTERN SYM;


// func
func: CACHE? FUNC ID args TYPE block ;
args: LPAREN (param (COMMA param)*)? RPAREN ;
param: TYPE ID ;

// func call
call_args: LPAREN (expr (COMMA expr)*)? RPAREN ;
func_call: SAFE? ID call_args ;

block: LBRACE stmt* RBRACE ;


expr
    : EXCL expr
    | AMPS expr
    | expr op=(MULT | DIV | MOD) expr
    | expr op=(PLUS | MINUS ) expr
    | expr op=(AND | OR | EQUAL | NOT_EQUAL | LE | GE | LT | GT) expr
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
    | func_call SEMI
    | NEXT SEMI
    | BREAK SEMI
    | if_stmt
    | for_stmt
    ;


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
AMPS: '&' ;
MOD: '%' ;
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
FUNC: 'func' ;
SQUOTE: '\'' ;
DQUOTE: '"' ;
SEMI: ';' ;

// TYPEs
TYPE: PTR | INT | STR | ANY ;
ANY:  'any' ;
PTR: 'ptr' ;
INT: 'int' ;
STR: 'str' ;

WS  : [ \t\n\r]+ -> skip ;
LINE_COMMENT: '/' '/' ~[\n\r]* -> skip ;
BLOCK_COMMENT: '/' '*' .*? '*' '/' -> skip ;
ID  : [a-zA-Z_] [a-zA-Z0-9_]* ;
