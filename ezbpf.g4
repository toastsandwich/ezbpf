grammar ezbpf;

# Fragment rules
fragment NUM      : [0-9] ;
fragment ALPHA    : [a-zA-Z] ;
fragment ALNUM    : (ALPHA | NUM)+ ;
fragment HEX      : '0x' [0-9a-fA-F]+ ;
fragment OCT      : '0' [0-7]+ ;
fragment BIN      : '0b' [01]+ ;
fragment ID_START : [a-zA-Z_] ;
fragment ID_PART  : [a-zA-Z_0-9] ;
fragment ESC      : '\\' [btnr'"\\] ;

# Lexer rules

LPAR: '(' ;
RPAR: ')' ;
LBRA: '{' ;
RBRA: '}' ;
LSQ: '[' ;
RSQ: ']' ;
TLBRA: '<' ;
TRBRA: '>' ;
COMMA: ',' ;
SEMI: ';' ;
COLON: ':' ;
DOT: '.' ;

# Arithmetic operators
ADD: '+' ;
SUB: '-' ;
MUL: '*' ;
DIV: '/' ;
MOD: '%' ;

# Bitwise operators
BIT_AND: '&' ;
BIT_OR: '|' ;
BIT_XOR: '^' ;
LSHIFT: '<<' ;
RSHIFT: '>>' ;

# Comparison operators
EQ: '==' ;
NEQ: '!=' ;
LT: '<' ;
GT: '>' ;
LTE: '<=' ;
GTE: '>=' ;

# Logical operators
AND: '&&' ;
OR: '||' ;
NOT: '!' ;

# Assignment
ASSIGN: '=' ;
ADD_ASSIGN: '+=' ;
SUB_ASSIGN: '-=' ;
MUL_ASSIGN: '*=' ;
DIV_ASSIGN: '/=' ;
MOD_ASSIGN: '%=' ;

HEX_LITERAL   : HEX ;
OCT_LITERAL   : OCT ;
BIN_LITERAL   : BIN ;
DEC_LITERAL   : [0-9]+ ;

CHAR_LITERAL  : '\'' (ESC | ~['\\]) '\'' ;  // Single character
STRING_LITERAL: '"' (ESC | ~["\\])* '"' ;   // String

__U32    : '__u32' ;
__U64    : '__u64' ;
__S32    : '__s32' ;
__S64    : '__s64' ;

U32      : 'u32' ;
U64      : 'u64' ;
S32      : 's32' ;
S64      : 's64' ;

VAR      : 'var' ;
CONST    : 'const' ;
CHAR     : 'char' ;
INT      : 'int' ;
LONG     : 'long' ;
SHORT    : 'short' ;
UINT     : 'uint' ;

STRUCT   : 'struct' ;

FN       : 'fn' ;
RETURN   : 'return' ;

IF       : 'if' ;
ELSEIF   : 'elif' ;
ELSE     : 'else' ;

# BPF map tokens
MAP      : 'map' ;

# Import statements
USE : 'use' ;

# Identifier (placed **AFTER** keywords)
IDENTIFIER : ID_START ID_PART* ;

# Whitespace & Comments
WS            : [ \t\r\n]+ -> skip ;
COMMENT       : '//' ~[\r\n]* -> skip ;
MULTI_COMMENT : '/*' ~[*]* '*/' -> skip ;

# Parser rules

type
    : U32 
    | U64 
    | __U32 
    | __U64 
    | S32 
    | S64 
    | __S32 
    | __S64 
    | CHAR 
    | SHORT 
    | LONG 
    | UINT
    ;

assign
    : ASSIGN
    | ADD_ASSIGN
    | SUB_ASSIGN
    | MUL_ASSIGN
    | DIV_ASSIGN
    | MOD_ASSIGN
    ;

compare
    : EQ
    | NEQ
    | GT
    | LT
    | GTE
    | LTE
    | AND
    | OR
    | NOT
    ;

expression
    : HEX_LITERAL
    | OCT_LITERAL
    | BIN_LITERAL
    | DEC_LITERAL
    | CHAR_LITERAL
    | STRING_LITERAL
    | IDENTIFIER
    | expression DOT IDENTIFIER     // Struct field access
    | expression ADD expression
    | expression SUB expression
    | expression MUL expression
    | expression DIV expression
    | expression MOD expression
    | expression BIT_AND expression
    | expression BIT_OR expression
    | expression BIT_XOR expression
    | expression LSHIFT expression
    | expression RSHIFT expression
    | expression compare expression
    | LPAR expression RPAR
    ;

arg: expression ;
args: arg (COMMA arg)* ;

param: IDENTIFIER COLON type ;
params: param (COMMA param)* ;

useStmt: USE IDENTIFIER SEMI ;

varInitStmt: VAR IDENTIFIER COLON type SEMI ;

varDeclStmt: VAR IDENTIFIER COLON type assign expression SEMI ;

constDeclStmt: CONST IDENTIFIER COLON type ASSIGN expression SEMI;

mapDeclStmt: MAP IDENTIFIER COLON TLBRA type COMMA type COMMA DEC_LITERAL TRBRA SEMI;

structDeclStmt: STRUCT IDENTIFIER LBRA varInitStmt* RBRA SEMI ;

ifStmt: IF LPAR expression RPAR LBRA stmts* RBRA ;

returnStmt: RETURN expression? SEMI ;  // Optional return value

funcDeclStmt: FN IDENTIFIER LPAR params? RPAR COLON type LBRA stmts* RBRA ;

funcCallStmt: IDENTIFIER LPAR args? RPAR SEMI ;

stmts: (varInitStmt
      | varDeclStmt
      | constDeclStmt
      | mapDeclStmt
      | structDeclStmt SEMI
      | ifStmt
      | returnStmt
      | funcDeclStmt
      | funcCallStmt
      )* ;

prog: useStmt? mapDeclStmt* funcDeclStmt+ ;
