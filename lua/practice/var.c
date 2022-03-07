--[[
local b =1;
function test()
	c = 1;
	local b = 2;
	b = 3
	print(c, b)
end

test()
print(c, b)
a = 5,2;

a,b = b, a
b,a = a, b
print(a, b)
--]]

--[[
if a == 5 then
	print("5")
	elseif a == 4
	print("4")
else
	print("0")
end
--]]

--[[
while a > 0 do
	print(a)
	a = a - 1
end
--]]

for i = 0, 10,4 do
	i = i + 10
	print(i)
end

if nil then
	print("1")
elseif 1 then
	print("2")
else
	print("3")
end

--[[
repeat
	print(a)
	a = a + 1
until a > 5
--]]

--[[
do 
	local b = 22
	print(b)
end
--]]
