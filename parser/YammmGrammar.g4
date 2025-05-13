grammar YammmGrammar; 

schema: schema_name (type | datatype)* EOF;
schema_name: DOC_COMMENT? 'schema' STRING;

type: DOC_COMMENT? ( is_abstract = 'abstract' | is_part = 'part' )?
  'type' type_name extends_types? LBRACE type_body RBRACE
  ;
datatype: DOC_COMMENT? 'type' type_name EQUALS built_in ;

type_name: UC_WORD ;
plural_name: type_name ;

extends_types: 'extends' type_name (COMMA type_name)* COMMA? ;
type_body: (property | association | composition | invariant)* ; 

property: DOC_COMMENT? property_name data_type_ref (is_primary = 'primary' | is_required = 'required')?;
rel_property: DOC_COMMENT? property_name data_type_ref is_required = 'required'?;
property_name: LC_WORD | lc_keyword;

data_type_ref: built_in | alias ;
alias: UC_WORD ;

association: DOC_COMMENT? ASSOC thisName=any_name thisMp=multiplicity? toType=type_name (SLASH reverse_name=any_name reverseMp=multiplicity?)?  (LBRACE relation_body? RBRACE)? ;
composition: DOC_COMMENT? COMP thisName=any_name thisMp=multiplicity? toType=type_name (SLASH reverse_name=any_name reverseMp=multiplicity?)? ;
any_name: UC_WORD | LC_WORD;
multiplicity
  : LPAR ((USCORE (COLON ('one'| 'many'))?) | ('one' (COLON ('one' | 'many'))?) | 'many') RPAR
  ;

relation_body: rel_property+ ;

built_in:
  integerT | floatT | boolT | stringT | enumT | patternT | timestampT | dateT | uuidT | spacevectorT
  ;

integerT: 'Integer' (LBRACK min=(USCORE | INTEGER) COMMA max=(USCORE | INTEGER) RBRACK)?;
floatT: 'Float'     (LBRACK min=(USCORE | INTEGER | FLOAT) COMMA max=(USCORE | INTEGER | FLOAT) RBRACK)?;
boolT: 'Boolean' ;
stringT: 'String'   (LBRACK min=(USCORE | INTEGER) COMMA max=(USCORE | INTEGER) RBRACK)?;
enumT: 'Enum'       LBRACK STRING (COMMA STRING)+ COMMA? RBRACK ;
patternT: 'Pattern' LBRACK STRING (COMMA STRING)? RBRACK;
timestampT: 'Timestamp' (LBRACK format=STRING RBRACK)?;
spacevectorT: 'Spacevector' LBRACK dimensions= INTEGER RBRACK;
dateT: 'Date' ;
uuidT: 'UUID' ;

datatypeKeyword
  : 'Integer' | 'Float' | 'Boolean' | 'String' | 'Enum' | 'Pattern' | 'Timestamp' | 'Date'
  | 'UUID' | 'Spacevector'
  ;
invariant: EXCLAMATION message=STRING constraint=expr ;

expr
  : left=literal # value
  | LBRACK (values+=expr (COMMA values+=expr)* COMMA?)? RBRACK # list
  | op=MINUS right=expr # uminus
  | left=expr LBRACK (right+=expr (COMMA right+=expr)* COMMA?)? RBRACK # at 
  | left=expr op=ARROW name=(LC_WORD | UC_WORD) args=arguments? params=parameters? (LBRACE body=expr RBRACE)? # fcall
  | left=expr PERIOD name=expr # period
  | op=EXCLAMATION right=expr # not
  | left=expr op=(STAR | SLASH | PERCENT) right=expr # muldiv
  | left=expr op=(PLUS | MINUS) right=expr #plusminus
  | left=expr op=(GT | GTE | LT | LTE) right=expr # compare
  | left=expr op='in' right=expr # in
  | left=expr op=(MATCH | NOTMATCH) right=expr # match
  | left=expr op=(EQUAL | NOTEQUAL) right=expr # equality
  | left=expr op=AND right=expr # and
  | left=expr op=(OR | HAT) right=expr # or
  | left=expr op=QMARK LBRACE trueBranch=expr (COLON falseBranch=expr)? RBRACE # if
  | LPAR left=expr RPAR # group
  | left=VARIABLE # variable
  | left=property_name # name
  | left=datatypeKeyword # datatypeName // shadows relation names by design.
  | left=UC_WORD # relationName
  | USCORE # literalNil
  ;

arguments: LPAR (args+=expr (COMMA args+=expr)*)? COMMA? RPAR;
parameters: (PIPE params+=VARIABLE (COMMA params+=VARIABLE)* COMMA? PIPE) ;

// relation_ref: ARROW any_name ;
literal : v=(STRING | BOOLEAN | FLOAT | INTEGER | REGEXP) ;

lc_keyword
  : 'schema' 
  | 'type'
  | 'datatype'
  | 'required'
  | 'primary'
  | 'extends'
  | 'includes'
  | 'abstract'
  | 'one'
  | 'many'
  ;


LBRACE : '{' ;
RBRACE : '}' ;
LBRACK : '[' ;
RBRACK : ']' ;
LPAR : '(' ;
RPAR : ')' ;
COLON : ':' ;
COMMA : ',';
EQUALS : '=';
ASSOC: '-->';
COMP: '*->';
ARROW: '->';
SLASH: '/';
USCORE: '_';
STAR: '*';
AT: '@';
EXCLAMATION: '!';
PLUS: '+';
MINUS: '-';
OR: '||';
AND: '&&';
EQUAL: '==';
NOTEQUAL: '!=';
MATCH: '=~';
NOTMATCH: '!~';
QMARK: '?';
GT: '>';
GTE: '>=';
LT: '<';
LTE: '<=';
DOLLAR: '$';
PIPE: '|';
PERIOD: '.';
PERCENT: '%';
HAT: '^';

// JS String support escaped b,t,n,f,r, and u HEX*4, x HEX*2, and (deprecated (0-7)), 
// A string converter may need to handle the u, x, and octal escapes.
STRING:
  '"' ('\\'('b'|'t'|'n'|'f'|'r'|'u'|'x'|'0'|'"'|'\''|'\\') | ~('\\'|'"'|'\r'|'\n') )* '"' |
  '\''('\\'('b'|'t'|'n'|'f'|'r'|'u'|'x'|'0'|'"'|'\''|'\\') | ~('\\'|'\''|'\r'|'\n') )* '\''
;
DOC_COMMENT : '/*' .*? '*/' ;
SL_COMMENT : '//' ~('\n'|'\r')* ('\r'? '\n')? -> skip;
//REGEXP: '/' ('\\'('\\' | '/') | ~('\\'|'/'|'\r'|'\n'))* '/' ;
REGEXP: '/' ('\\' ('\\' | '/' | .) | ~('\\'|'/'|'\r'|'\n'))* '/' ;
WS : (' '|'\t'|'\r'|'\n')+ -> channel(HIDDEN);
fragment DIGITS : [0-9]+ ;
fragment EDIGITS: DIGITS ('e'|'E')('-'|'+') DIGITS;
VARIABLE: '$' (DIGITS | LC_WORD);
INTEGER: DIGITS;
FLOAT: DIGITS '.' (EDIGITS | DIGITS);
BOOLEAN: 'true' | 'false' ;
UC_WORD: [A-Z]([A-Z]|[a-z]|[0-9]|'_')* ;
LC_WORD: [a-z]([A-Z]|[a-z]|[0-9]|'_')* ;  
ANY_OTHER : . ;
