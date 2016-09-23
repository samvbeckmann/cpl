:- table expression/2, expr_val/3, simple_expr/3, term/3.

assign([id | [equal_sign | S]], R) :- expression(S, R).
expression([0 | R], R).
expression([1 | R], R).
expression([2 | R], R).
expression([3 | R], R).
expression([4 | R], R).
expression([5 | R], R).
expression([6 | R], R).
expression([7 | R], R).
expression([8 | R], R).
expression([9 | R], R).

expression([lt_paren | S], R) :- expression(S, [rt_paren | R]).
expression(S, R) :- expression(S, [add_op | T]), expression(T, R).
expression(S, R) :- expression(S, [mul_op | T]), expression(T, R).

% Value Section
assign_val([id | [equal_sign | S]], R, V) :- expr_val(S, R, V).

expr_val(S, R, V) :- simple_expr(S, R, V).
expr_val(S, R, V) :- expr_val(S, [pow_op | T], V1),
                     simple_expr(T, R, V2),
                     V is V1 ** V2.

simple_expr(S, R, V) :- term(S, R, V).
simple_expr(S, R, V) :- simple_expr(S, [add_op | T], V1),
                     term(T, R, V2),
                     V is V1 + V2.
simple_expr(S, R, V) :- simple_expr(S, [sub_op | T], V1),
                     term(T, R, V2),
                     V is V1 - V2.

term(S, R, V) :- factor(S, R, V).
term(S, R, V) :- term(S, [mul_op | T], V1),
                 factor(T, R, V2),
                 V is V1 * V2.
term(S, R, V) :- term(S, [div_op | T], V1),
                 factor(T, R, V2),
                 V is V1 / V2.

factor([lt_paren | S], R, V) :- expr_val(S, [rt_paren | R], V).
factor([0 | R], R, 0).
factor([1 | R], R, 1).
factor([2 | R], R, 2).
factor([3 | R], R, 3).
factor([4 | R], R, 4).
factor([5 | R], R, 5).
factor([6 | R], R, 6).
factor([7 | R], R, 7).
factor([8 | R], R, 8).
factor([9 | R], R, 9).
