# gfw parser

The following will detail the grammar for gfw

# Grammar


PROG := OPTS INTS ALIAS FIREWALL DEFAULTS NATIVE

//options
OPTS := Option Lcurly [Identifier Identifier Endstmt]* Rcurly


// interfaces
INTS := Interface Lcurly [Identifier Identifier NetAddr Endstmt]* Rcurly

// aliases
ALIAS := Alias Lcurly [Identifier HostAddr Endstmt]* Rcurly

// firewall
FIREWALL := Firwall Lculy (Identifier | Star) OP
// defaults


// native
NATIVE := '{'NL SEQ NL'}'NL


KV := IDENTIFIER IDENTIFIER
STMT 
OP := Allow | Block | Reject | NAT

