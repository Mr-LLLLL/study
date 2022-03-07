#!/usr/local/bin/lua

local t = {1,3,4}

local function test(x, y, z)
    print(x, y, z)
end

test(table.unpack(t))
