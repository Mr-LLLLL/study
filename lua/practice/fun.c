#!/usr/local/bin/lua

function test(...)
	local arg = {...}
	for k,v in pairs(arg) do
		print(v)
	end
	print(#arg, "\n")
end

test()
test(1)
test(1,2)
test(1,2,5)

