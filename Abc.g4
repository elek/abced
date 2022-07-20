grammar Abc;

line: unit+ CONT? WS? NL?;
unit: (H? (de | SEP) WS?)+;
note: (ROLL)? (SHARP | NATURAL | FLAT)? PITCH DOWN? NUMBER? DIVIDER* NUMBER?;
de: triplet WS? de | GS DIVIDER? note+ GE | de LT de | de GT de | de (H WS)? SEP | PS de+ PE | note WS? | H WS? de | note TIE;
triplet: PS NUMBER ((COLON NUMBER?)? COLON NUMBER?)?;


PITCH: [cdefgabCDEFGABz];
NUMBER: [0-9];
DIVIDER: '/';
CONT: '\\';
NL: '\n'->skip;
WS: ' ';
SEP: '|';
GS: '{';
GE: '}';
PS: '(';
PE: ')';
COLON: ':';
LT: '<';
GT: '>';
SHARP: '^';
NATURAL: '=';
FLAT: '_';
DOWN: ',';
H: 'H';
ROLL: '!roll!';
TIE: '-';

