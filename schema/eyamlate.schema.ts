type Type = 'int' | 'num' | 'bool' | 'str' | 'null' | 'obj' | 'arr';
type Evaluation = {
    as?: Type;
    using?: Array<FunctionDefinition>;
    eval: Expression;
};
type FunctionDefinition = {
    def: string;
    with?: Array<string>;
    value: Expression;
};
type Expression =
    | ScalarExpression
    | null
    | ObjectExpression
    | ArrayExpression
    | ComplexExpression;
type ComplexExpression =
    | ComplexOperation
    | Evaluation
    | VariableReference
    | FunctionCall
    | ConditionalCases;
type ScalarExpression =
    | IntegerExpression
    | NumberExpression
    | BooleanExpression
    | StringExpression;

type IntegerExpression =
    | IntegerValue
    | IntegerOperation
    | (ComplexExpression & { as: 'int' });
type IntegerValue = bigint;
type IntegerOperation =
    | {
        as?: 'int';
        int: Expression | [Expression]
    }
    | {
        as?: 'int';
        plus: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        minus: IntegerExpression | [IntegerExpression] | [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        mul: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        div: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        mod: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        min: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        max: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        bitand: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        bitor: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        bitxor: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        bitinv: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        bitshl: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        bitshr: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        bitushr: [IntegerExpression, IntegerExpression]
    }
    | {
        as?: 'int';
        len: [ArrayExpression]
    }
    | {
        as?: 'int';
        cmp: [ScalarExpression | ArrayExpression, ScalarExpression | ArrayExpression]
    }


type NumberExpression =
    | NumberValue
    | NumberOperation
    | (ComplexExpression & { as: 'num' });
type NumberValue = number;
type NumberOperation =
    | {
        as?: 'num';
        num: Expression | [Expression]
    }
    | {
        as?: 'num';
        plus: [NumberExpression, NumberExpression]
    }
    | {
        as?: 'num';
        minus: NumberExpression | [NumberExpression] | [NumberExpression, NumberExpression]
    }
    | {
        as?: 'num';
        mul: [NumberExpression, NumberExpression]
    }
    | {
        as?: 'num';
        div: [NumberExpression, NumberExpression]
    }
    | {
        as?: 'num';
        min: [NumberExpression, NumberExpression]
    }
    | {
        as?: 'num';
        max: [NumberExpression, NumberExpression]
    }


type BooleanExpression =
    | BooleanValue
    | BooleanOperation
    | (ComplexExpression & { as: 'bool' });
type BooleanValue = boolean;
type BooleanOperation =
    | {
        as?: 'bool';
        bool: Expression | [Expression]
    }
    | {
        as?: 'bool';
        not: [BooleanExpression]
    }
    | {
        as?: 'bool';
        and: [BooleanExpression, BooleanExpression]
    }
    | {
        as?: 'bool';
        or: [BooleanExpression, BooleanExpression]
    }
    | {
        as?: 'bool';
        eq: [Expression, Expression]
    }
    | {
        as?: 'bool';
        ne: [Expression, Expression]
    }
    | {
        as?: 'bool';
        lt: [ScalarExpression | ArrayExpression, ScalarExpression | ArrayExpression]
    }
    | {
        as?: 'bool';
        lte: [ScalarExpression | ArrayExpression, ScalarExpression | ArrayExpression]
    }
    | {
        as?: 'bool';
        gt: [ScalarExpression | ArrayExpression, ScalarExpression | ArrayExpression]
    }
    | {
        as?: 'bool';
        gte: [ScalarExpression | ArrayExpression, ScalarExpression | ArrayExpression]
    }


type StringExpression =
    | StringValue
    | StringOperation
    | (ComplexExpression & { as: 'str' });
type StringValue = string;
type StringOperation =
    | {
        as?: 'str';
        str: Expression | [Expression]
    }
    | {
        as?: 'str';
        get: [StringExpression, IntegerExpression]
    }
    | {
        as?: 'str';
        concat: [StringExpression, StringExpression]
    }
    | {
        as?: 'str';
        head: StringExpression | [StringExpression]
    }
    | {
        as?: 'str';
        tail: StringExpression | [StringExpression]
    }
    | {
        as?: 'str';
        last: StringExpression | [StringExpression]
    }
    | {
        as?: 'str';
        init: StringExpression | [StringExpression]
    }

type ObjectExpression =
    | ObjectValue
    | ObjectOperation
    | (ComplexExpression & { as: 'obj' });
type ObjectValue = {
    as?: 'obj';
    obj: object;
};
type ObjectOperation =
    | {
        as?: 'obj';
        put: [ObjectExpression, StringExpression, Expression]
    }
    | {
        as?: 'obj';
        del: [ObjectExpression, StringExpression]
    }

type ArrayExpression =
    | ArrayValue
    | ArrayOperation
    | ArrayTransform
    | (ComplexExpression & { as: 'arr' });
type ArrayValue = {
    as?: 'arr';
    arr: Array<Expression>;
};
type ArrayOperation =
    | {
        as?: 'arr';
        concat: [ArrayExpression, ArrayExpression]
    }
    | {
        as?: 'arr';
        len: ArrayExpression | [ArrayExpression]
    }
    | {
        as?: 'arr';
        tail: ArrayExpression | [ArrayExpression]
    }
    | {
        as?: 'arr';
        init: ArrayExpression | [ArrayExpression]
    }
    | {
        as?: 'arr';
        add: [ArrayExpression, Expression]
    }
    | {
        as?: 'arr';
        set: [ArrayExpression, IntegerExpression, Expression]
    }
    | {
        as?: 'arr';
        keys: [ArrayExpression]
    }
type ArrayTransform = {
    as?: 'arr';
    for: string;
    in: Expression;
    do: Expression;
};

type ComplexOperation =
    | {
        as?: Type;
        get: [ArrayExpression, IntegerExpression] | [ObjectExpression, StringExpression]
    }


type VariableReference = {
    as?: Type;
    ref: string;
};

type FunctionCall = {
    as?: Type;
    ref: string;
    with: Array<Expression>;
};

type ConditionalCases = {
    as?: Type;
    cases: Array<ConditionalIfThen | ConditionalOtherwise>
};
type ConditionalIfThen = {
    as?: Type;
    if: Expression;
    then: Expression;
};
type ConditionalOtherwise = {
    as?: Type;
    otherwise: Expression;
};