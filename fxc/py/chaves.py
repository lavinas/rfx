# Definição das chaves de duas fontes diferentes

# keys fonte A
key1_A = "alpha"
key2_A = "beta"
key3_A = "gamma"
# keys fonte B
key1_B = "delta"
key2_B = None
key3_B = "gamma"


# Verifica se alguma das chaves coincide
found = False
if key1_A == key1_B:
    found = True
elif key2_A is not None and key2_B is not None and key2_A == key2_B:
    found = True
elif key3_A is not None and key3_B is not None and key3_A == key3_B:
    found = True
print(found)
    
