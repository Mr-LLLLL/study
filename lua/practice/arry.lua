arr={"lua", "c"}
for i = 1, 2 do
	print(arr[i])
end

arr[-1] = "hello"
arr.key = "world";

for k,v in pairs(arr) do
	print(k, v)
end


arr = {{"hello", "world"}, {"nihao", "shijie"}}
for k, v in pairs(arr) do
	for k, v in pairs(v) do 
		print(k, v)
	end
end
