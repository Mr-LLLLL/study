people = {}

people.sayHi = function(self)
	print("People say hi:"..self.name)
end

function clone(tab)
	local ins = {}
	for key, var in pairs(tab) do
		ins[key] = var
	end
	return ins
end

people.new = function(name)
	local self = clone(people)
	self.name = name

	return self
end

function People(name)
	local self = {}

	local function init()
		self.name = name
	end

	self.sayHi = function ()
		print("say Hi "..self.name)
	end

	init()
	return self
end

---[[
function Man(name)
	local self = People(name)
	
	self.sayHello = function()
		print("Hello "..self.name)
	end

	return self
end
--]]

local m = Man("lisi")
m:sayHello()
m:sayHi()

tab = {1,2,3,4,word = "hedkfjkdfjkdfjllo", "five"}
tab[1]=nil

for key,val in pairs(tab) do
	print(key, val, "end")
end

if 0 then
	print('0 is true')
elseif 1 then
	print("1 is true")
end

print(2e-1)

html =
[[
<html>
<head></head>
<body>
	<a href = "kdjfkdfjk></a>
</body>
</html>
]]

print(html)

str=[[dkfj]]
print("2e2")

print(#str)

num = 5

function fun(n)
	if n == 1 then
		return n;
	else
		return n * fun(n - 1);
	end
end

print(fun(5))

function testFun(tab, fun)
	for k, v in pairs(tab) do
		fun(k, v)
	end
end

tab = {1,2,3}

function f1(k, v)
	print(k,v)
end

testFun(tab, f1)
testFun(tab, 
	function (k, v)
		print(k..","..v)
	end
)

