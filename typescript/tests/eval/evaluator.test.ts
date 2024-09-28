import {describe, expect, test} from '@jest/globals';
import {create} from "@bufbuild/protobuf";
import {ValueSchema} from "../../src/yaml/value_pb.js";

describe('sum module', () => {
    test('adds 1 + 2 to equal 3', () => {

        expect(create(ValueSchema)).toBe(3);
    });
});