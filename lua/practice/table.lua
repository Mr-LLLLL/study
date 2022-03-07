tab = {}

tab[1] = "lua"

tab[1] = nil

tab["key"] = "hello"

tab.key1 = "world"

for k, v in pairs(tab) do
	print(k, v)
end

tab = {"lua", "c#", "c", "C++"}

print(table.concat(tab, ",", 2, 4))
print(#tab)

table.insert(tab, 1, "java")

for k, v in pairs(tab) do
	print(k, v)
end

table.remove(tab, 1)

for k, v in ipairs(tab) do
	print(k, v)
end

table.sort(tab)

for k, v in ipairs(tab) do
	print(k, v)
end

tab = {1,2,3,3,3,4}

module = {}
function module.fun()
	print("module.fun()")
e
	d
module.fun()
