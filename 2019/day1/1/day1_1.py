#!/usr/bin/env python3

import numpy

def removeNL(s):
    return s.rstrip()

def fuel_needed(moduleMass):
    
    return moduleMass//3 - 2

def full_operation(a):
    return fuel_needed(float(removeNL(a)))

fint = open('day1_1_input.txt', 'r')
fuelLines = fint.readlines()
fuelLines = list(map(full_operation,fuelLines))

totalFuel = numpy.sum(fuelLines)

print(totalFuel)





