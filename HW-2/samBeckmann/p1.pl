tokenizer([], []).
tokenizer([32 | R], T) :- tokenizer(R, T).
tokenizer([S0 | R], [TK | T]) :- token(S0, TK), tokenizer(R, T).

token(97, id).
token(98, id).
token(99, id).
token(40, lt_paren).
token(41, rt_paren).
token(43, add_op).
token(42, mul_op).
token(45, sub_op).
token(47, div_op).
token(94, pow_op).
token(61, equal_sign).
token(48, 0).
token(49, 1).
token(50, 2).
token(51, 3).
token(52, 4).
token(53, 5).
token(54, 6).
token(55, 7).
token(56, 8).
token(57, 9).
