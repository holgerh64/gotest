# Der num-Test

# Genauigkeit
x = 1
eps=1.0
while 1.0+eps != 1.0:
    x+=1
    eps=eps/2.0

print('1.0+',eps,'=1.0 nach ',x, 'Iterationen')

# Granularität
x=1
a=1.0
b=1.0
while a!=0.0:
    x+=1
    b=a
    a/=2.0

print(b,'/2.0=0 nach',x,'Iterationen')

# Reiskörner auf einem Sachbrett

a=1.0
b=1.0
print('Feld Nr.: 1 \t Reiskörner: 1 \tSumme: 1')
for x in range(2,65):
    a=a*2
    b=b+a
    print('Feld Nr.:',x, '\tReiskörner: ',a, '\tSumme: ',b)

