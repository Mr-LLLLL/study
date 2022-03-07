arr = {"lua", "c", "cpp"}
for k, v in pairs(arr) do
	print(k, v)
end

arr.key = "hello"
arr[4] = "world"
for k, v in ipairs(arr) do
	print(k, v)
end

function square(state, control)
	if control >= state then
		return nil
	else
		control = control + 1
		return control, control + control
	end
end

for i, j in square, 9, 0 do
	print(i, j)
end

