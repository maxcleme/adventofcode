grammar AOC;

program: registers '\n' instructions ;

registers: (register '\n')* ;

register: 'Register ' ID ': ' INTEGER ;

ID: 'A' | 'B' | 'C' ;

INTEGER: INT ;

fragment INT: '0' | [1-9] DIGIT*;

DIGIT: [0-9] ;


instructions: 'Program: ' instruction (',' instruction)* ;

instruction: adv | bxl | bst | jnz | bxc | out | bdv | cdv ;

adv: '0' ',' operand ;

bxl: '1' ',' operand ;

bst: '2' ',' operand ;

jnz: '3' ',' operand ;

bxc: '4' ',' operand ;

out: '5' ',' operand ;

bdv: '6' ',' operand ;

cdv: '7' ',' operand ;

operand: literal_value | reg_a | reg_b | reg_c ;

literal_value: '0' | '1' | '2' | '3' ;

reg_a: '4';

reg_b: '5';

reg_c: '6';


