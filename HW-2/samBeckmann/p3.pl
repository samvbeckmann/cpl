:- table expression/2, expr_val/3, mul_expr_val/3, pow_expr_val/3.

assign([id, equal_sign | S], R) :- expression(S, R).
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
expression(S, R) :- expression(S, [sub_op | T]), expression(T, R).
expression(S, R) :- expression(S, [div_op | T]), expression(T, R).
expression(S, R) :- expression(S, [pow_op | T]), expression(T, R).

% Value Section
assign_val([id, equal_sign | S], R, V) :- expr_val(S, R, V).

expr_val(S, R, V) :- mul_expr_val(S, R, V).
expr_val(S, R, V) :- expr_val(S, [add_op | T], V1),
                     mul_expr_val(T, R, V2),
                     V is V1 + V2.
expr_val(S, R, V) :- expr_val(S, [sub_op | T], V1),
                     mul_expr_val(T, R, V2),
                     V is V1 - V2.

mul_expr_val(S, R, V) :- pow_expr_val(S, R, V).
mul_expr_val(S, R, V) :- mul_expr_val(S, [mul_op | T], V1),
                         pow_expr_val(T, R, V2),
                         V is V1 * V2.
mul_expr_val(S, R, V) :- mul_expr_val(S, [div_op | T], V1),
                         pow_expr_val(T, R, V2),
                         V is V1 / V2.

pow_expr_val(S, R, V) :- factor_val(S, R, V).
pow_expr_val(S, R, V) :- pow_expr_val(S, [pow_op | T], V1),
                         factor_val(T, R, V2),
                         V is V1 ** V2.

factor_val([lt_paren | S], R, V) :- expr_val(S, [rt_paren | R], V).
factor_val([0 | R], R, 0).
factor_val([1 | R], R, 1).
factor_val([2 | R], R, 2).
factor_val([3 | R], R, 3).
factor_val([4 | R], R, 4).
factor_val([5 | R], R, 5).
factor_val([6 | R], R, 6).
factor_val([7 | R], R, 7).
factor_val([8 | R], R, 8).
factor_val([9 | R], R, 9).
