grammar ezbpf;

compilationUnit : (structDefinition | bpfMapDefinition | globalVariableDeclaration | functionDefinition)* EOF;

structDefinition : 'struct' Identifier '{' structMember* '}' ;
structMember     : Identifier ':' type ';' ;

bpfMapDefinition : 'bpf_map' Identifier '{' bpfMapOption* '}' ';' ;
bpfMapOption     : 'type' ':' Identifier ';'
                 | 'key' ':' type ';'
                 | 'value' ':' type ';'
                 | 'max_entries' ':' INTEGER_LITERAL ';'
                 ;

globalVariableDeclaration : 'var' Identifier ':' type ('=' expression)? ';' ;

functionDefinition : SEC '(' STRING_LITERAL ')' 'fn' Identifier '(' parameterList? ')' (':' type)? block ;
parameterList      : parameter (',' parameter)* ;
parameter          : Identifier ':' type ;

block : '{' statement* '}' ;
statement : variableDeclarationStatement
          | assignmentStatement
          | returnStatement
          | ifStatement
          | whileStatement
          | forStatement
          | functionCallStatement
          | bpfMapOperationStatement
          ;

variableDeclarationStatement : 'var' Identifier ':' type ('=' expression)? ';' ;
assignmentStatement        : Identifier '=' expression ';' ;
returnStatement            : 'return' expression ';' ;
ifStatement                : 'if' '(' expression ')' block ('else' block)? ;
whileStatement             : 'while' '(' expression ')' block ;
forStatement               : 'for' '(' simpleStatement ';' expression ';' simpleStatement ')' block ;
simpleStatement            : assignmentStatement | variableDeclarationStatement | functionCallStatement ;
functionCallStatement      : Identifier '(' argumentList? ')' ';' ;
argumentList             : expression (',' expression)* ;

bpfMapOperationStatement : Identifier '[' expression ']' '=' expression ';' // map[key] = value
                         | Identifier '[' expression ']'                       // Accessing map[key] in an expression
                         | 'delete' Identifier '[' expression ']' ';'         // delete map[key]
                         ;

type : Identifier ('*')? ;

expression : primaryExpression (operator primaryExpression)* ;
primaryExpression : Identifier
                  | INTEGER_LITERAL
                  | CHAR_LITERAL
                  | STRING_LITERAL
                  | '(' expression ')'
                  | functionCallStatement // Allow function calls within expressions
                  | bpfMapOperationStatement // Allow map access within expressions
                  ;

operator : '==' | '!=' | '<' | '>' | '<=' | '>=' | '+' | '-' | '*' | '/' | '%' | '&&' | '||' | '=' ; // Include assignment here for simplicity in 'for' loops

// Fragments for common token patterns
fragment LETTER : [a-zA-Z];
fragment DIGIT  : [0-9];
fragment HEX_DIGIT : [0-9a-fA-F];
fragment ALNUM  : LETTER | DIGIT | '_';

// Lexer rules (tokens)
Identifier       : LETTER+ ALNUM* ;
INTEGER_LITERAL  : DIGIT+ ;
CHAR_LITERAL     : '\'' . '\'' ; // Simple character literal
STRING_LITERAL   : '"' .*? '"' ;
WS               : [ \t\r\n]+ -> skip ;
COMMENT          : '//' .*? '\r'? '\n' -> skip ;
SEC              : 'SEC' ;
LPAREN           : '(' ;
RPAREN           : ')' ;
LBRACE           : '{' ;
RBRACE           : '}' ;
LBRACK           : '[' ;
RBRACK           : ']' ;
COLON            : ':' ;
SEMI             : ';' ;
COMMA            : ',' ;
EQUAL            : '=' ;
PLUS             : '+' ;
MINUS            : '-' ;
STAR             : '*' ;
SLASH            : '/' ;
PERCENT          : '%' ;
EQUAL_EQUAL      : '==' ;
BANG_EQUAL       : '!=' ;
LT               : '<' ;
GT               : '>' ;
LTE              : '<=' ;
GTE              : '>=' ;
AND_AND          : '&&' ;
OR_OR            : '||' ;
IF               : 'if' ;
ELSE             : 'else' ;
WHILE            : 'while' ;
FOR              : 'for' ;
VAR              : 'var' ;
RETURN           : 'return' ;
STRUCT           : 'struct' ;
BPF_MAP          : 'bpf_map' ;
TYPE             : 'type' ;
KEY              : 'key' ;
VALUE            : 'value' ;
MAX_ENTRIES      : 'max_entries' ;
DELETE           : 'delete' ;
FN               : 'fn' ;
