
#!/usr/bin/env python3

import numpy

def removeNL(s):
    return s.rstrip()

def fuel_needed(moduleMass):
    
    return moduleMass//3 - 2


def get_total_fuel_needed_per_module(moduleMass) :

    totalFuelCount = 0
    moduleMass = fuel_needed(moduleMass)

    while moduleMass > 0:

        totalFuelCount += moduleMass
        moduleMass = fuel_needed(moduleMass)

    return totalFuelCount


def full_operation(a):
    return get_total_fuel_needed_per_module(float(removeNL(a)))

fint = open('day1_2_input.txt', 'r')
fuelLines = fint.readlines()
fuelLines = list(map(full_operation,fuelLines))

totalFuel = numpy.sum(fuelLines)

print(totalFuel)
