print(string.char(97, 98, 99))
print(string.byte("ab", 2))
print(string.len("helo"))
print(#"hello")

str = "halo"
print(string.rep(str,3))

for word in string.gmatch("hello lua user", "%a+") do
	print(word)
end
