#!/usr/bin/env python3

fint = open('day2_2_input.txt', 'r')
inputArray = list(map(int,fint.read().split(',')))


def replaceInput(input,verb,noun):
    input[1] = noun
    input[2] = verb

def processMem(input):

    for i in range(0,len(input),4):
        opcode = input[i]

        if opcode == 99:
            break

        op1Addr = input[i+1]
        op2Addr = input[i+2]
        resultAddr = input[i+3]
        result = 0

        if opcode == 1:
            result = input[op1Addr] + input[op2Addr]
        elif opcode == 2:
            result = input[op1Addr] * input[op2Addr]

        input[resultAddr] = result


for noun in range(99):

    for verb in range(99):

        newInput = inputArray.copy()

        replaceInput(newInput,verb,noun)
        processMem(newInput)

        if newInput[0] == 19690720:
            print("noun: %d verb: %d total_ %d" %(noun,verb,noun*100+verb))
            break
    
        

        
