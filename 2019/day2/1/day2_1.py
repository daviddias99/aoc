#!/usr/bin/env python3

fint = open('day2_1_input.txt', 'r')

inputArray = list(map(int,fint.read().split(',')))

print(inputArray)

for i in range(0,len(inputArray),4):
    opcode = inputArray[i]

    if opcode == 99:
        break

    op1Addr = inputArray[i+1]
    op2Addr = inputArray[i+2]
    resultAddr = inputArray[i+3]
    result = 0

    if opcode == 1:
        result = inputArray[op1Addr] + inputArray[op2Addr]
    elif opcode == 2:
        result = inputArray[op1Addr] * inputArray[op2Addr]
    
    inputArray[resultAddr] = result
    
        

print(inputArray)
        
