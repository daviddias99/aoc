#!/usr/bin/env python3

fint = open('day3_1_input.txt', 'r')

line1 = fint.readline().strip('\n').split(',')
line2 = fint.readline().strip('\n').split(',')


def processOrder(x,y,order):

    dir = order[0]
    dist = int(order[1 : 1 + len(order)])

    if dir == 'U':
        y += dist
    elif dir == 'D':
        y -= dist
    elif dir == 'R':
        x += dist
    elif dir == 'L':
        x -= dist

    return x,y

x = 0
y = 0

for s in line1:
    
    x,y = processOrder(x,y,s)



print(x)
print(y)

