def derive(coefficient, exponent): 

    return str(coefficient*exponent) + "x^" +str(exponent-1)

# el mejor
def derive(coefficient, exponent): 
    return f'{coefficient * exponent}x^{exponent - 1}'

print(derive(7, 8))
print(derive(5, 9))