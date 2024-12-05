grammar AOC;

// Entry point
program: instruction* EOF ;

// Valid mul instruction
instruction: mul | do | don_t | INVALID;

mul: 'mul' '(' NUM ',' NUM ')' ;

do: 'do()' ;

don_t: 'don\'t()' ;

// Numbers (1-3 digit)
NUM: [0-9][0-9]*;

INVALID: . ;