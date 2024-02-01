# from _typeshed import ProfileFunction
# import torch

# print(torch.__version__)

a = ["hello", 1]
a[0] = "world"
b = range(1, 5)
c = (1, 5)
print(c)
print(a)

for value in a:
    print(value)
print(b)
print(list(b))

print("================")


def func(size, *args):
    print(args)
    print(size)


func(1, "ksdjfks")
