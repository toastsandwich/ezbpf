grammar ezbpf;

// Fragment rules
fragment NUM      : [0-9] ;
fragment ALPHA    : [a-zA-Z] ;
fragment ALNUM    : (ALPHA | NUM)+ ;
fragment HEX      : '0x' [0-9a-fA-F]+ ;
fragment OCT      : '0' [0-7]+ ;
fragment BIN      : '0b' [01]+ ;
fragment ID_START : [a-zA-Z_] ;
fragment ID_PART  : [a-zA-Z_0-9] ;
fragment ESC      : '\\' [btnr'"\\] ;

// Operators
ADD: '+' ;
SUB: '-' ;
MUL: '*' ;
DIV: '/' ;
MOD: '%' ;

BIT_AND: '&' ;
BIT_OR: '|' ;
BIT_XOR: '^' ;
LSHIFT: '<<' ;
RSHIFT: '>>' ;

EQ: '==' ;
NEQ: '!=' ;
LT: '<' ;
GT: '>' ;
LTE: '<=' ;
GTE: '>=' ;

AND: '&&' ;
OR: '||' ;
NOT: '!' ;

ASSIGN: '=' ;
ADD_ASSIGN: '+=' ;
SUB_ASSIGN: '-=' ;
MUL_ASSIGN: '*=' ;
DIV_ASSIGN: '/=' ;
MOD_ASSIGN: '%=' ;

// Symbols
LPAR: '(' ; RPAR: ')' ; LBRA: '{' ; RBRA: '}' ;
LSQ: '[' ; RSQ: ']' ; COMMA: ',' ; SEMI: ';' ; COLON: ':' ; DOT: '.' ;

// Literals
HEX_LITERAL   : HEX ;
OCT_LITERAL   : OCT ;
BIN_LITERAL   : BIN ;
DEC_LITERAL   : [0-9]+ ;
CHAR_LITERAL  : '\'' (ESC | ~['\\]) '\'' ;
STRING_LITERAL: '"' (ESC | ~["\\])* '"' ;

// Data types
UPTR32: '__u32' ; UPTR64: '__u64' ; SPTR32: '__s32' ; SPTR64: '__s64' ;
U32: 'u32' ; U64: 'u64' ; S32: 's32' ; S64: 's64' ;
VAR: 'var' ; CONST: 'const' ; CHAR: 'char' ; INT: 'int' ;
LONG: 'long' ; SHORT: 'short' ; UINT: 'uint' ; VOID: 'void' ;
STRUCT: 'struct' ; ETHHDR: 'ethhdr'; IPHDR: 'iphdr'; TCPHDR: 'tcphdr'; UDPHDR:'udphdr';

// Keywords
FN: 'fn' ; RETURN: 'return' ;
IF: 'if' ; ELSEIF: 'elif' ; ELSE: 'else' ;

// BPF Map tokens
MAP: 'map' ;
BPF_MAP_TYPE_HASH: 'BPF_MAP_TYPE_HASH' ;
BPF_MAP_TYPE_ARRAY: 'BPF_MAP_TYPE_ARRAY' ;
BPF_MAP_TYPE_PERCPU_HASH: 'BPF_MAP_TYPE_PERCPU_HASH' ;
BPF_MAP_TYPE_PERCPU_ARRAY: 'BPF_MAP_TYPE_PERCPU_ARRAY' ;
BPF_MAP_TYPE_LRU_HASH: 'BPF_MAP_TYPE_LRU_HASH' ;
BPF_MAP_TYPE_LRU_PERCPU_HASH: 'BPF_MAP_TYPE_LRU_PERCPU_HASH' ;

// Identifier
IDENTIFIER: ID_START ID_PART* ;

// Whitespace & Comments
WS: [ \t\r\n]+ -> skip ;
COMMENT: '//' ~[\r\n]* -> skip ;
MULTI_COMMENT: '/*' ~[*]* '*/' -> skip ;

// Parser rules
type: U32 | U64 | UPTR32 | UPTR64 | S32 | S64 | SPTR32 | SPTR64 | CHAR | SHORT | LONG | UINT | INT | VOID | IDENTIFIER;
assign: ASSIGN | ADD_ASSIGN | SUB_ASSIGN | MUL_ASSIGN | DIV_ASSIGN | MOD_ASSIGN;
compare: EQ | NEQ | GT | LT | GTE | LTE | AND | OR | NOT;

expression
	: HEX_LITERAL #hexLiteralExpression
    | OCT_LITERAL #octLiteralExpression
    | BIN_LITERAL #binLiteralExpression
    | DEC_LITERAL #decLiteralExpression
    | CHAR_LITERAL #charLiteralExpression
    | STRING_LITERAL #stringLiteralExpression
    | IDENTIFIER #identifierExpression
    | funcCallExpression #funccallExpression
    | structInitExpression #structinitExpression
    | expression DOT IDENTIFIER #dotExpression
    | LPAR expression RPAR #parExpression
    | expression MUL expression #mulExpression
    | expression DIV expression #divExpression
    | expression MOD expression #modExpression
    | expression ADD expression #addExpression
    | expression SUB expression #subExpression
    | expression LSHIFT expression #lshiftExpression
    | expression RSHIFT expression #rshiftExpression
    | expression BIT_AND expression #bitandExpression
    | expression BIT_OR expression #bitorExpression
    | expression BIT_XOR expression #bitxorExpression
    | expression compare expression #compareExpression
    ;

funcCallExpression: IDENTIFIER LPAR args? RPAR ;
structInitExpression: IDENTIFIER LBRA structFieldAssign? RBRA ;

structFieldAssign: IDENTIFIER COLON expression (COMMA IDENTIFIER COLON expression)* ;


arg: expression ;
args: arg (COMMA arg)* ;
param: IDENTIFIER COLON type ;
params: param (COMMA param)* ;

varInitStmt: VAR IDENTIFIER COLON type SEMI ;
varDeclStmt: VAR IDENTIFIER COLON type assign expression SEMI ;
constDeclStmt: CONST IDENTIFIER COLON type assign expression SEMI;

mapType: BPF_MAP_TYPE_HASH | BPF_MAP_TYPE_ARRAY | BPF_MAP_TYPE_PERCPU_HASH |
         BPF_MAP_TYPE_PERCPU_ARRAY | BPF_MAP_TYPE_LRU_HASH | BPF_MAP_TYPE_LRU_PERCPU_HASH;

mapDeclStmt: MAP IDENTIFIER COLON mapType LT type COMMA type COMMA DEC_LITERAL GT SEMI;

structDataMemStmt: IDENTIFIER COLON type SEMI ;
structDeclStmt: STRUCT IDENTIFIER LBRA structDataMemStmt* RBRA ;
ifStmt: IF LPAR expression RPAR LBRA stmts RBRA ;
returnStmt: RETURN expression? SEMI ;
funcDeclStmt: FN IDENTIFIER LPAR params? RPAR COLON type LBRA stmts RBRA;

stmt: varInitStmt | varDeclStmt | constDeclStmt | mapDeclStmt | structDeclStmt | ifStmt | returnStmt | funcDeclStmt ;

stmts: stmt* ;

prog: structDeclStmt* mapDeclStmt* funcDeclStmt+ ;

