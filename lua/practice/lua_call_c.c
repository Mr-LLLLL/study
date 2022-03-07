#include <stdio.h>
#include <string.h>
#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>

static int add(lua_State* L)
{
    double op1 = luaL_checknumber(L, 1);
    double op2 = luaL_checknumber(L, 2);

    lua_pushnumber(L, op1 + op2);

    return 1;
}

static int sub(lua_State* L)
{
    double op1 = luaL_checknumber(L, 1);
    double op2 = luaL_checknumber(L, 2);

    lua_pushnumber(L, op1 - op2);

    return 1;
}

int main(int argc, char** argv)
{
    lua_State* L = luaL_newstate();
    luaL_openlibs(L);

    char* testfunc = 
        "print('add1:  ' , add1(1.0, 2.0)) \
                    print('sub1:    ' , sub1(20, 1)) \
                    print('add:     ' , _G.add) \
                    print('sub:     ' , _G.sub)   \
                    print('add1:    ' , _G.add1) \
                    print('sub1:    ' , _G.sub1)";

    lua_register(L, "add1", add);
    lua_pushcfunction(L, sub);
    lua_setglobal(L, "sub1");

    if (luaL_dostring(L, testfunc))
        printf("Failed to invoke\n");
    lua_close(L);
    return 0;
}
